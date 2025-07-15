// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRamUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRamUsersResponseBody
	GetCode() *string
	SetData(v *ListRamUsersResponseBodyData) *ListRamUsersResponseBody
	GetData() *ListRamUsersResponseBodyData
	SetHttpStatusCode(v int32) *ListRamUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRamUsersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListRamUsersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListRamUsersResponseBody
	GetRequestId() *string
}

type ListRamUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListRamUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 24753D71-C91D-1A38-A8AD-372BF12453F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRamUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRamUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRamUsersResponseBody) GetData() *ListRamUsersResponseBodyData {
	return s.Data
}

func (s *ListRamUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRamUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRamUsersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListRamUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRamUsersResponseBody) SetCode(v string) *ListRamUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListRamUsersResponseBody) SetData(v *ListRamUsersResponseBodyData) *ListRamUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListRamUsersResponseBody) SetHttpStatusCode(v int32) *ListRamUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRamUsersResponseBody) SetMessage(v string) *ListRamUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListRamUsersResponseBody) SetParams(v []*string) *ListRamUsersResponseBody {
	s.Params = v
	return s
}

func (s *ListRamUsersResponseBody) SetRequestId(v string) *ListRamUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRamUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRamUsersResponseBodyData struct {
	List []*ListRamUsersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRamUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRamUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBodyData) GetList() []*ListRamUsersResponseBodyDataList {
	return s.List
}

func (s *ListRamUsersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRamUsersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRamUsersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRamUsersResponseBodyData) SetList(v []*ListRamUsersResponseBodyDataList) *ListRamUsersResponseBodyData {
	s.List = v
	return s
}

func (s *ListRamUsersResponseBodyData) SetPageNumber(v int32) *ListRamUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRamUsersResponseBodyData) SetPageSize(v int32) *ListRamUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRamUsersResponseBodyData) SetTotalCount(v int32) *ListRamUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRamUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListRamUsersResponseBodyDataList struct {
	// example:
	//
	// 15772400000****
	AliyunUid   *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// false
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// example:
	//
	// 28036411123456****
	RamId *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
}

func (s ListRamUsersResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListRamUsersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBodyDataList) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *ListRamUsersResponseBodyDataList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListRamUsersResponseBodyDataList) GetEmail() *string {
	return s.Email
}

func (s *ListRamUsersResponseBodyDataList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListRamUsersResponseBodyDataList) GetMobile() *string {
	return s.Mobile
}

func (s *ListRamUsersResponseBodyDataList) GetPrimary() *bool {
	return s.Primary
}

func (s *ListRamUsersResponseBodyDataList) GetRamId() *string {
	return s.RamId
}

func (s *ListRamUsersResponseBodyDataList) SetAliyunUid(v int64) *ListRamUsersResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetDisplayName(v string) *ListRamUsersResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetEmail(v string) *ListRamUsersResponseBodyDataList {
	s.Email = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetLoginName(v string) *ListRamUsersResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetMobile(v string) *ListRamUsersResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetPrimary(v bool) *ListRamUsersResponseBodyDataList {
	s.Primary = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetRamId(v string) *ListRamUsersResponseBodyDataList {
	s.RamId = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
