// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOIDCProviderResponseBody
	GetRequestId() *string
}

type DeleteOIDCProviderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOIDCProviderResponseBody) SetRequestId(v string) *DeleteOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOIDCProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
