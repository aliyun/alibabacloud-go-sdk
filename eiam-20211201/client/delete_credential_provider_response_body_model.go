// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCredentialProviderResponseBody
	GetRequestId() *string
}

type DeleteCredentialProviderResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCredentialProviderResponseBody) SetRequestId(v string) *DeleteCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
