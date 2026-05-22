// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteImageTransformResponseBody
	GetRequestId() *string
}

type DeleteImageTransformResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageTransformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageTransformResponseBody) SetRequestId(v string) *DeleteImageTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
