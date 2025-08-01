// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListEurekaServicesResponseBodyData) *ListEurekaServicesResponseBody
	GetData() []*ListEurekaServicesResponseBodyData
	SetErrorCode(v string) *ListEurekaServicesResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListEurekaServicesResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListEurekaServicesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListEurekaServicesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEurekaServicesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEurekaServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEurekaServicesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListEurekaServicesResponseBody
	GetTotalCount() *int32
}

type ListEurekaServicesResponseBody struct {
	// The details of the data.
	Data []*ListEurekaServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEurekaServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponseBody) GetData() []*ListEurekaServicesResponseBodyData {
	return s.Data
}

func (s *ListEurekaServicesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEurekaServicesResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListEurekaServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEurekaServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEurekaServicesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEurekaServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEurekaServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEurekaServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEurekaServicesResponseBody) SetData(v []*ListEurekaServicesResponseBodyData) *ListEurekaServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListEurekaServicesResponseBody) SetErrorCode(v string) *ListEurekaServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetHttpCode(v string) *ListEurekaServicesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetMessage(v string) *ListEurekaServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetPageNumber(v int32) *ListEurekaServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetPageSize(v int32) *ListEurekaServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetRequestId(v string) *ListEurekaServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetSuccess(v bool) *ListEurekaServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetTotalCount(v int32) *ListEurekaServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEurekaServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEurekaServicesResponseBodyData struct {
	// The details of the instance.
	InstancesId []*string `json:"InstancesId,omitempty" xml:"InstancesId,omitempty" type:"Repeated"`
	// The name of the service.
	//
	// example:
	//
	// CONTACTINFO
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of service providers. The value is in the following format: Number of healthy instances/Total number of instances.
	//
	// example:
	//
	// 1/1
	UpStatus *string `json:"UpStatus,omitempty" xml:"UpStatus,omitempty"`
}

func (s ListEurekaServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponseBodyData) GetInstancesId() []*string {
	return s.InstancesId
}

func (s *ListEurekaServicesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListEurekaServicesResponseBodyData) GetUpStatus() *string {
	return s.UpStatus
}

func (s *ListEurekaServicesResponseBodyData) SetInstancesId(v []*string) *ListEurekaServicesResponseBodyData {
	s.InstancesId = v
	return s
}

func (s *ListEurekaServicesResponseBodyData) SetName(v string) *ListEurekaServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListEurekaServicesResponseBodyData) SetUpStatus(v string) *ListEurekaServicesResponseBodyData {
	s.UpStatus = &v
	return s
}

func (s *ListEurekaServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
