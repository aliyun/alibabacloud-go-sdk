// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemberResponseBody
	GetCode() *string
	SetData(v *ListMemberResponseBodyData) *ListMemberResponseBody
	GetData() *ListMemberResponseBodyData
	SetHttpStatusCode(v int32) *ListMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemberResponseBody
	GetSuccess() *bool
}

type ListMemberResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemberResponseBody) GetData() *ListMemberResponseBodyData {
	return s.Data
}

func (s *ListMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemberResponseBody) SetCode(v string) *ListMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemberResponseBody) SetData(v *ListMemberResponseBodyData) *ListMemberResponseBody {
	s.Data = v
	return s
}

func (s *ListMemberResponseBody) SetHttpStatusCode(v int32) *ListMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemberResponseBody) SetMessage(v string) *ListMemberResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemberResponseBody) SetRequestId(v string) *ListMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemberResponseBody) SetSuccess(v bool) *ListMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemberResponseBodyData struct {
	Num      *int32                                `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMemberResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListMemberResponseBodyData) GetPageData() []*ListMemberResponseBodyDataPageData {
	return s.PageData
}

func (s *ListMemberResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListMemberResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMemberResponseBodyData) SetNum(v int32) *ListMemberResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMemberResponseBodyData) SetPageData(v []*ListMemberResponseBodyDataPageData) *ListMemberResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMemberResponseBodyData) SetSize(v int32) *ListMemberResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMemberResponseBodyData) SetTotal(v int32) *ListMemberResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMemberResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemberResponseBodyDataPageData struct {
	AdminName             *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	AuthorizedCount       *int64  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	AuthorizedDeviceCount *int32  `json:"AuthorizedDeviceCount,omitempty" xml:"AuthorizedDeviceCount,omitempty"`
	Contactor             *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	MemberId              *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Telephony             *string `json:"Telephony,omitempty" xml:"Telephony,omitempty"`
	Uid                   *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListMemberResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListMemberResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMemberResponseBodyDataPageData) GetAdminName() *string {
	return s.AdminName
}

func (s *ListMemberResponseBodyDataPageData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *ListMemberResponseBodyDataPageData) GetAuthorizedDeviceCount() *int32 {
	return s.AuthorizedDeviceCount
}

func (s *ListMemberResponseBodyDataPageData) GetContactor() *string {
	return s.Contactor
}

func (s *ListMemberResponseBodyDataPageData) GetMemberId() *string {
	return s.MemberId
}

func (s *ListMemberResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListMemberResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListMemberResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListMemberResponseBodyDataPageData) GetTelephony() *string {
	return s.Telephony
}

func (s *ListMemberResponseBodyDataPageData) GetUid() *string {
	return s.Uid
}

func (s *ListMemberResponseBodyDataPageData) SetAdminName(v string) *ListMemberResponseBodyDataPageData {
	s.AdminName = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetAuthorizedCount(v int64) *ListMemberResponseBodyDataPageData {
	s.AuthorizedCount = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetAuthorizedDeviceCount(v int32) *ListMemberResponseBodyDataPageData {
	s.AuthorizedDeviceCount = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetContactor(v string) *ListMemberResponseBodyDataPageData {
	s.Contactor = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetMemberId(v string) *ListMemberResponseBodyDataPageData {
	s.MemberId = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetName(v string) *ListMemberResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetRemark(v string) *ListMemberResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetStatus(v string) *ListMemberResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetTelephony(v string) *ListMemberResponseBodyDataPageData {
	s.Telephony = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) SetUid(v string) *ListMemberResponseBodyDataPageData {
	s.Uid = &v
	return s
}

func (s *ListMemberResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
