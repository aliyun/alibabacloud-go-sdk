// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDelegatedAdministratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegisterDelegatedAdministratorResponseBody
	GetRequestId() *string
}

type RegisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A45FC8F-54D2-4A65-8338-25E5DEBDA304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDelegatedAdministratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDelegatedAdministratorResponseBody) SetRequestId(v string) *RegisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDelegatedAdministratorResponseBody) Validate() error {
	return dara.Validate(s)
}
