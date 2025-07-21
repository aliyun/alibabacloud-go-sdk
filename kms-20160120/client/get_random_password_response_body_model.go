// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRandomPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRandomPassword(v string) *GetRandomPasswordResponseBody
	GetRandomPassword() *string
	SetRequestId(v string) *GetRandomPasswordResponseBody
	GetRequestId() *string
}

type GetRandomPasswordResponseBody struct {
	// The generated random password.
	//
	// example:
	//
	// IxGn>NMmNB(y?iZ<Yc,_H/{2GC\\"U****
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6b0cbe25-5e33-467e-972e-7a83c6c97604
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRandomPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRandomPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordResponseBody) GetRandomPassword() *string {
	return s.RandomPassword
}

func (s *GetRandomPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRandomPasswordResponseBody) SetRandomPassword(v string) *GetRandomPasswordResponseBody {
	s.RandomPassword = &v
	return s
}

func (s *GetRandomPasswordResponseBody) SetRequestId(v string) *GetRandomPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRandomPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
