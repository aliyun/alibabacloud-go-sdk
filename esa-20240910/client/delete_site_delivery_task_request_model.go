// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteSiteDeliveryTaskRequest
	GetSiteId() *int64
	SetTaskName(v string) *DeleteSiteDeliveryTaskRequest
	GetTaskName() *string
}

type DeleteSiteDeliveryTaskRequest struct {
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteSiteDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteSiteDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DeleteSiteDeliveryTaskRequest) SetSiteId(v int64) *DeleteSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteSiteDeliveryTaskRequest) SetTaskName(v string) *DeleteSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *DeleteSiteDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
