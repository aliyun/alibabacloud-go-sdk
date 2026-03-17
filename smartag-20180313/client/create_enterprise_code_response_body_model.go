// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEnterpriseCodeResponseBody
	GetRequestId() *string
}

type CreateEnterpriseCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnterpriseCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnterpriseCodeResponseBody) SetRequestId(v string) *CreateEnterpriseCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnterpriseCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
