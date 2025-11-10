// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetricMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMetricMetaResponseBody
	GetCode() *string
	SetData(v *ListMetricMetaResponseBodyData) *ListMetricMetaResponseBody
	GetData() *ListMetricMetaResponseBodyData
	SetDynamicCode(v string) *ListMetricMetaResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMetricMetaResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListMetricMetaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMetricMetaResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMetricMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetricMetaResponseBody
	GetSuccess() *bool
}

type ListMetricMetaResponseBody struct {
	// Error code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result
	Data *ListMetricMetaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// xxx
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8B459455-4A35-5796-BA9D-98EF1AB9A931
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMetricMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetricMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetricMetaResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMetricMetaResponseBody) GetData() *ListMetricMetaResponseBodyData {
	return s.Data
}

func (s *ListMetricMetaResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMetricMetaResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMetricMetaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMetricMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMetricMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetricMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetricMetaResponseBody) SetCode(v string) *ListMetricMetaResponseBody {
	s.Code = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetData(v *ListMetricMetaResponseBodyData) *ListMetricMetaResponseBody {
	s.Data = v
	return s
}

func (s *ListMetricMetaResponseBody) SetDynamicCode(v string) *ListMetricMetaResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetDynamicMessage(v string) *ListMetricMetaResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetHttpStatusCode(v int32) *ListMetricMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetMessage(v string) *ListMetricMetaResponseBody {
	s.Message = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetRequestId(v string) *ListMetricMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetricMetaResponseBody) SetSuccess(v bool) *ListMetricMetaResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetricMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetricMetaResponseBodyData struct {
	// Paged data
	List []*ListMetricMetaResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total record count
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetricMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetricMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetricMetaResponseBodyData) GetList() []*ListMetricMetaResponseBodyDataList {
	return s.List
}

func (s *ListMetricMetaResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListMetricMetaResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMetricMetaResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMetricMetaResponseBodyData) SetList(v []*ListMetricMetaResponseBodyDataList) *ListMetricMetaResponseBodyData {
	s.List = v
	return s
}

func (s *ListMetricMetaResponseBodyData) SetPageNumber(v int64) *ListMetricMetaResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMetricMetaResponseBodyData) SetPageSize(v int64) *ListMetricMetaResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMetricMetaResponseBodyData) SetTotalCount(v int64) *ListMetricMetaResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMetricMetaResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetricMetaResponseBodyDataList struct {
	// Monitoring item tag
	//
	// example:
	//
	// Bug
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Monitoring item description
	//
	// example:
	//
	// Using Serverless Devs to deploy the infrastructure of project:get-userinfo-v1-infrastructure-as-template-project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Monitoring item name
	//
	// example:
	//
	// SendMessageCountPerInstance
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
}

func (s ListMetricMetaResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListMetricMetaResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListMetricMetaResponseBodyDataList) GetCategory() *string {
	return s.Category
}

func (s *ListMetricMetaResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListMetricMetaResponseBodyDataList) GetMetricName() *string {
	return s.MetricName
}

func (s *ListMetricMetaResponseBodyDataList) SetCategory(v string) *ListMetricMetaResponseBodyDataList {
	s.Category = &v
	return s
}

func (s *ListMetricMetaResponseBodyDataList) SetDescription(v string) *ListMetricMetaResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListMetricMetaResponseBodyDataList) SetMetricName(v string) *ListMetricMetaResponseBodyDataList {
	s.MetricName = &v
	return s
}

func (s *ListMetricMetaResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
