// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateApp(v bool) *ImportWorkflowsRequest
	GetAutoCreateApp() *bool
	SetClusterId(v string) *ImportWorkflowsRequest
	GetClusterId() *string
	SetContent(v string) *ImportWorkflowsRequest
	GetContent() *string
	SetOverwrite(v bool) *ImportWorkflowsRequest
	GetOverwrite() *bool
}

type ImportWorkflowsRequest struct {
	// example:
	//
	// true
	AutoCreateApp *bool `json:"AutoCreateApp,omitempty" xml:"AutoCreateApp,omitempty"`
	// A short description of struct
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// [{"kind":"SchedulerXWorkflows","type":"JSON","version":"2.0","workflowInfo":{"name":"myWorkflow","description":"","appName":"xuerentest","appType":1,"maxConcurrency":1,"currentExecuteStatus":0,"timeConfig":{"timeType":1,"timeExpression":"0 0 12 	- 	- ?","dataOffset":0}},"nodes":[{"name":"Java1","startTime":-1,"coordinate":{"x":-222.0,"y":40.0,"width":220.0,"height":76.0},"appName":"xuerentest","description":"","jobType":"xxljob","executeMode":"standalone","contentType":1,"content":"{\\"jobHandler\\":\\"helloworld\\"}","xattrs":"{\\"executorBlockStrategy\\":1}","dependentStrategy":1,"routeStrategy":1,"parameters":"","maxConcurrency":1,"maxAttempt":0,"attemptInterval":30,"priority":5,"weight":1,"timeConfig":{"timeType":1,"calendar":"","dataOffset":0},"monitorConfigInfo":{"timeoutEnable":true,"failEnable":true,"failLimitTimes":1,"failRate":100,"missWorkerEnable":true,"timeout":300,"timeoutKillEnable":false,"daysOfDeadline":0,"sendChannel":"","alarmType":"CustomContacts","successNotice":false,"endEarlyEnable":false,"endEarly":30},"contactInfoList":[]},{"name":"shell1","startTime":-1,"coordinate":{"x":102.0,"y":-51.0,"width":220.0,"height":76.0},"appName":"xuerentest","description":"","jobType":"script_shell","executeMode":"standalone","contentType":2,"content":"echo \\"hello world\\"","xattrs":"{\\"executorBlockStrategy\\":1}","dependentStrategy":1,"routeStrategy":1,"parameters":"","maxConcurrency":1,"maxAttempt":0,"attemptInterval":30,"priority":5,"weight":1,"timeConfig":{"timeType":1,"calendar":"","dataOffset":0},"monitorConfigInfo":{"timeoutEnable":true,"failEnable":true,"failLimitTimes":1,"failRate":100,"missWorkerEnable":true,"timeout":300,"timeoutKillEnable":false,"daysOfDeadline":0,"sendChannel":"","alarmType":"CustomContacts","successNotice":false,"endEarlyEnable":false,"endEarly":30},"contactInfoList":[]},{"name":"Java2","startTime":-1,"coordinate":{"x":390.0,"y":55.0,"width":220.0,"height":76.0},"appName":"xuerentest","description":"","jobType":"xxljob","executeMode":"standalone","contentType":1,"content":"{\\"jobHandler\\":\\"helloworld2\\"}","xattrs":"{\\"executorBlockStrategy\\":1,\\"localParams\\":[]}","dependentStrategy":1,"routeStrategy":1,"parameters":"","maxConcurrency":1,"maxAttempt":0,"attemptInterval":30,"priority":5,"weight":1,"timeConfig":{"timeType":1,"calendar":"","dataOffset":0},"monitorConfigInfo":{"timeoutEnable":true,"failEnable":true,"failLimitTimes":1,"failRate":100,"missWorkerEnable":true,"timeout":300,"timeoutKillEnable":false,"daysOfDeadline":0,"sendChannel":"","alarmType":"CustomContacts","successNotice":false,"endEarlyEnable":false,"endEarly":30},"contactInfoList":[]},{"name":"shell2","startTime":-1,"coordinate":{"x":89.0,"y":161.0,"width":220.0,"height":76.0},"appName":"xuerentest","description":"","jobType":"script_shell","executeMode":"standalone","contentType":2,"content":"echo \\"hello world2\\"","xattrs":"{\\"executorBlockStrategy\\":1}","dependentStrategy":1,"routeStrategy":1,"parameters":"","maxConcurrency":1,"maxAttempt":0,"attemptInterval":30,"priority":5,"weight":1,"timeConfig":{"timeType":1,"calendar":"","dataOffset":0},"monitorConfigInfo":{"timeoutEnable":true,"failEnable":true,"failLimitTimes":1,"failRate":100,"missWorkerEnable":true,"timeout":300,"timeoutKillEnable":false,"daysOfDeadline":0,"sendChannel":"","alarmType":"CustomContacts","successNotice":false,"endEarlyEnable":false,"endEarly":30},"contactInfoList":[]}],"edges":[{"from":"Java1","to":"shell1"},{"from":"Java1","to":"shell2"},{"from":"Schedulerx-Root","to":"Java1"},{"from":"shell1","to":"Java2"},{"from":"shell2","to":"Java2"}]}]
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
}

func (s ImportWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowsRequest) GetAutoCreateApp() *bool {
	return s.AutoCreateApp
}

func (s *ImportWorkflowsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportWorkflowsRequest) GetContent() *string {
	return s.Content
}

func (s *ImportWorkflowsRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *ImportWorkflowsRequest) SetAutoCreateApp(v bool) *ImportWorkflowsRequest {
	s.AutoCreateApp = &v
	return s
}

func (s *ImportWorkflowsRequest) SetClusterId(v string) *ImportWorkflowsRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportWorkflowsRequest) SetContent(v string) *ImportWorkflowsRequest {
	s.Content = &v
	return s
}

func (s *ImportWorkflowsRequest) SetOverwrite(v bool) *ImportWorkflowsRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
