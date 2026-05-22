// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetSiteDeliveryTaskResponseBody
	GetBusinessType() *string
	SetDataCenter(v string) *GetSiteDeliveryTaskResponseBody
	GetDataCenter() *string
	SetDeliveryType(v string) *GetSiteDeliveryTaskResponseBody
	GetDeliveryType() *string
	SetDiscardRate(v float32) *GetSiteDeliveryTaskResponseBody
	GetDiscardRate() *float32
	SetFieldList(v string) *GetSiteDeliveryTaskResponseBody
	GetFieldList() *string
	SetFilterRules(v string) *GetSiteDeliveryTaskResponseBody
	GetFilterRules() *string
	SetFilterVer(v string) *GetSiteDeliveryTaskResponseBody
	GetFilterVer() *string
	SetRawRule(v string) *GetSiteDeliveryTaskResponseBody
	GetRawRule() *string
	SetRequestId(v string) *GetSiteDeliveryTaskResponseBody
	GetRequestId() *string
	SetSinkConfig(v interface{}) *GetSiteDeliveryTaskResponseBody
	GetSinkConfig() interface{}
	SetSiteId(v int64) *GetSiteDeliveryTaskResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetSiteDeliveryTaskResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetSiteDeliveryTaskResponseBody
	GetStatus() *string
	SetTaskName(v string) *GetSiteDeliveryTaskResponseBody
	GetTaskName() *string
}

type GetSiteDeliveryTaskResponseBody struct {
	BusinessType *string     `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string     `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	DeliveryType *string     `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	DiscardRate  *float32    `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	FieldList    *string     `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	FilterRules  *string     `json:"FilterRules,omitempty" xml:"FilterRules,omitempty"`
	FilterVer    *string     `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	RawRule      *string     `json:"RawRule,omitempty" xml:"RawRule,omitempty"`
	RequestId    *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SinkConfig   interface{} `json:"SinkConfig,omitempty" xml:"SinkConfig,omitempty"`
	SiteId       *int64      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName     *string     `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status       *string     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskName     *string     `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetSiteDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteDeliveryTaskResponseBody) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetSiteDeliveryTaskResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *GetSiteDeliveryTaskResponseBody) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *GetSiteDeliveryTaskResponseBody) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *GetSiteDeliveryTaskResponseBody) GetFieldList() *string {
	return s.FieldList
}

func (s *GetSiteDeliveryTaskResponseBody) GetFilterRules() *string {
	return s.FilterRules
}

func (s *GetSiteDeliveryTaskResponseBody) GetFilterVer() *string {
	return s.FilterVer
}

func (s *GetSiteDeliveryTaskResponseBody) GetRawRule() *string {
	return s.RawRule
}

func (s *GetSiteDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteDeliveryTaskResponseBody) GetSinkConfig() interface{} {
	return s.SinkConfig
}

func (s *GetSiteDeliveryTaskResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteDeliveryTaskResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetSiteDeliveryTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSiteDeliveryTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSiteDeliveryTaskResponseBody) SetBusinessType(v string) *GetSiteDeliveryTaskResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDataCenter(v string) *GetSiteDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDeliveryType(v string) *GetSiteDeliveryTaskResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetDiscardRate(v float32) *GetSiteDeliveryTaskResponseBody {
	s.DiscardRate = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetFieldList(v string) *GetSiteDeliveryTaskResponseBody {
	s.FieldList = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetFilterRules(v string) *GetSiteDeliveryTaskResponseBody {
	s.FilterRules = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetFilterVer(v string) *GetSiteDeliveryTaskResponseBody {
	s.FilterVer = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetRawRule(v string) *GetSiteDeliveryTaskResponseBody {
	s.RawRule = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetRequestId(v string) *GetSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSinkConfig(v interface{}) *GetSiteDeliveryTaskResponseBody {
	s.SinkConfig = v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSiteId(v int64) *GetSiteDeliveryTaskResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetSiteName(v string) *GetSiteDeliveryTaskResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetStatus(v string) *GetSiteDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) SetTaskName(v string) *GetSiteDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetSiteDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
