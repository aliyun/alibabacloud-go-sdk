// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTokenVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTokenVaultResponseBody
	GetRequestId() *string
}

type UpdateTokenVaultResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTokenVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTokenVaultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTokenVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTokenVaultResponseBody) SetRequestId(v string) *UpdateTokenVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTokenVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
