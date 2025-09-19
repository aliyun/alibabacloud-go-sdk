// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserSettingResponseBody
	GetRequestId() *string
}

type CreateUserSettingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1EE5AFD-0867-5F4F-9BE1-EBDD2C35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserSettingResponseBody) SetRequestId(v string) *CreateUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
