// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppServicesPageResponseBody
	GetCode() *string
	SetData(v []*ListAppServicesPageResponseBodyData) *ListAppServicesPageResponseBody
	GetData() []*ListAppServicesPageResponseBodyData
	SetErrorCode(v string) *ListAppServicesPageResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppServicesPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppServicesPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppServicesPageResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListAppServicesPageResponseBody
	GetTraceId() *string
}

type ListAppServicesPageResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: client error.
	//
	// - **5xx**: server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The service list.
	Data []*ListAppServicesPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// A message that describes the outcome of the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2583E089-99C2-562E-8B7E-73512136****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the details of a request.
	//
	// example:
	//
	// 0be3e0c816394483660457498e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListAppServicesPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppServicesPageResponseBody) GetData() []*ListAppServicesPageResponseBodyData {
	return s.Data
}

func (s *ListAppServicesPageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppServicesPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppServicesPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppServicesPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppServicesPageResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListAppServicesPageResponseBody) SetCode(v string) *ListAppServicesPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetData(v []*ListAppServicesPageResponseBodyData) *ListAppServicesPageResponseBody {
	s.Data = v
	return s
}

func (s *ListAppServicesPageResponseBody) SetErrorCode(v string) *ListAppServicesPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetMessage(v string) *ListAppServicesPageResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetRequestId(v string) *ListAppServicesPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetSuccess(v bool) *ListAppServicesPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetTraceId(v string) *ListAppServicesPageResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListAppServicesPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppServicesPageResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value must be in the range of 0 to 9999.
	//
	// example:
	//
	// 9999
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned results.
	Result []*ListAppServicesPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalSize *string `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAppServicesPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBodyData) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListAppServicesPageResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAppServicesPageResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAppServicesPageResponseBodyData) GetResult() []*ListAppServicesPageResponseBodyDataResult {
	return s.Result
}

func (s *ListAppServicesPageResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *ListAppServicesPageResponseBodyData) SetCurrentPage(v string) *ListAppServicesPageResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetPageNumber(v string) *ListAppServicesPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetPageSize(v string) *ListAppServicesPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetResult(v []*ListAppServicesPageResponseBodyDataResult) *ListAppServicesPageResponseBodyData {
	s.Result = v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetTotalSize(v string) *ListAppServicesPageResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppServicesPageResponseBodyDataResult struct {
	// The application ID.
	//
	// example:
	//
	// hc4fs1****@98314c8790b****
	EdasAppId *string `json:"EdasAppId,omitempty" xml:"EdasAppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// cn-zhangjiakou-micro-service-******
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	// The service group. This value is user-defined.
	//
	// example:
	//
	// springCloud
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceNum *int64 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The service name.
	//
	// example:
	//
	// edas.service.provider
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service version. This value is user-defined.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppServicesPageResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBodyDataResult) GetEdasAppId() *string {
	return s.EdasAppId
}

func (s *ListAppServicesPageResponseBodyDataResult) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *ListAppServicesPageResponseBodyDataResult) GetGroup() *string {
	return s.Group
}

func (s *ListAppServicesPageResponseBodyDataResult) GetInstanceNum() *int64 {
	return s.InstanceNum
}

func (s *ListAppServicesPageResponseBodyDataResult) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAppServicesPageResponseBodyDataResult) GetVersion() *string {
	return s.Version
}

func (s *ListAppServicesPageResponseBodyDataResult) SetEdasAppId(v string) *ListAppServicesPageResponseBodyDataResult {
	s.EdasAppId = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetEdasAppName(v string) *ListAppServicesPageResponseBodyDataResult {
	s.EdasAppName = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetGroup(v string) *ListAppServicesPageResponseBodyDataResult {
	s.Group = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetInstanceNum(v int64) *ListAppServicesPageResponseBodyDataResult {
	s.InstanceNum = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetServiceName(v string) *ListAppServicesPageResponseBodyDataResult {
	s.ServiceName = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetVersion(v string) *ListAppServicesPageResponseBodyDataResult {
	s.Version = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
