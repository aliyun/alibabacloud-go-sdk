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
	// The log category. Valid values:
	//
	// 	- dcdn_log_access_l1 (default): access logs.
	//
	// 	- dcdn_log_er: Edge Routine logs.
	//
	// 	- dcdn_log_waf: firewall logs.
	//
	// 	- dcdn_log_ipa: TCP/UDP proxy logs.
	//
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The discard rate.
	//
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The log fields that you want to include in real-time logs to be delivered. Separate the log fields with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ClientIP,UserAgent
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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
