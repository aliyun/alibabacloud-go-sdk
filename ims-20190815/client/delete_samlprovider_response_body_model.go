// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSAMLProviderResponseBody
	GetRequestId() *string
}

type DeleteSAMLProviderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSAMLProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSAMLProviderResponseBody) SetRequestId(v string) *DeleteSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSAMLProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
