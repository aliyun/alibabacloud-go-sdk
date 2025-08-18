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
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// 	- cn: the Chinese mainland.
	//
	// 	- sg: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The destination of the delivery. Valid values:
	//
	// 1.  sls: Alibaba Cloud SLS.
	//
	// 2.  http: HTTP server.
	//
	// 3.  aws3: Amazon S3.
	//
	// 4.  oss: Alibaba Cloud OSS.
	//
	// 5.  kafka: Kafka.
	//
	// 6.  aws3cmpt: S3-compatible storage service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	Details      *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The discard rate. Default value: 0.
	//
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The log field. If you specify multiple fields, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_address,ip_port
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The configurations for delivery to an HTTP server.
	HttpDeliveryShrink *string `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty"`
	// The configurations for delivery to Kafka.
	KafkaDeliveryShrink *string `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty"`
	// The configurations for delivery to OSS.
	OssDeliveryShrink *string `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty"`
	// The configurations for delivery to Amazon S3 or an S3-compatible service.
	S3DeliveryShrink *string `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty"`
	// The configurations for delivery to SLS.
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
