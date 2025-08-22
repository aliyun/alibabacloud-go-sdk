// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnRealTimeDeliveryProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListDcdnRealTimeDeliveryProjectResponseBodyContent) *ListDcdnRealTimeDeliveryProjectResponseBody
	GetContent() *ListDcdnRealTimeDeliveryProjectResponseBodyContent
	SetRequestId(v string) *ListDcdnRealTimeDeliveryProjectResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDcdnRealTimeDeliveryProjectResponseBody
	GetTotalCount() *int32
}

type ListDcdnRealTimeDeliveryProjectResponseBody struct {
	// The configuration results of the domain name.
	Content *ListDcdnRealTimeDeliveryProjectResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3EACD23C-F49F-4BF7-B9AD-C2CD3BA888C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) GetContent() *ListDcdnRealTimeDeliveryProjectResponseBodyContent {
	return s.Content
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetContent(v *ListDcdnRealTimeDeliveryProjectResponseBodyContent) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.Content = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetRequestId(v string) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) SetTotalCount(v int32) *ListDcdnRealTimeDeliveryProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDcdnRealTimeDeliveryProjectResponseBodyContent struct {
	Projects []*ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContent) GetProjects() []*ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	return s.Projects
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContent) SetProjects(v []*ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) *ListDcdnRealTimeDeliveryProjectResponseBodyContent {
	s.Projects = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of DCDN POPs
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
	// 1
	SamplingRate *float32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The type of log delivery. Only **SLS_POST*	- is supported.
	//
	// example:
	//
	// SLS_POST
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetFieldName() *string {
	return s.FieldName
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetSLSProject() *string {
	return s.SLSProject
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetSamplingRate() *float32 {
	return s.SamplingRate
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) GetType() *string {
	return s.Type
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetBusinessType(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.BusinessType = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetDataCenter(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.DataCenter = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetDomainName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.DomainName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetFieldName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.FieldName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetProjectName(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.ProjectName = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSLogStore(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSLogStore = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSProject(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSProject = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSLSRegion(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SLSRegion = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetSamplingRate(v float32) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.SamplingRate = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) SetType(v string) *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects {
	s.Type = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponseBodyContentProjects) Validate() error {
	return dara.Validate(s)
}
