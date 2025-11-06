// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEmailVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListEmailVerificationResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListEmailVerificationResponseBodyData) *ListEmailVerificationResponseBody
	GetData() []*ListEmailVerificationResponseBodyData
	SetNextPage(v bool) *ListEmailVerificationResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *ListEmailVerificationResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *ListEmailVerificationResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *ListEmailVerificationResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListEmailVerificationResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListEmailVerificationResponseBody
	GetTotalPageNum() *int32
}

type ListEmailVerificationResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListEmailVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 78C60CC3-FF0A-44E2-989A-DDE0597791C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListEmailVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListEmailVerificationResponseBody) GetData() []*ListEmailVerificationResponseBodyData {
	return s.Data
}

func (s *ListEmailVerificationResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListEmailVerificationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEmailVerificationResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListEmailVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEmailVerificationResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListEmailVerificationResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListEmailVerificationResponseBody) SetCurrentPageNum(v int32) *ListEmailVerificationResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetData(v []*ListEmailVerificationResponseBodyData) *ListEmailVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ListEmailVerificationResponseBody) SetNextPage(v bool) *ListEmailVerificationResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetPageSize(v int32) *ListEmailVerificationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetPrePage(v bool) *ListEmailVerificationResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetRequestId(v string) *ListEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetTotalItemNum(v int32) *ListEmailVerificationResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetTotalPageNum(v int32) *ListEmailVerificationResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListEmailVerificationResponseBody) Validate() error {
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

type ListEmailVerificationResponseBodyData struct {
	// example:
	//
	// 127.0.0.1
	ConfirmIp *string `json:"ConfirmIp,omitempty" xml:"ConfirmIp,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 00000a21fd374da99d9c95b48500000
	EmailVerificationNo *string `json:"EmailVerificationNo,omitempty" xml:"EmailVerificationNo,omitempty"`
	// example:
	//
	// 2017-12-25 03:38:46
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2017-12-25 03:41:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 127.0.0.1
	SendIp *string `json:"SendIp,omitempty" xml:"SendIp,omitempty"`
	// example:
	//
	// 2017-12-25 03:38:46
	TokenSendTime *string `json:"TokenSendTime,omitempty" xml:"TokenSendTime,omitempty"`
	// example:
	//
	// 0000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	VerificationStatus *int32 `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	// example:
	//
	// 2017-12-25 03:41:11
	VerificationTime *string `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s ListEmailVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEmailVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponseBodyData) GetConfirmIp() *string {
	return s.ConfirmIp
}

func (s *ListEmailVerificationResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ListEmailVerificationResponseBodyData) GetEmailVerificationNo() *string {
	return s.EmailVerificationNo
}

func (s *ListEmailVerificationResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListEmailVerificationResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListEmailVerificationResponseBodyData) GetSendIp() *string {
	return s.SendIp
}

func (s *ListEmailVerificationResponseBodyData) GetTokenSendTime() *string {
	return s.TokenSendTime
}

func (s *ListEmailVerificationResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListEmailVerificationResponseBodyData) GetVerificationStatus() *int32 {
	return s.VerificationStatus
}

func (s *ListEmailVerificationResponseBodyData) GetVerificationTime() *string {
	return s.VerificationTime
}

func (s *ListEmailVerificationResponseBodyData) SetConfirmIp(v string) *ListEmailVerificationResponseBodyData {
	s.ConfirmIp = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetEmail(v string) *ListEmailVerificationResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetEmailVerificationNo(v string) *ListEmailVerificationResponseBodyData {
	s.EmailVerificationNo = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetGmtCreate(v string) *ListEmailVerificationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetGmtModified(v string) *ListEmailVerificationResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetSendIp(v string) *ListEmailVerificationResponseBodyData {
	s.SendIp = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetTokenSendTime(v string) *ListEmailVerificationResponseBodyData {
	s.TokenSendTime = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetUserId(v string) *ListEmailVerificationResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetVerificationStatus(v int32) *ListEmailVerificationResponseBodyData {
	s.VerificationStatus = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetVerificationTime(v string) *ListEmailVerificationResponseBodyData {
	s.VerificationTime = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
