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
	// The type of real-time log for Dynamic Route for CDN (DCDN). Valid values:
	//
	// - **dcdn_log_access_l1*	- (default): access log.
	//
	// - **dcdn_log_er**: edge function log.
	//
	// - **dcdn_log_waf**: WAF log.
	//
	// - **dcdn_log_ipa**: layer 4 acceleration log.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// - **cn**: Chinese mainland.
	//
	// - **sg**: Global (excluding Chinese mainland). Note that the value for this region is "sg".
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The delivery type. Valid values:
	//
	// - **sls**: Log Service.
	//
	// - **http**: HTTP service.
	//
	// - **aws3**: Amazon S3.
	//
	// - **oss**: Object Storage Service.
	//
	// - **kafka**: Kafka service.
	//
	// - **aws3cmpt**: Amazon S3-compatible service.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The discard rate.
	//
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// A comma-separated list of log fields to deliver.
	//
	// example:
	//
	// Client,UserAgent
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// The filter rules.
	//
	// example:
	//
	// []
	FilterRules *string `json:"FilterRules,omitempty" xml:"FilterRules,omitempty"`
	FilterVer   *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	RawRule     *string `json:"RawRule,omitempty" xml:"RawRule,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The delivery configuration.
	//
	// example:
	//
	// {\\"Region\\": \\"cn-hangzhou\\", \\"Endpoint\\": \\"https://***.oss-cn-hangzhou.aliyuncs.com\\", \\"BucketPath\\": \\"hjy-test002/online-logs\\"}
	SinkConfig interface{} `json:"SinkConfig,omitempty" xml:"SinkConfig,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// test.***.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The status of the task. Valid values:
	//
	// - **online**: The task is delivering logs.
	//
	// - **offline**: The task is paused.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
