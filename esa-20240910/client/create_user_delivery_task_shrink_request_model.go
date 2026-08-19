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
	// The real-time log type. Valid values:
	//
	// - **dcdn_log_er_pod**: edge container logs.
	//
	// - **dcdn_log_dns**: edge DNS logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_er_pod
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// - **cn**: the Chinese mainland.
	//
	// - **sg**: global (excluding the Chinese mainland).
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The delivery type. Valid values:
	//
	// - **sls**: Alibaba Cloud Simple Log Service.
	//
	// - **http**: HTTP service.
	//
	// - **aws3**: Amazon S3 service.
	//
	// - **oss**: Alibaba Cloud Object Storage Service.
	//
	// - **kafka**: Kafka service.
	//
	// - **aws3cmpt**: Amazon S3-compatible service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The list of Edge Routine (ER) pods to configure.
	//
	// example:
	//
	// xxx,xxx
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The discard rate. Default value: 0.
	//
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The fields to deliver, separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ClientIP,ClientRequestURI,EdgeResponseStatusCode
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The version of the filter rule.
	//
	// > This parameter is used for backward compatibility with legacy filter rules. The default value is v1. New tasks use v2.
	//
	// example:
	//
	// v2
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The HTTP delivery configuration parameters.
	HttpDeliveryShrink *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	// The Kafka delivery configuration parameters.
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	// The OSS delivery configuration parameters.
	OssDeliveryShrink *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	// The S3 or S3-compatible delivery configuration parameters.
	S3DeliveryShrink *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	// The SLS delivery configuration.
	SlsDeliveryShrink *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
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
