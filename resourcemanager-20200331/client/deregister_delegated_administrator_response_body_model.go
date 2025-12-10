// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterDelegatedAdministratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeregisterDelegatedAdministratorResponseBody
	GetRequestId() *string
}

type DeregisterDelegatedAdministratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DF5D5C52-7BD0-40E7-94C6-23A1505038A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeregisterDelegatedAdministratorResponseBody) SetRequestId(v string) *DeregisterDelegatedAdministratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterDelegatedAdministratorResponseBody) Validate() error {
	return dara.Validate(s)
}
