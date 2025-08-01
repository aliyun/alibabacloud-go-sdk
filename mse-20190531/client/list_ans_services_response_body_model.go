// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAnsServicesResponseBodyData) *ListAnsServicesResponseBody
	GetData() []*ListAnsServicesResponseBodyData
	SetErrorCode(v string) *ListAnsServicesResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListAnsServicesResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListAnsServicesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListAnsServicesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAnsServicesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAnsServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnsServicesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAnsServicesResponseBody
	GetTotalCount() *int32
}

type ListAnsServicesResponseBody struct {
	// The details of the data.
	Data []*ListAnsServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page.
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
	// 52BA6DA6-A702-4362-A32F-DFF79655****
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
	// The total number of instances returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnsServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponseBody) GetData() []*ListAnsServicesResponseBodyData {
	return s.Data
}

func (s *ListAnsServicesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAnsServicesResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListAnsServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnsServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAnsServicesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnsServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnsServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnsServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAnsServicesResponseBody) SetData(v []*ListAnsServicesResponseBodyData) *ListAnsServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsServicesResponseBody) SetErrorCode(v string) *ListAnsServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetHttpCode(v string) *ListAnsServicesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetMessage(v string) *ListAnsServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetPageNumber(v int32) *ListAnsServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetPageSize(v int32) *ListAnsServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetRequestId(v string) *ListAnsServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetSuccess(v bool) *ListAnsServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetTotalCount(v int32) *ListAnsServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAnsServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAnsServicesResponseBodyData struct {
	// The total number of clusters.
	//
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The name of the contact group.
	//
	// example:
	//
	// name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The total number of instances with healthy heartbeats.
	//
	// example:
	//
	// 1
	HealthyInstanceCount *int32 `json:"HealthyInstanceCount,omitempty" xml:"HealthyInstanceCount,omitempty"`
	// The total number of instances that are used for the current service.
	//
	// example:
	//
	// 1
	IpCount *int32 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// name
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListAnsServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponseBodyData) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *ListAnsServicesResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAnsServicesResponseBodyData) GetHealthyInstanceCount() *int32 {
	return s.HealthyInstanceCount
}

func (s *ListAnsServicesResponseBodyData) GetIpCount() *int32 {
	return s.IpCount
}

func (s *ListAnsServicesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAnsServicesResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListAnsServicesResponseBodyData) SetClusterCount(v int32) *ListAnsServicesResponseBodyData {
	s.ClusterCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetGroupName(v string) *ListAnsServicesResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetHealthyInstanceCount(v int32) *ListAnsServicesResponseBodyData {
	s.HealthyInstanceCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetIpCount(v int32) *ListAnsServicesResponseBodyData {
	s.IpCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetName(v string) *ListAnsServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetSource(v string) *ListAnsServicesResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
