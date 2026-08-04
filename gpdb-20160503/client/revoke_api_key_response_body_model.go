// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeApiKeyResponseBody
	GetRequestId() *string
}

type RevokeApiKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeApiKeyResponseBody) SetRequestId(v string) *RevokeApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
