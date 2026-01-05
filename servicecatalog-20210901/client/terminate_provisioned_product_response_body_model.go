// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateProvisionedProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateProvisionedProductResponseBody
	GetRequestId() *string
}

type TerminateProvisionedProductResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateProvisionedProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateProvisionedProductResponseBody) SetRequestId(v string) *TerminateProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateProvisionedProductResponseBody) Validate() error {
	return dara.Validate(s)
}
