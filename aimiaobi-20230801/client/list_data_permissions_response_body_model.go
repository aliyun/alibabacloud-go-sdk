// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataPermissionsResponseBody
	GetCode() *string
	SetData(v []*ListDataPermissionsResponseBodyData) *ListDataPermissionsResponseBody
	GetData() []*ListDataPermissionsResponseBodyData
	SetHttpStatusCode(v int32) *ListDataPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataPermissionsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListDataPermissionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataPermissionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataPermissionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataPermissionsResponseBody
	GetTotalCount() *int32
}

type ListDataPermissionsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDataPermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataPermissionsResponseBody) GetData() []*ListDataPermissionsResponseBodyData {
	return s.Data
}

func (s *ListDataPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataPermissionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataPermissionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataPermissionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataPermissionsResponseBody) SetCode(v string) *ListDataPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetData(v []*ListDataPermissionsResponseBodyData) *ListDataPermissionsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataPermissionsResponseBody) SetHttpStatusCode(v int32) *ListDataPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetMessage(v string) *ListDataPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetPageNumber(v int32) *ListDataPermissionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetPageSize(v int32) *ListDataPermissionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetRequestId(v string) *ListDataPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetSuccess(v bool) *ListDataPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataPermissionsResponseBody) SetTotalCount(v int32) *ListDataPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataPermissionsResponseBody) Validate() error {
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

type ListDataPermissionsResponseBodyData struct {
	// example:
	//
	// 2024-11-12 21:46:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// SystemSearch.QuarkCommonNews
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// xxx
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// read
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// xxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListDataPermissionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataPermissionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataPermissionsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDataPermissionsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataPermissionsResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ListDataPermissionsResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *ListDataPermissionsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDataPermissionsResponseBodyData) GetPermission() *string {
	return s.Permission
}

func (s *ListDataPermissionsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListDataPermissionsResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *ListDataPermissionsResponseBodyData) SetCreateTime(v string) *ListDataPermissionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetCreateUser(v string) *ListDataPermissionsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetDataId(v string) *ListDataPermissionsResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetDataType(v string) *ListDataPermissionsResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetId(v int64) *ListDataPermissionsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetPermission(v string) *ListDataPermissionsResponseBodyData {
	s.Permission = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetUserId(v string) *ListDataPermissionsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) SetUsername(v string) *ListDataPermissionsResponseBodyData {
	s.Username = &v
	return s
}

func (s *ListDataPermissionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
