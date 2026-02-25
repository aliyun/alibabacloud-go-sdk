// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParameterSetResponseBody
	GetRequestId() *string
}

type DeleteParameterSetResponseBody struct {
	// example:
	//
	// 708D670B-67C4-5653-9F88-8F7800429EE1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteParameterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterSetResponseBody) SetRequestId(v string) *DeleteParameterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterSetResponseBody) Validate() error {
	return dara.Validate(s)
}
