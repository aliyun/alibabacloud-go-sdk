// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstanceIpWhitelistResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstanceIpWhitelistResponseBody
	GetCode() *string
	SetData(v *ListInstanceIpWhitelistResponseBodyData) *ListInstanceIpWhitelistResponseBody
	GetData() *ListInstanceIpWhitelistResponseBodyData
	SetDynamicCode(v string) *ListInstanceIpWhitelistResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListInstanceIpWhitelistResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListInstanceIpWhitelistResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceIpWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceIpWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceIpWhitelistResponseBody
	GetSuccess() *bool
}

type ListInstanceIpWhitelistResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListInstanceIpWhitelistResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-611C63E05674
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstanceIpWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceIpWhitelistResponseBody) GetData() *ListInstanceIpWhitelistResponseBodyData {
	return s.Data
}

func (s *ListInstanceIpWhitelistResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListInstanceIpWhitelistResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListInstanceIpWhitelistResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceIpWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceIpWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *ListInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetCode(v string) *ListInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetData(v *ListInstanceIpWhitelistResponseBodyData) *ListInstanceIpWhitelistResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *ListInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *ListInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *ListInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetMessage(v string) *ListInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetRequestId(v string) *ListInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetSuccess(v bool) *ListInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceIpWhitelistResponseBodyData struct {
	// The pagination information.
	List []*string `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceIpWhitelistResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIpWhitelistResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponseBodyData) GetList() []*string {
	return s.List
}

func (s *ListInstanceIpWhitelistResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceIpWhitelistResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceIpWhitelistResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetList(v []*string) *ListInstanceIpWhitelistResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetPageNumber(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetPageSize(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetTotalCount(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) Validate() error {
	return dara.Validate(s)
}
