
Vue.component('issue-details', {
  props: ['issues', 'currentID', 'handleUpdateClick'],
  template: `<div class="">
              <main>
              <h2> Issue Details </h2>
              <form>
                  <div class="form-field">

                  Customer:<br>
                  <input v-model="issues[currentID].Customer"><br><br>
                  Description:<br>
                  <input v-model="issues[currentID].Description"><br><br>
                  Tags:<br>
                  <input v-model="issues[currentID].Tags"><br><br>
                  Status:<br>
                  <input v-model="issues[currentID].Status"><br><br>
                  Devices (optional):<br>
                  <input v-model="issues[currentID].Devices"><br><br>
                  Contact Email:<br>
                  <input v-model="issues[currentID].Contact_email"><br><br>
                  Contact Name:<br>
                  <input v-model="issues[currentID].Contact_name"><br><br>
                  <input class="a-btn--filled" v-on:click="handleUpdateClick($event, issues[currentID])" type="submit" value="Submit">
                  {{issues[currentID].ID}}
                  </div>
              </form>

              </main>

            </div>`
})
Vue.component('issue-list', {
   props: ['issues', 'handleDetailClick', 'handleNewClick'],
  template: `<div>
        <h2>Issues</h2>
        <div style="margin-bottom: 10px">

          <span>
            Status:
            <select>
              <option>Open</option>
              <option>Resolved</option>
              <option>All</option>
            </select>
          </span>

          <span>
            Tag:
            <select>
              <option>All</option>
              <option>gpslev</option>
              <option>billing</option>
              <option>cell-reception</option>
            </select>
          </span>

        </div>

        <table class="table table-hover">
            <thead>
                <tr>
                    <!-- <th>ID</th> -->
                    <th>Date Created/Modified</th>
                    <th>Customer (optional)</th>
                    <th>Description</th>
                    <th>Tags</th>
                    <th>Status</th>
                    <th>Devices (optional)</th>
                    <th>Contact Email</th>
                    <th>Contact Name</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(issue, index) in issues" >
                    <!-- <th scope="row">10</th> -->
                    <td> {{ issue.CreatedAt }} /<br/> {{ issue.UpdatedAt }}</td>
                    <td> {{ issue.Customer }} </td>
                    <td> {{ issue.Description }} </td>
                    <td>
                      <span class="label label-info"> {{ issue.Tags }} </span>
                    </td>
                    <td>Open</td>
                    <td> {{ issue.Devices }} </td>
                    <td> {{ issue.Contact_email }} </td>
                    <td> {{ issue.Contact_name }} </td>

                    <td>
                    <a href="#" v-on:click="handleDetailClick($event, index)">View/Edit &raquo;</a>
                    </td>
                </tr>

            </tbody>
        </table>

        <button v-on:click="handleNewClick" type="button" class="new-signup-button btn btn-primary pull-right">New Issue</button>

        TODO: can we put images in here?

      </div>`


})


new Vue({
  el: '#issueTracker',
  data: {
    currentView: 'issue-list',
    currentIssue: [],
    currentID: '',
    //debug: true,
    issues: []
  },
  methods: {
    loadIssues: function() {
      var request = new XMLHttpRequest();
      request.open('GET', 'http://localhost:8080/api/issue', true);

      request.onload = function() {
        if (request.status >= 200 && request.status < 400) {
          // Success!
          var data = JSON.parse(request.responseText);
           this.issues = data.Value;
        } else {
          // We reached our target server, but it returned an error
        }
      }.bind(this);

      request.onerror = function() {
        // There was a connection error of some sort
      };

      request.send();

    },createIssue: function() {
      var request = new XMLHttpRequest();
      request.open('POST', 'http://localhost:8080/api/issue', true);

      request.onload = function() {
        if (request.status >= 200 && request.status < 400) {
          // Success!
          var data = JSON.parse(request.responseText);
           this.issues = data.Value;
        } else {
          // We reached our target server, but it returned an error
        }
      }.bind(this);

      request.onerror = function() {
        // There was a connection error of some sort
      };

      request.send();

    },handleDetailClick: function(event, id) {
            this.currentID = id
            //var request = new XMLHttpRequest();
            //request.open('GET', 'http://localhost:8080/api/issue/' + id, true);

            //request.onload = function() {
            //  if (request.status >= 200 && request.status < 400) {
                // Success!
            //    var data = JSON.parse(request.responseText);
          //       this.currentIssue = data.Value;
          //    } else {
                // We reached our target server, but it returned an error
          //    }
          //  }.bind(this);

          //  request.onerror = function() {
              // There was a connection error of some sort
        //    };

          //  request.send();
            this.currentView = "issue-details"

    },handleNewClick: function() {
      console.log('test new item')

    },handleUpdateClick: function($event, issue) {
      console.log(issue)
      //this.currentID = id
      var request = new XMLHttpRequest();
      request.open('PATCH', 'http://localhost:8080/api/issue/' + issue.ID, true);
      request.setRequestHeader('Content-Type', 'application/json');

      request.onload = function() {
        if (request.status >= 200 && request.status < 400) {
          // Success!
        } else {
          // We reached our target server, but it returned an error
        }
      };

      request.onerror = function() {
        // There was a connection error of some sort
      };

      request.send(JSON.stringify(issue));
      this.currentView = "issue-list"

    }

  }, beforeMount(){
    this.loadIssues()
 }
});
