// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateApp(v bool) *ImportJobsRequest
	GetAutoCreateApp() *bool
	SetClusterId(v string) *ImportJobsRequest
	GetClusterId() *string
	SetContent(v string) *ImportJobsRequest
	GetContent() *string
	SetOverwrite(v bool) *ImportJobsRequest
	GetOverwrite() *bool
}

type ImportJobsRequest struct {
	// example:
	//
	// true
	AutoCreateApp *bool `json:"AutoCreateApp,omitempty" xml:"AutoCreateApp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// {
	//
	//   "kind": "SchedulerXJobs",
	//
	//   "type": "JSON",
	//
	//   "version": "2.0",
	//
	//   "content": [
	//
	//     {
	//
	//       "appName": "xxl-job-executor-perf-test-xx",
	//
	//       "groupId": "xxl-job-executor-perf-test-xx",
	//
	//       "description": "xxl-job-executor-xx",
	//
	//       "jobConfigInfo": [
	//
	//         {
	//
	//           "jobHandler": "testJobVoidHandler",
	//
	//           "dataOffset": 0,
	//
	//           "executeMode": "standalone",
	//
	//           "monitorConfigInfo": {
	//
	//             "alarmType": "CustomContacts",
	//
	//             "failLimitTimes": 1,
	//
	//             "failEnable": true,
	//
	//             "failRate": 100,
	//
	//             "timeoutKillEnable": false,
	//
	//             "missWorkerEnable": false,
	//
	//             "sendChannel": "webhook",
	//
	//             "timeoutEnable": true,
	//
	//             "timeout": 7200,
	//
	//             "daysOfDeadline": 0,
	//
	//             "successNotice": false
	//
	//           },
	//
	//           "attemptInterval": 30,
	//
	//           "cleanMode": "{\\"cleanMode\\":\\"NUM_ONLY\\",\\"totalRemain\\":300}",
	//
	//           "description": "",
	//
	//           "routeStrategy": 1,
	//
	//           "userName": "xx",
	//
	//           "userId": "xx",
	//
	//           "content": "{\\"jobHandler\\":\\"testJobVoidHandler\\"}",
	//
	//           "maxConcurrency": 1,
	//
	//           "maxAttempt": 0,
	//
	//           "name": "perf_auto_test_0",
	//
	//           "xattrs": "",
	//
	//           "jobType": "xxljob",
	//
	//           "contentType": 1,
	//
	//           "parameters": "success-withMsg",
	//
	//           "timeConfig": {
	//
	//             "calendar": "",
	//
	//             "dataOffset": 0,
	//
	//             "timeType": 1,
	//
	//             "paramMap": {},
	//
	//             "timeExpression": "	- 	- 	- 	- 	- ?"
	//
	//           },
	//
	//           "contactInfoList": [],
	//
	//           "status": 0
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
}

func (s ImportJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportJobsRequest) GoString() string {
	return s.String()
}

func (s *ImportJobsRequest) GetAutoCreateApp() *bool {
	return s.AutoCreateApp
}

func (s *ImportJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportJobsRequest) GetContent() *string {
	return s.Content
}

func (s *ImportJobsRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportJobsRequest) SetAutoCreateApp(v bool) *ImportJobsRequest {
	s.AutoCreateApp = &v
	return s
}

func (s *ImportJobsRequest) SetClusterId(v string) *ImportJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportJobsRequest) SetContent(v string) *ImportJobsRequest {
	s.Content = &v
	return s
}

func (s *ImportJobsRequest) SetOverwrite(v bool) *ImportJobsRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportJobsRequest) Validate() error {
	return dara.Validate(s)
}
