// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdminInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAdminInfoResponseBody
	GetCode() *string
	SetData(v *DescribeAdminInfoResponseBodyData) *DescribeAdminInfoResponseBody
	GetData() *DescribeAdminInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeAdminInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeAdminInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdminInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAdminInfoResponseBody
	GetSuccess() *bool
}

type DescribeAdminInfoResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeAdminInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAdminInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdminInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdminInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAdminInfoResponseBody) GetData() *DescribeAdminInfoResponseBodyData {
	return s.Data
}

func (s *DescribeAdminInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAdminInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdminInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdminInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdminInfoResponseBody) SetCode(v string) *DescribeAdminInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAdminInfoResponseBody) SetData(v *DescribeAdminInfoResponseBodyData) *DescribeAdminInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdminInfoResponseBody) SetHttpStatusCode(v int32) *DescribeAdminInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAdminInfoResponseBody) SetMessage(v string) *DescribeAdminInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdminInfoResponseBody) SetRequestId(v string) *DescribeAdminInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdminInfoResponseBody) SetSuccess(v bool) *DescribeAdminInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdminInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdminInfoResponseBodyData struct {
	AccessDate            *int64  `json:"AccessDate,omitempty" xml:"AccessDate,omitempty"`
	AccessStatus          *string `json:"AccessStatus,omitempty" xml:"AccessStatus,omitempty"`
	AuthorizedCount       *int64  `json:"AuthorizedCount,omitempty" xml:"AuthorizedCount,omitempty"`
	AuthorizedDeviceCount *int64  `json:"AuthorizedDeviceCount,omitempty" xml:"AuthorizedDeviceCount,omitempty"`
	Contactor             *string `json:"Contactor,omitempty" xml:"Contactor,omitempty"`
	MemberId              *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Telephony             *string `json:"Telephony,omitempty" xml:"Telephony,omitempty"`
	Uid                   *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeAdminInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdminInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdminInfoResponseBodyData) GetAccessDate() *int64 {
	return s.AccessDate
}

func (s *DescribeAdminInfoResponseBodyData) GetAccessStatus() *string {
	return s.AccessStatus
}

func (s *DescribeAdminInfoResponseBodyData) GetAuthorizedCount() *int64 {
	return s.AuthorizedCount
}

func (s *DescribeAdminInfoResponseBodyData) GetAuthorizedDeviceCount() *int64 {
	return s.AuthorizedDeviceCount
}

func (s *DescribeAdminInfoResponseBodyData) GetContactor() *string {
	return s.Contactor
}

func (s *DescribeAdminInfoResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *DescribeAdminInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeAdminInfoResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAdminInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdminInfoResponseBodyData) GetTelephony() *string {
	return s.Telephony
}

func (s *DescribeAdminInfoResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *DescribeAdminInfoResponseBodyData) SetAccessDate(v int64) *DescribeAdminInfoResponseBodyData {
	s.AccessDate = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetAccessStatus(v string) *DescribeAdminInfoResponseBodyData {
	s.AccessStatus = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetAuthorizedCount(v int64) *DescribeAdminInfoResponseBodyData {
	s.AuthorizedCount = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetAuthorizedDeviceCount(v int64) *DescribeAdminInfoResponseBodyData {
	s.AuthorizedDeviceCount = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetContactor(v string) *DescribeAdminInfoResponseBodyData {
	s.Contactor = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetMemberId(v string) *DescribeAdminInfoResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetName(v string) *DescribeAdminInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetRemark(v string) *DescribeAdminInfoResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetStatus(v string) *DescribeAdminInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetTelephony(v string) *DescribeAdminInfoResponseBodyData {
	s.Telephony = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) SetUid(v string) *DescribeAdminInfoResponseBodyData {
	s.Uid = &v
	return s
}

func (s *DescribeAdminInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
