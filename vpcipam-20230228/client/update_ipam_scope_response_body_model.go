// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpamScopeResponseBody
	GetRequestId() *string
}

type UpdateIpamScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BA8054F5-852A-570A-ACFF-BCA63E0B02D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpamScopeResponseBody) SetRequestId(v string) *UpdateIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpamScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
