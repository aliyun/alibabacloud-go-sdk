// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOriginPoolResponseBody
	GetRequestId() *string
}

type DeleteOriginPoolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOriginPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOriginPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOriginPoolResponseBody) SetRequestId(v string) *DeleteOriginPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOriginPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
