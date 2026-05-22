// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetUserDeliveryTaskResponseBody
	GetBusinessType() *string
	SetDataCenter(v string) *GetUserDeliveryTaskResponseBody
	GetDataCenter() *string
	SetDeliveryType(v string) *GetUserDeliveryTaskResponseBody
	GetDeliveryType() *string
	SetDetails(v string) *GetUserDeliveryTaskResponseBody
	GetDetails() *string
	SetDiscardRate(v float32) *GetUserDeliveryTaskResponseBody
	GetDiscardRate() *float32
	SetFieldList(v string) *GetUserDeliveryTaskResponseBody
	GetFieldList() *string
	SetFilterRules(v string) *GetUserDeliveryTaskResponseBody
	GetFilterRules() *string
	SetFilterVer(v string) *GetUserDeliveryTaskResponseBody
	GetFilterVer() *string
	SetRawRule(v string) *GetUserDeliveryTaskResponseBody
	GetRawRule() *string
	SetRequestId(v string) *GetUserDeliveryTaskResponseBody
	GetRequestId() *string
	SetSinkConfig(v interface{}) *GetUserDeliveryTaskResponseBody
	GetSinkConfig() interface{}
	SetStatus(v string) *GetUserDeliveryTaskResponseBody
	GetStatus() *string
	SetTaskName(v string) *GetUserDeliveryTaskResponseBody
	GetTaskName() *string
}

type GetUserDeliveryTaskResponseBody struct {
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
	// example:
	//
	// dcdn_log_er
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
	// example:
	//
	// oss
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	Details      *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The discard rate.
	//
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The fields.
	//
	// example:
	//
	// ClientRequestID,ClientRequestHost
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// The filtering rules.
	//
	// example:
	//
	// [{"ClientSSLProtocol": {"equals": ["TLSv1.3"]}}]
	FilterRules *string `json:"FilterRules,omitempty" xml:"FilterRules,omitempty"`
	FilterVer   *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	RawRule     *string `json:"RawRule,omitempty" xml:"RawRule,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7072132a-bd3c-46a6-9e81-aba3e0e3f861
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The delivery configuration.
	//
	// example:
	//
	// {\\"Project\\": \\"er-online-hjy-pro\\", \\"Logstore\\": \\"er-online-hjy-log\\", \\"Region\\": \\"cn-hangzhou\\", \\"Endpoint\\": \\"cn-hangzhou.log.aliyuncs.com\\", \\"Aliuid\\": \\"1077912128805410\\"}
	SinkConfig interface{} `json:"SinkConfig,omitempty" xml:"SinkConfig,omitempty"`
	// The status of the delivery task.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the delivery task.
	//
	// example:
	//
	// testoss11
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetUserDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskResponseBody) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUserDeliveryTaskResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *GetUserDeliveryTaskResponseBody) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *GetUserDeliveryTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetUserDeliveryTaskResponseBody) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *GetUserDeliveryTaskResponseBody) GetFieldList() *string {
	return s.FieldList
}

func (s *GetUserDeliveryTaskResponseBody) GetFilterRules() *string {
	return s.FilterRules
}

func (s *GetUserDeliveryTaskResponseBody) GetFilterVer() *string {
	return s.FilterVer
}

func (s *GetUserDeliveryTaskResponseBody) GetRawRule() *string {
	return s.RawRule
}

func (s *GetUserDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserDeliveryTaskResponseBody) GetSinkConfig() interface{} {
	return s.SinkConfig
}

func (s *GetUserDeliveryTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUserDeliveryTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *GetUserDeliveryTaskResponseBody) SetBusinessType(v string) *GetUserDeliveryTaskResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDataCenter(v string) *GetUserDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDeliveryType(v string) *GetUserDeliveryTaskResponseBody {
	s.DeliveryType = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDetails(v string) *GetUserDeliveryTaskResponseBody {
	s.Details = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetDiscardRate(v float32) *GetUserDeliveryTaskResponseBody {
	s.DiscardRate = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetFieldList(v string) *GetUserDeliveryTaskResponseBody {
	s.FieldList = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetFilterRules(v string) *GetUserDeliveryTaskResponseBody {
	s.FilterRules = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetFilterVer(v string) *GetUserDeliveryTaskResponseBody {
	s.FilterVer = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetRawRule(v string) *GetUserDeliveryTaskResponseBody {
	s.RawRule = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetRequestId(v string) *GetUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetSinkConfig(v interface{}) *GetUserDeliveryTaskResponseBody {
	s.SinkConfig = v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetStatus(v string) *GetUserDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) SetTaskName(v string) *GetUserDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetUserDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
