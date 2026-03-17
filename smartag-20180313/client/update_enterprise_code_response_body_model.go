// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEnterpriseCodeResponseBody
	GetRequestId() *string
}

type UpdateEnterpriseCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnterpriseCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnterpriseCodeResponseBody) SetRequestId(v string) *UpdateEnterpriseCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnterpriseCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
