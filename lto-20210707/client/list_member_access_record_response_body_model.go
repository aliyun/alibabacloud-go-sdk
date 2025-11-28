// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAccessRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemberAccessRecordResponseBody
	GetCode() *string
	SetData(v *ListMemberAccessRecordResponseBodyData) *ListMemberAccessRecordResponseBody
	GetData() *ListMemberAccessRecordResponseBodyData
	SetHttpStatusCode(v int32) *ListMemberAccessRecordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMemberAccessRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMemberAccessRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemberAccessRecordResponseBody
	GetSuccess() *bool
}

type ListMemberAccessRecordResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListMemberAccessRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMemberAccessRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAccessRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemberAccessRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemberAccessRecordResponseBody) GetData() *ListMemberAccessRecordResponseBodyData {
	return s.Data
}

func (s *ListMemberAccessRecordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemberAccessRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemberAccessRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemberAccessRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemberAccessRecordResponseBody) SetCode(v string) *ListMemberAccessRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemberAccessRecordResponseBody) SetData(v *ListMemberAccessRecordResponseBodyData) *ListMemberAccessRecordResponseBody {
	s.Data = v
	return s
}

func (s *ListMemberAccessRecordResponseBody) SetHttpStatusCode(v int32) *ListMemberAccessRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemberAccessRecordResponseBody) SetMessage(v string) *ListMemberAccessRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemberAccessRecordResponseBody) SetRequestId(v string) *ListMemberAccessRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemberAccessRecordResponseBody) SetSuccess(v bool) *ListMemberAccessRecordResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemberAccessRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemberAccessRecordResponseBodyData struct {
	Num      *int32                                            `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListMemberAccessRecordResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                            `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMemberAccessRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAccessRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemberAccessRecordResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListMemberAccessRecordResponseBodyData) GetPageData() []*ListMemberAccessRecordResponseBodyDataPageData {
	return s.PageData
}

func (s *ListMemberAccessRecordResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListMemberAccessRecordResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMemberAccessRecordResponseBodyData) SetNum(v int32) *ListMemberAccessRecordResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyData) SetPageData(v []*ListMemberAccessRecordResponseBodyDataPageData) *ListMemberAccessRecordResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListMemberAccessRecordResponseBodyData) SetSize(v int32) *ListMemberAccessRecordResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyData) SetTotal(v int32) *ListMemberAccessRecordResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyData) Validate() error {
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

type ListMemberAccessRecordResponseBodyDataPageData struct {
	AccessDate            *int64  `json:"AccessDate,omitempty" xml:"AccessDate,omitempty"`
	AccessStatus          *string `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	AdminName             *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	AuthorizedCount       *int64  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	AuthorizedDeviceCount *int32  `json:"AuthorizedDeviceCount,omitempty" xml:"AuthorizedDeviceCount,omitempty"`
	Contactor             *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	MemberId              *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberResponseDate    *int64  `json:"MemberResponseDate,omitempty" xml:"MemberResponseDate,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Telephony             *string `json:"Telephony,omitempty" xml:"Telephony,omitempty"`
	Uid                   *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListMemberAccessRecordResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAccessRecordResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetAccessDate() *int64 {
	return s.AccessDate
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetAccessStatus() *string {
	return s.AccessStatus
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetAdminName() *string {
	return s.AdminName
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetAuthorizedDeviceCount() *int32 {
	return s.AuthorizedDeviceCount
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetContactor() *string {
	return s.Contactor
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetMemberId() *string {
	return s.MemberId
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetMemberResponseDate() *int64 {
	return s.MemberResponseDate
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetTelephony() *string {
	return s.Telephony
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) GetUid() *string {
	return s.Uid
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetAccessDate(v int64) *ListMemberAccessRecordResponseBodyDataPageData {
	s.AccessDate = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetAccessStatus(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.AccessStatus = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetAdminName(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.AdminName = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetAuthorizedCount(v int64) *ListMemberAccessRecordResponseBodyDataPageData {
	s.AuthorizedCount = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetAuthorizedDeviceCount(v int32) *ListMemberAccessRecordResponseBodyDataPageData {
	s.AuthorizedDeviceCount = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetContactor(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Contactor = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetMemberId(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.MemberId = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetMemberResponseDate(v int64) *ListMemberAccessRecordResponseBodyDataPageData {
	s.MemberResponseDate = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetName(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetRemark(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetStatus(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetTelephony(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Telephony = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) SetUid(v string) *ListMemberAccessRecordResponseBodyDataPageData {
	s.Uid = &v
	return s
}

func (s *ListMemberAccessRecordResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
