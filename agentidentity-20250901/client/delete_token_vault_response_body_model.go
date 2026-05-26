// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTokenVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTokenVaultResponseBody
	GetRequestId() *string
}

type DeleteTokenVaultResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTokenVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTokenVaultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTokenVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTokenVaultResponseBody) SetRequestId(v string) *DeleteTokenVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTokenVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
