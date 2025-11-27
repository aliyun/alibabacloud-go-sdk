// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRdsCustomInitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegisterUidSuccess(v bool) *RdsCustomInitResponseBody
	GetRegisterUidSuccess() *bool
	SetRequestId(v string) *RdsCustomInitResponseBody
	GetRequestId() *string
}

type RdsCustomInitResponseBody struct {
	// example:
	//
	// true
	RegisterUidSuccess *bool `json:"RegisterUidSuccess,omitempty" xml:"RegisterUidSuccess,omitempty"`
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RdsCustomInitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RdsCustomInitResponseBody) GoString() string {
	return s.String()
}

func (s *RdsCustomInitResponseBody) GetRegisterUidSuccess() *bool {
	return s.RegisterUidSuccess
}

func (s *RdsCustomInitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RdsCustomInitResponseBody) SetRegisterUidSuccess(v bool) *RdsCustomInitResponseBody {
	s.RegisterUidSuccess = &v
	return s
}

func (s *RdsCustomInitResponseBody) SetRequestId(v string) *RdsCustomInitResponseBody {
	s.RequestId = &v
	return s
}

func (s *RdsCustomInitResponseBody) Validate() error {
	return dara.Validate(s)
}
