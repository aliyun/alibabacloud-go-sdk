// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceTaskResponseBody
	GetRequestId() *string
	SetServiceTask(v map[string]interface{}) *GetServiceTaskResponseBody
	GetServiceTask() map[string]interface{}
}

type GetServiceTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task details object. Common fields include taskId, serviceId, workspace, regionId, ip, taskType, extraInfo (taskConfig JSON for LiveDebug), createTime, and updateTime.
	//
	// example:
	//
	// {"taskId":"a1b2c3d4-e5f6-7890-abcd-ef1234567890","serviceId":"ggxw4lnjuz@f2fd3a6265a254a052afb","taskType":"live_debug_log_probe","ip":"10.0.0.1","extraInfo":"{\\"probeType\\":\\"LOG\\",\\"language\\":\\"java\\"}"}
	ServiceTask map[string]interface{} `json:"serviceTask,omitempty" xml:"serviceTask,omitempty"`
}

func (s GetServiceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceTaskResponseBody) GetServiceTask() map[string]interface{} {
	return s.ServiceTask
}

func (s *GetServiceTaskResponseBody) SetRequestId(v string) *GetServiceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTaskResponseBody) SetServiceTask(v map[string]interface{}) *GetServiceTaskResponseBody {
	s.ServiceTask = v
	return s
}

func (s *GetServiceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
