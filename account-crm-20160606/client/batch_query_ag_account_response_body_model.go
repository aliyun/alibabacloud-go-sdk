// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccounts(v []*BatchQueryAgAccountResponseBodyAgAccounts) *BatchQueryAgAccountResponseBody
	GetAgAccounts() []*BatchQueryAgAccountResponseBodyAgAccounts
	SetCode(v string) *BatchQueryAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *BatchQueryAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchQueryAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchQueryAgAccountResponseBody
	GetSuccess() *bool
}

type BatchQueryAgAccountResponseBody struct {
	AgAccounts []*BatchQueryAgAccountResponseBodyAgAccounts `json:"AgAccounts,omitempty" xml:"AgAccounts,omitempty" type:"Repeated"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchQueryAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryAgAccountResponseBody) GetAgAccounts() []*BatchQueryAgAccountResponseBodyAgAccounts {
	return s.AgAccounts
}

func (s *BatchQueryAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchQueryAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryAgAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchQueryAgAccountResponseBody) SetAgAccounts(v []*BatchQueryAgAccountResponseBodyAgAccounts) *BatchQueryAgAccountResponseBody {
	s.AgAccounts = v
	return s
}

func (s *BatchQueryAgAccountResponseBody) SetCode(v string) *BatchQueryAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryAgAccountResponseBody) SetMessage(v string) *BatchQueryAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryAgAccountResponseBody) SetRequestId(v string) *BatchQueryAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryAgAccountResponseBody) SetSuccess(v bool) *BatchQueryAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryAgAccountResponseBody) Validate() error {
	if s.AgAccounts != nil {
		for _, item := range s.AgAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryAgAccountResponseBodyAgAccounts struct {
	LoginEmail   *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	Pk           *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ShowNickName *string `json:"ShowNickName,omitempty" xml:"ShowNickName,omitempty"`
}

func (s BatchQueryAgAccountResponseBodyAgAccounts) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryAgAccountResponseBodyAgAccounts) GoString() string {
	return s.String()
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) GetLoginEmail() *string {
	return s.LoginEmail
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) GetPk() *string {
	return s.Pk
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) GetShowNickName() *string {
	return s.ShowNickName
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) SetLoginEmail(v string) *BatchQueryAgAccountResponseBodyAgAccounts {
	s.LoginEmail = &v
	return s
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) SetPk(v string) *BatchQueryAgAccountResponseBodyAgAccounts {
	s.Pk = &v
	return s
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) SetShowNickName(v string) *BatchQueryAgAccountResponseBodyAgAccounts {
	s.ShowNickName = &v
	return s
}

func (s *BatchQueryAgAccountResponseBodyAgAccounts) Validate() error {
	return dara.Validate(s)
}
