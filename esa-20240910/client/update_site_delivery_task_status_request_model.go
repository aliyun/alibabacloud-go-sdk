// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *UpdateSiteDeliveryTaskStatusRequest
	GetMethod() *string
	SetSiteId(v int64) *UpdateSiteDeliveryTaskStatusRequest
	GetSiteId() *int64
	SetTaskName(v string) *UpdateSiteDeliveryTaskStatusRequest
	GetTaskName() *string
}

type UpdateSiteDeliveryTaskStatusRequest struct {
	// Specifies whether to enable the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The name of the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskStatusRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateSiteDeliveryTaskStatusRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteDeliveryTaskStatusRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetMethod(v string) *UpdateSiteDeliveryTaskStatusRequest {
	s.Method = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetSiteId(v int64) *UpdateSiteDeliveryTaskStatusRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusRequest) SetTaskName(v string) *UpdateSiteDeliveryTaskStatusRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateSiteDeliveryTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
