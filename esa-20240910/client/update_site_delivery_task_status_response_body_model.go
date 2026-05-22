// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteDeliveryTaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateSiteDeliveryTaskStatusResponseBody
	GetStatus() *string
	SetTaskName(v string) *UpdateSiteDeliveryTaskStatusResponseBody
	GetTaskName() *string
}

type UpdateSiteDeliveryTaskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the delivery task. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the delivery task.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetRequestId(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetStatus(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) SetTaskName(v string) *UpdateSiteDeliveryTaskStatusResponseBody {
	s.TaskName = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
