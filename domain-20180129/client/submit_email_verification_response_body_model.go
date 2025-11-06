// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEmailVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExistList(v []*SubmitEmailVerificationResponseBodyExistList) *SubmitEmailVerificationResponseBody
	GetExistList() []*SubmitEmailVerificationResponseBodyExistList
	SetFailList(v []*SubmitEmailVerificationResponseBodyFailList) *SubmitEmailVerificationResponseBody
	GetFailList() []*SubmitEmailVerificationResponseBodyFailList
	SetRequestId(v string) *SubmitEmailVerificationResponseBody
	GetRequestId() *string
	SetSuccessList(v []*SubmitEmailVerificationResponseBodySuccessList) *SubmitEmailVerificationResponseBody
	GetSuccessList() []*SubmitEmailVerificationResponseBodySuccessList
}

type SubmitEmailVerificationResponseBody struct {
	ExistList []*SubmitEmailVerificationResponseBodyExistList `json:"ExistList,omitempty" xml:"ExistList,omitempty" type:"Repeated"`
	FailList  []*SubmitEmailVerificationResponseBodyFailList  `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	// example:
	//
	// E2A8A5EF-DF8A-4C48-8FD4-9F6BD71AB26D
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*SubmitEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s SubmitEmailVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBody) GetExistList() []*SubmitEmailVerificationResponseBodyExistList {
	return s.ExistList
}

func (s *SubmitEmailVerificationResponseBody) GetFailList() []*SubmitEmailVerificationResponseBodyFailList {
	return s.FailList
}

func (s *SubmitEmailVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitEmailVerificationResponseBody) GetSuccessList() []*SubmitEmailVerificationResponseBodySuccessList {
	return s.SuccessList
}

func (s *SubmitEmailVerificationResponseBody) SetExistList(v []*SubmitEmailVerificationResponseBodyExistList) *SubmitEmailVerificationResponseBody {
	s.ExistList = v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetFailList(v []*SubmitEmailVerificationResponseBodyFailList) *SubmitEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetRequestId(v string) *SubmitEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetSuccessList(v []*SubmitEmailVerificationResponseBodySuccessList) *SubmitEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

func (s *SubmitEmailVerificationResponseBody) Validate() error {
	if s.ExistList != nil {
		for _, item := range s.ExistList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type SubmitEmailVerificationResponseBodyExistList struct {
	// example:
	//
	// SendTokenQuotaExceeded
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// The maximum number of attempts allowed to send the email verification link is exceeded.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodyExistList) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodyExistList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodyExistList) GetCode() *string {
	return s.Code
}

func (s *SubmitEmailVerificationResponseBodyExistList) GetEmail() *string {
	return s.Email
}

func (s *SubmitEmailVerificationResponseBodyExistList) GetMessage() *string {
	return s.Message
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetCode(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetEmail(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetMessage(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Message = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyExistList) Validate() error {
	return dara.Validate(s)
}

type SubmitEmailVerificationResponseBodyFailList struct {
	// example:
	//
	// SendTokenQuotaExceeded
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// The maximum number of attempts allowed to send the email verification link is exceeded
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodyFailList) GetCode() *string {
	return s.Code
}

func (s *SubmitEmailVerificationResponseBodyFailList) GetEmail() *string {
	return s.Email
}

func (s *SubmitEmailVerificationResponseBodyFailList) GetMessage() *string {
	return s.Message
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetCode(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetEmail(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetMessage(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type SubmitEmailVerificationResponseBodySuccessList struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodySuccessList) GetCode() *string {
	return s.Code
}

func (s *SubmitEmailVerificationResponseBodySuccessList) GetEmail() *string {
	return s.Email
}

func (s *SubmitEmailVerificationResponseBodySuccessList) GetMessage() *string {
	return s.Message
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetCode(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetEmail(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetMessage(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
