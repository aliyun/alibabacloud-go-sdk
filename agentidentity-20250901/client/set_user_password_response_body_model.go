// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRandomPassword(v string) *SetUserPasswordResponseBody
	GetRandomPassword() *string
	SetRequestId(v string) *SetUserPasswordResponseBody
	GetRequestId() *string
}

type SetUserPasswordResponseBody struct {
	// The system-generated random password returned when GenerateRandomPassword is set to true.
	//
	// example:
	//
	// Xk9mP2vL5nQ8
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPasswordResponseBody) GetRandomPassword() *string {
	return s.RandomPassword
}

func (s *SetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserPasswordResponseBody) SetRandomPassword(v string) *SetUserPasswordResponseBody {
	s.RandomPassword = &v
	return s
}

func (s *SetUserPasswordResponseBody) SetRequestId(v string) *SetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
