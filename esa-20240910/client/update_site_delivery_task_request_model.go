// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *UpdateSiteDeliveryTaskRequest
	GetBusinessType() *string
	SetDiscardRate(v float32) *UpdateSiteDeliveryTaskRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *UpdateSiteDeliveryTaskRequest
	GetFieldName() *string
	SetFilterVer(v string) *UpdateSiteDeliveryTaskRequest
	GetFilterVer() *string
	SetSiteId(v int64) *UpdateSiteDeliveryTaskRequest
	GetSiteId() *int64
	SetTaskName(v string) *UpdateSiteDeliveryTaskRequest
	GetTaskName() *string
}

type UpdateSiteDeliveryTaskRequest struct {
	BusinessType *string  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DiscardRate  *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateSiteDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *UpdateSiteDeliveryTaskRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *UpdateSiteDeliveryTaskRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateSiteDeliveryTaskRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *UpdateSiteDeliveryTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateSiteDeliveryTaskRequest) SetBusinessType(v string) *UpdateSiteDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetDiscardRate(v float32) *UpdateSiteDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetFieldName(v string) *UpdateSiteDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetFilterVer(v string) *UpdateSiteDeliveryTaskRequest {
	s.FilterVer = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetSiteId(v int64) *UpdateSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) SetTaskName(v string) *UpdateSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateSiteDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
