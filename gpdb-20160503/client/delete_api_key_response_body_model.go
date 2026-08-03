// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DeleteApiKeyResponseBody
	GetKeyId() *string
	SetRequestId(v string) *DeleteApiKeyResponseBody
	GetRequestId() *string
}

type DeleteApiKeyResponseBody struct {
	// API KEY ID。
	//
	// example:
	//
	// api-xxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *DeleteApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiKeyResponseBody) SetKeyId(v string) *DeleteApiKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetRequestId(v string) *DeleteApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
