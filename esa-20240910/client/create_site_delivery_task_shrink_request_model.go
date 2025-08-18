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
	// The log category. Valid values:
	//
	// 	- **dcdn_log_access_l1*	- (default): access logs.
	//
	// 	- **dcdn_log_er**: Edge Routine logs.
	//
	// 	- **dcdn_log_waf**: firewall logs.
	//
	// 	- **dcdn_log_ipa**: TCP/UDP proxy logs.
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
	// 	- oversea: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The destination of the delivery. Valid values:
	//
	// 	- sls: Alibaba Cloud Simple Log Service (SLS).
	//
	// 	- http: HTTP server.
	//
	// 	- aws3: Amazon Simple Storage Service (S3).
	//
	// 	- oss: Alibaba Cloud Object Storage Service (OSS).
	//
	// 	- kafka: Kafka.
	//
	// 	- aws3cmpt: S3-compatible storage service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The discard rate. Default value: 0.
	//
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The log fields, which are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_adress,ip_port
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
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12312312112***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The configurations for delivery to SLS.
	SlsDeliveryShrink *string `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty"`
	// The name of the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn-test-task
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
