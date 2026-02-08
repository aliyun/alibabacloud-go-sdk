// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContactWhiteListResponseBody
	GetCode() *string
	SetContactWhitelistList(v *GetContactWhiteListResponseBodyContactWhitelistList) *GetContactWhiteListResponseBody
	GetContactWhitelistList() *GetContactWhiteListResponseBodyContactWhitelistList
	SetHttpStatusCode(v int32) *GetContactWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetContactWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContactWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContactWhiteListResponseBody
	GetSuccess() *bool
}

type GetContactWhiteListResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Contact list
	//
	// example:
	//
	// {}
	ContactWhitelistList *GetContactWhiteListResponseBodyContactWhitelistList `json:"ContactWhitelistList,omitempty" xml:"ContactWhitelistList,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContactWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContactWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContactWhiteListResponseBody) GetContactWhitelistList() *GetContactWhiteListResponseBodyContactWhitelistList {
	return s.ContactWhitelistList
}

func (s *GetContactWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetContactWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContactWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContactWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContactWhiteListResponseBody) SetCode(v string) *GetContactWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *GetContactWhiteListResponseBody) SetContactWhitelistList(v *GetContactWhiteListResponseBodyContactWhitelistList) *GetContactWhiteListResponseBody {
	s.ContactWhitelistList = v
	return s
}

func (s *GetContactWhiteListResponseBody) SetHttpStatusCode(v int32) *GetContactWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContactWhiteListResponseBody) SetMessage(v string) *GetContactWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *GetContactWhiteListResponseBody) SetRequestId(v string) *GetContactWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactWhiteListResponseBody) SetSuccess(v bool) *GetContactWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *GetContactWhiteListResponseBody) Validate() error {
	if s.ContactWhitelistList != nil {
		if err := s.ContactWhitelistList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContactWhiteListResponseBodyContactWhitelistList struct {
	// Contact list
	//
	// example:
	//
	// {}
	List []*GetContactWhiteListResponseBodyContactWhitelistListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetContactWhiteListResponseBodyContactWhitelistList) String() string {
	return dara.Prettify(s)
}

func (s GetContactWhiteListResponseBodyContactWhitelistList) GoString() string {
	return s.String()
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) GetList() []*GetContactWhiteListResponseBodyContactWhitelistListList {
	return s.List
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) SetList(v []*GetContactWhiteListResponseBodyContactWhitelistListList) *GetContactWhiteListResponseBodyContactWhitelistList {
	s.List = v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) SetPageNumber(v int32) *GetContactWhiteListResponseBodyContactWhitelistList {
	s.PageNumber = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) SetPageSize(v int32) *GetContactWhiteListResponseBodyContactWhitelistList {
	s.PageSize = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) SetTotalCount(v int32) *GetContactWhiteListResponseBodyContactWhitelistList {
	s.TotalCount = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistList) Validate() error {
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

type GetContactWhiteListResponseBodyContactWhitelistListList struct {
	// Whitelist ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	ContactWhiteListId *string `json:"ContactWhiteListId,omitempty" xml:"ContactWhiteListId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1640174411848
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Creator
	//
	// example:
	//
	// 测试123
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Instance ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Name
	//
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operator
	//
	// example:
	//
	// xxx
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Phone number
	//
	// example:
	//
	// 13959999999
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Remark
	//
	// example:
	//
	// xxxx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s GetContactWhiteListResponseBodyContactWhitelistListList) String() string {
	return dara.Prettify(s)
}

func (s GetContactWhiteListResponseBodyContactWhitelistListList) GoString() string {
	return s.String()
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetContactWhiteListId() *string {
	return s.ContactWhiteListId
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetCreator() *string {
	return s.Creator
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetName() *string {
	return s.Name
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetOperator() *string {
	return s.Operator
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) GetRemark() *string {
	return s.Remark
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetContactWhiteListId(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.ContactWhiteListId = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetCreationTime(v int64) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.CreationTime = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetCreator(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.Creator = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetInstanceId(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.InstanceId = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetName(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.Name = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetOperator(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.Operator = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetPhoneNumber(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.PhoneNumber = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) SetRemark(v string) *GetContactWhiteListResponseBodyContactWhitelistListList {
	s.Remark = &v
	return s
}

func (s *GetContactWhiteListResponseBodyContactWhitelistListList) Validate() error {
	return dara.Validate(s)
}
