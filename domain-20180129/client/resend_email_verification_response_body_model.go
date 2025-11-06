// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailList(v []*ResendEmailVerificationResponseBodyFailList) *ResendEmailVerificationResponseBody
	GetFailList() []*ResendEmailVerificationResponseBodyFailList
	SetRequestId(v string) *ResendEmailVerificationResponseBody
	GetRequestId() *string
	SetSuccessList(v []*ResendEmailVerificationResponseBodySuccessList) *ResendEmailVerificationResponseBody
	GetSuccessList() []*ResendEmailVerificationResponseBodySuccessList
}

type ResendEmailVerificationResponseBody struct {
	FailList []*ResendEmailVerificationResponseBodyFailList `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	// example:
	//
	// 0EA54E99-DB48-4CE3-A099-6ED8E451B8AC
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*ResendEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s ResendEmailVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBody) GetFailList() []*ResendEmailVerificationResponseBodyFailList {
	return s.FailList
}

func (s *ResendEmailVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendEmailVerificationResponseBody) GetSuccessList() []*ResendEmailVerificationResponseBodySuccessList {
	return s.SuccessList
}

func (s *ResendEmailVerificationResponseBody) SetFailList(v []*ResendEmailVerificationResponseBodyFailList) *ResendEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *ResendEmailVerificationResponseBody) SetRequestId(v string) *ResendEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendEmailVerificationResponseBody) SetSuccessList(v []*ResendEmailVerificationResponseBodySuccessList) *ResendEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

func (s *ResendEmailVerificationResponseBody) Validate() error {
	if s.FailList != nil {
		for _, item := range s.FailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessList != nil {
		for _, item := range s.SuccessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResendEmailVerificationResponseBodyFailList struct {
	// example:
	//
	// SendTokenQuotaExceeded
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test1@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// The maximum number of attempts allowed to send the email verification link is exceeded.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ResendEmailVerificationResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBodyFailList) GetCode() *string {
	return s.Code
}

func (s *ResendEmailVerificationResponseBodyFailList) GetEmail() *string {
	return s.Email
}

func (s *ResendEmailVerificationResponseBodyFailList) GetMessage() *string {
	return s.Message
}

func (s *ResendEmailVerificationResponseBodyFailList) SetCode(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *ResendEmailVerificationResponseBodyFailList) SetEmail(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationResponseBodyFailList) SetMessage(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

func (s *ResendEmailVerificationResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type ResendEmailVerificationResponseBodySuccessList struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test2@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ResendEmailVerificationResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBodySuccessList) GetCode() *string {
	return s.Code
}

func (s *ResendEmailVerificationResponseBodySuccessList) GetEmail() *string {
	return s.Email
}

func (s *ResendEmailVerificationResponseBodySuccessList) GetMessage() *string {
	return s.Message
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetCode(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetEmail(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetMessage(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

func (s *ResendEmailVerificationResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
