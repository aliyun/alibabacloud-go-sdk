// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGEnterpriseCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmartAGEnterpriseCodeResponseBody
	GetRequestId() *string
}

type UpdateSmartAGEnterpriseCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmartAGEnterpriseCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGEnterpriseCodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGEnterpriseCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAGEnterpriseCodeResponseBody) SetRequestId(v string) *UpdateSmartAGEnterpriseCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
