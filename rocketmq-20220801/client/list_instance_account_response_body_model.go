// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstanceAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListInstanceAccountResponseBody
	GetCode() *string
	SetData(v *ListInstanceAccountResponseBodyData) *ListInstanceAccountResponseBody
	GetData() *ListInstanceAccountResponseBodyData
	SetDynamicCode(v string) *ListInstanceAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListInstanceAccountResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListInstanceAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceAccountResponseBody
	GetSuccess() *bool
}

type ListInstanceAccountResponseBody struct {
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
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListInstanceAccountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// C115601B-8736-5BBF-AC99-7FEAE1245A80
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstanceAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceAccountResponseBody) GetData() *ListInstanceAccountResponseBodyData {
	return s.Data
}

func (s *ListInstanceAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListInstanceAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListInstanceAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *ListInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetCode(v string) *ListInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetData(v *ListInstanceAccountResponseBodyData) *ListInstanceAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAccountResponseBody) SetDynamicCode(v string) *ListInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetDynamicMessage(v string) *ListInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetHttpStatusCode(v int32) *ListInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetMessage(v string) *ListInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetRequestId(v string) *ListInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetSuccess(v bool) *ListInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceAccountResponseBodyData struct {
	// The pagination information.
	List []*ListInstanceAccountResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	// 24
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBodyData) GetList() []*ListInstanceAccountResponseBodyDataList {
	return s.List
}

func (s *ListInstanceAccountResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceAccountResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceAccountResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceAccountResponseBodyData) SetList(v []*ListInstanceAccountResponseBodyDataList) *ListInstanceAccountResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetPageNumber(v int64) *ListInstanceAccountResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetPageSize(v int64) *ListInstanceAccountResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetTotalCount(v int64) *ListInstanceAccountResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListInstanceAccountResponseBodyDataList struct {
	// The status of the account.
	//
	// Valid values:
	//
	//   - DISABLE
	//
	//   - ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The account type.
	//
	//   - CUSTOMER
	//
	//   - DEFAULT
	//
	// example:
	//
	// CUSTOMER
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAccountResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAccountResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBodyDataList) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *ListInstanceAccountResponseBodyDataList) GetAccountType() *string {
	return s.AccountType
}

func (s *ListInstanceAccountResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceAccountResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceAccountResponseBodyDataList) GetUsername() *string {
	return s.Username
}

func (s *ListInstanceAccountResponseBodyDataList) SetAccountStatus(v string) *ListInstanceAccountResponseBodyDataList {
	s.AccountStatus = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetAccountType(v string) *ListInstanceAccountResponseBodyDataList {
	s.AccountType = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetInstanceId(v string) *ListInstanceAccountResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetRegionId(v string) *ListInstanceAccountResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetUsername(v string) *ListInstanceAccountResponseBodyDataList {
	s.Username = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
