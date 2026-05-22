// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserDeliveryTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateUserDeliveryTaskShrinkRequest
	GetBusinessType() *string
	SetDataCenter(v string) *CreateUserDeliveryTaskShrinkRequest
	GetDataCenter() *string
	SetDeliveryType(v string) *CreateUserDeliveryTaskShrinkRequest
	GetDeliveryType() *string
	SetDetails(v string) *CreateUserDeliveryTaskShrinkRequest
	GetDetails() *string
	SetDiscardRate(v float32) *CreateUserDeliveryTaskShrinkRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *CreateUserDeliveryTaskShrinkRequest
	GetFieldName() *string
	SetFilterVer(v string) *CreateUserDeliveryTaskShrinkRequest
	GetFilterVer() *string
	SetHttpDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest
	GetHttpDeliveryShrink() *string
	SetKafkaDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest
	GetKafkaDeliveryShrink() *string
	SetOssDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest
	GetOssDeliveryShrink() *string
	SetS3DeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest
	GetS3DeliveryShrink() *string
	SetSlsDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest
	GetSlsDeliveryShrink() *string
	SetTaskName(v string) *CreateUserDeliveryTaskShrinkRequest
	GetTaskName() *string
}

type CreateUserDeliveryTaskShrinkRequest struct {
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	DataCenter   *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// This parameter is required.
	DeliveryType *string  `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	Details      *string  `json:"Details,omitempty" xml:"Details,omitempty"`
	DiscardRate  *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// This parameter is required.
	FieldName           *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer           *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	HttpDeliveryShrink  *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	OssDeliveryShrink   *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	S3DeliveryShrink    *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	SlsDeliveryShrink   *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetDetails() *string {
	return s.Details
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetHttpDeliveryShrink() *string {
	return s.HttpDeliveryShrink
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetKafkaDeliveryShrink() *string {
	return s.KafkaDeliveryShrink
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetOssDeliveryShrink() *string {
	return s.OssDeliveryShrink
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetS3DeliveryShrink() *string {
	return s.S3DeliveryShrink
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetSlsDeliveryShrink() *string {
	return s.SlsDeliveryShrink
}

func (s *CreateUserDeliveryTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetBusinessType(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDataCenter(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDeliveryType(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDetails(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.Details = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetDiscardRate(v float32) *CreateUserDeliveryTaskShrinkRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetFieldName(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.FieldName = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetFilterVer(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.FilterVer = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetHttpDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.HttpDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetKafkaDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.KafkaDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetOssDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.OssDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetS3DeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.S3DeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetSlsDeliveryShrink(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.SlsDeliveryShrink = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) SetTaskName(v string) *CreateUserDeliveryTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUserDeliveryTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
