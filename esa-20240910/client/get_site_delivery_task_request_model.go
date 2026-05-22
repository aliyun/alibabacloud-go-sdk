// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetSiteDeliveryTaskRequest
	GetSiteId() *int64
	SetTaskName(v string) *GetSiteDeliveryTaskRequest
	GetTaskName() *string
}

type GetSiteDeliveryTaskRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 123456***
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

func (s GetSiteDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSiteDeliveryTaskRequest) SetSiteId(v int64) *GetSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteDeliveryTaskRequest) SetTaskName(v string) *GetSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *GetSiteDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
