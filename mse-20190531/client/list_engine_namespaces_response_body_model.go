// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListEngineNamespacesResponseBodyData) *ListEngineNamespacesResponseBody
	GetData() []*ListEngineNamespacesResponseBodyData
	SetErrorCode(v string) *ListEngineNamespacesResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListEngineNamespacesResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListEngineNamespacesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListEngineNamespacesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEngineNamespacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEngineNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEngineNamespacesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListEngineNamespacesResponseBody
	GetTotalCount() *int32
}

type ListEngineNamespacesResponseBody struct {
	// The details of the data.
	Data []*ListEngineNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
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
	// 062D13C5-DEA3-4921-8918-C49A0F1B****
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

func (s ListEngineNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEngineNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponseBody) GetData() []*ListEngineNamespacesResponseBodyData {
	return s.Data
}

func (s *ListEngineNamespacesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEngineNamespacesResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListEngineNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEngineNamespacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEngineNamespacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEngineNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEngineNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEngineNamespacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEngineNamespacesResponseBody) SetData(v []*ListEngineNamespacesResponseBodyData) *ListEngineNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetErrorCode(v string) *ListEngineNamespacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetHttpCode(v string) *ListEngineNamespacesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetMessage(v string) *ListEngineNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetPageNumber(v int32) *ListEngineNamespacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetPageSize(v int32) *ListEngineNamespacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetRequestId(v string) *ListEngineNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetSuccess(v bool) *ListEngineNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetTotalCount(v int32) *ListEngineNamespacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEngineNamespacesResponseBodyData struct {
	// The quota value.
	//
	// example:
	//
	// 1
	ConfigCount *int32 `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	// The namespace.
	//
	// example:
	//
	// DEFAULT
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// mytest
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// public
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	// The quota.
	//
	// example:
	//
	// 200
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The number of active services.
	//
	// example:
	//
	// 3
	ServiceCount *string `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
	// The source from which the namespace was created.
	//
	// example:
	//
	// mse
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the namespace. Valid values:
	//
	// 	- `0`: global configuration
	//
	// 	- `1`: default namespace
	//
	// 	- `2`: custom namespace
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEngineNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEngineNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponseBodyData) GetConfigCount() *int32 {
	return s.ConfigCount
}

func (s *ListEngineNamespacesResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEngineNamespacesResponseBodyData) GetNamespaceDesc() *string {
	return s.NamespaceDesc
}

func (s *ListEngineNamespacesResponseBodyData) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *ListEngineNamespacesResponseBodyData) GetQuota() *int32 {
	return s.Quota
}

func (s *ListEngineNamespacesResponseBodyData) GetServiceCount() *string {
	return s.ServiceCount
}

func (s *ListEngineNamespacesResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *ListEngineNamespacesResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListEngineNamespacesResponseBodyData) SetConfigCount(v int32) *ListEngineNamespacesResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespace(v string) *ListEngineNamespacesResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespaceDesc(v string) *ListEngineNamespacesResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespaceShowName(v string) *ListEngineNamespacesResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetQuota(v int32) *ListEngineNamespacesResponseBodyData {
	s.Quota = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetServiceCount(v string) *ListEngineNamespacesResponseBodyData {
	s.ServiceCount = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetSourceType(v string) *ListEngineNamespacesResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetType(v int32) *ListEngineNamespacesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
