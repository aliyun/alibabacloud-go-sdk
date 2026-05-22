// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteDeliveryTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetBusinessType() *string
	SetDataCenter(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetDataCenter() *string
	SetDeliveryType(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetDeliveryType() *string
	SetDiscardRate(v float32) *CreateSiteDeliveryTaskShrinkRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetFieldName() *string
	SetFilterVer(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetFilterVer() *string
	SetHttpDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetHttpDeliveryShrink() *string
	SetKafkaDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetKafkaDeliveryShrink() *string
	SetOssDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetOssDeliveryShrink() *string
	SetS3DeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetS3DeliveryShrink() *string
	SetSiteId(v int64) *CreateSiteDeliveryTaskShrinkRequest
	GetSiteId() *int64
	SetSlsDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetSlsDeliveryShrink() *string
	SetTaskName(v string) *CreateSiteDeliveryTaskShrinkRequest
	GetTaskName() *string
}

type CreateSiteDeliveryTaskShrinkRequest struct {
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	DeliveryType *string  `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	DiscardRate  *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	FieldName           *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer           *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	HttpDeliveryShrink  *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	OssDeliveryShrink   *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	S3DeliveryShrink    *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	// This parameter is required.
	SiteId            *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SlsDeliveryShrink *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetHttpDeliveryShrink() *string {
	return s.HttpDeliveryShrink
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetKafkaDeliveryShrink() *string {
	return s.KafkaDeliveryShrink
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetOssDeliveryShrink() *string {
	return s.OssDeliveryShrink
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetS3DeliveryShrink() *string {
	return s.S3DeliveryShrink
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetSlsDeliveryShrink() *string {
	return s.SlsDeliveryShrink
}

func (s *CreateSiteDeliveryTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetBusinessType(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDataCenter(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDeliveryType(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetDiscardRate(v float32) *CreateSiteDeliveryTaskShrinkRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetFieldName(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.FieldName = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetFilterVer(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.FilterVer = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetHttpDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.HttpDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetKafkaDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.KafkaDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetOssDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.OssDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetS3DeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.S3DeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetSiteId(v int64) *CreateSiteDeliveryTaskShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetSlsDeliveryShrink(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.SlsDeliveryShrink = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) SetTaskName(v string) *CreateSiteDeliveryTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSiteDeliveryTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
