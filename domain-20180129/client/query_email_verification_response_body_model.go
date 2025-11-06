// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmailVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmIp(v string) *QueryEmailVerificationResponseBody
	GetConfirmIp() *string
	SetEmail(v string) *QueryEmailVerificationResponseBody
	GetEmail() *string
	SetEmailVerificationNo(v string) *QueryEmailVerificationResponseBody
	GetEmailVerificationNo() *string
	SetGmtCreate(v string) *QueryEmailVerificationResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *QueryEmailVerificationResponseBody
	GetGmtModified() *string
	SetRequestId(v string) *QueryEmailVerificationResponseBody
	GetRequestId() *string
	SetSendIp(v string) *QueryEmailVerificationResponseBody
	GetSendIp() *string
	SetTokenSendTime(v string) *QueryEmailVerificationResponseBody
	GetTokenSendTime() *string
	SetUserId(v string) *QueryEmailVerificationResponseBody
	GetUserId() *string
	SetVerificationStatus(v int32) *QueryEmailVerificationResponseBody
	GetVerificationStatus() *int32
	SetVerificationTime(v string) *QueryEmailVerificationResponseBody
	GetVerificationTime() *string
}

type QueryEmailVerificationResponseBody struct {
	// example:
	//
	// 42.*.*.31
	ConfirmIp *string `json:"ConfirmIp,omitempty" xml:"ConfirmIp,omitempty"`
	// example:
	//
	// abc@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 72b36ba0572e423bbb3f19640896****
	EmailVerificationNo *string `json:"EmailVerificationNo,omitempty" xml:"EmailVerificationNo,omitempty"`
	// example:
	//
	// 2019-02-19 16:38:07
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2019-02-19 16:40:38
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// FC4F7D02-8A83-4E37-B935-2D48A1B8423E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 42.*.*.115
	SendIp *string `json:"SendIp,omitempty" xml:"SendIp,omitempty"`
	// example:
	//
	// 2019-02-19 16:38:07
	TokenSendTime *string `json:"TokenSendTime,omitempty" xml:"TokenSendTime,omitempty"`
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	VerificationStatus *int32 `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	// example:
	//
	// 2019-02-19 16:40:38
	VerificationTime *string `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s QueryEmailVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationResponseBody) GetConfirmIp() *string {
	return s.ConfirmIp
}

func (s *QueryEmailVerificationResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryEmailVerificationResponseBody) GetEmailVerificationNo() *string {
	return s.EmailVerificationNo
}

func (s *QueryEmailVerificationResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryEmailVerificationResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryEmailVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEmailVerificationResponseBody) GetSendIp() *string {
	return s.SendIp
}

func (s *QueryEmailVerificationResponseBody) GetTokenSendTime() *string {
	return s.TokenSendTime
}

func (s *QueryEmailVerificationResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *QueryEmailVerificationResponseBody) GetVerificationStatus() *int32 {
	return s.VerificationStatus
}

func (s *QueryEmailVerificationResponseBody) GetVerificationTime() *string {
	return s.VerificationTime
}

func (s *QueryEmailVerificationResponseBody) SetConfirmIp(v string) *QueryEmailVerificationResponseBody {
	s.ConfirmIp = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetEmail(v string) *QueryEmailVerificationResponseBody {
	s.Email = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetEmailVerificationNo(v string) *QueryEmailVerificationResponseBody {
	s.EmailVerificationNo = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetGmtCreate(v string) *QueryEmailVerificationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetGmtModified(v string) *QueryEmailVerificationResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetRequestId(v string) *QueryEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetSendIp(v string) *QueryEmailVerificationResponseBody {
	s.SendIp = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetTokenSendTime(v string) *QueryEmailVerificationResponseBody {
	s.TokenSendTime = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetUserId(v string) *QueryEmailVerificationResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetVerificationStatus(v int32) *QueryEmailVerificationResponseBody {
	s.VerificationStatus = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetVerificationTime(v string) *QueryEmailVerificationResponseBody {
	s.VerificationTime = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
