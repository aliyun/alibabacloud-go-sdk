// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIdentityProviderResponseBody
	GetRequestId() *string
}

type DeleteIdentityProviderResponseBody struct {
	// example:
	//
	// EE854F60-E275-534A-B102-F75346B6DA38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdentityProviderResponseBody) SetRequestId(v string) *DeleteIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
