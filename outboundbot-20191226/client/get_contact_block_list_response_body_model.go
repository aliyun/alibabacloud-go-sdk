// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactBlockListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContactBlockListResponseBody
	GetCode() *string
	SetContactBlocklistList(v *GetContactBlockListResponseBodyContactBlocklistList) *GetContactBlockListResponseBody
	GetContactBlocklistList() *GetContactBlockListResponseBodyContactBlocklistList
	SetHttpStatusCode(v int32) *GetContactBlockListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetContactBlockListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContactBlockListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContactBlockListResponseBody
	GetSuccess() *bool
}

type GetContactBlockListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of blocked contacts.
	//
	// example:
	//
	// {}
	ContactBlocklistList *GetContactBlockListResponseBodyContactBlocklistList `json:"ContactBlocklistList,omitempty" xml:"ContactBlocklistList,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContactBlockListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContactBlockListResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactBlockListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContactBlockListResponseBody) GetContactBlocklistList() *GetContactBlockListResponseBodyContactBlocklistList {
	return s.ContactBlocklistList
}

func (s *GetContactBlockListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetContactBlockListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContactBlockListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContactBlockListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContactBlockListResponseBody) SetCode(v string) *GetContactBlockListResponseBody {
	s.Code = &v
	return s
}

func (s *GetContactBlockListResponseBody) SetContactBlocklistList(v *GetContactBlockListResponseBodyContactBlocklistList) *GetContactBlockListResponseBody {
	s.ContactBlocklistList = v
	return s
}

func (s *GetContactBlockListResponseBody) SetHttpStatusCode(v int32) *GetContactBlockListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContactBlockListResponseBody) SetMessage(v string) *GetContactBlockListResponseBody {
	s.Message = &v
	return s
}

func (s *GetContactBlockListResponseBody) SetRequestId(v string) *GetContactBlockListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactBlockListResponseBody) SetSuccess(v bool) *GetContactBlockListResponseBody {
	s.Success = &v
	return s
}

func (s *GetContactBlockListResponseBody) Validate() error {
	if s.ContactBlocklistList != nil {
		if err := s.ContactBlocklistList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContactBlockListResponseBodyContactBlocklistList struct {
	// The data array.
	//
	// example:
	//
	// []
	List []*GetContactBlockListResponseBodyContactBlocklistListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetContactBlockListResponseBodyContactBlocklistList) String() string {
	return dara.Prettify(s)
}

func (s GetContactBlockListResponseBodyContactBlocklistList) GoString() string {
	return s.String()
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) GetList() []*GetContactBlockListResponseBodyContactBlocklistListList {
	return s.List
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) SetList(v []*GetContactBlockListResponseBodyContactBlocklistListList) *GetContactBlockListResponseBodyContactBlocklistList {
	s.List = v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) SetPageNumber(v int32) *GetContactBlockListResponseBodyContactBlocklistList {
	s.PageNumber = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) SetPageSize(v int32) *GetContactBlockListResponseBodyContactBlocklistList {
	s.PageSize = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) SetTotalCount(v int32) *GetContactBlockListResponseBodyContactBlocklistList {
	s.TotalCount = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistList) Validate() error {
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

type GetContactBlockListResponseBodyContactBlocklistListList struct {
	// The unique key of the entry in the do-not-call list.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	ContactBlockListId *string `json:"ContactBlockListId,omitempty" xml:"ContactBlockListId,omitempty"`
	// The time when the entry was created.
	//
	// example:
	//
	// 1640077685465
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The user who created the entry.
	//
	// example:
	//
	// xxx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operator.
	//
	// example:
	//
	// xxx
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 1388888888
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The remarks.
	//
	// example:
	//
	// xxxx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s GetContactBlockListResponseBodyContactBlocklistListList) String() string {
	return dara.Prettify(s)
}

func (s GetContactBlockListResponseBodyContactBlocklistListList) GoString() string {
	return s.String()
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetContactBlockListId() *string {
	return s.ContactBlockListId
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetCreator() *string {
	return s.Creator
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetName() *string {
	return s.Name
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetOperator() *string {
	return s.Operator
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) GetRemark() *string {
	return s.Remark
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetContactBlockListId(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.ContactBlockListId = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetCreationTime(v int64) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.CreationTime = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetCreator(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.Creator = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetInstanceId(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.InstanceId = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetName(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.Name = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetOperator(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.Operator = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetPhoneNumber(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.PhoneNumber = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) SetRemark(v string) *GetContactBlockListResponseBodyContactBlocklistListList {
	s.Remark = &v
	return s
}

func (s *GetContactBlockListResponseBodyContactBlocklistListList) Validate() error {
	return dara.Validate(s)
}
