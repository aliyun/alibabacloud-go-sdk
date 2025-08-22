// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSLSRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody
	GetContent() *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent
	SetRequestId(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DescribeDcdnSLSRealtimeLogDeliveryResponseBody struct {
	// The configuration results of the domain name.
	Content *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) GetContent() *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) SetContent(v *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of Dynamic Content Delivery Network (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The region from which logs were collected.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The domain names from which logs were collected. You can specify one or more domain names. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the field. For more information about fields in real-time log entries, see [Fields in a real-time log](https://help.aliyun.com/document_detail/324199.html).
	//
	// example:
	//
	// field1,field2
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// example-cn
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// The name of the log file.
	//
	// example:
	//
	// example-cn
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region to which logs were delivered.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 1.0
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The status of real-time logs.
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of log delivery. Only **SLS_POST*	- is supported.
	//
	// example:
	//
	// SLS_POST
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetSLSProject() *string {
	return s.SLSProject
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetBusinessType(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDataCenter(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.DataCenter = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetDomainName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetFieldName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.FieldName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetProjectName(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.ProjectName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSLogStore(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSLogStore = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSProject(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSProject = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSLSRegion(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SLSRegion = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetSamplingRate(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.SamplingRate = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetStatus(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.Status = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) SetType(v string) *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent {
	s.Type = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
