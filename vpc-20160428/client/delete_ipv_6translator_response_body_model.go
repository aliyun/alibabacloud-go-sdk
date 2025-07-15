// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIPv6TranslatorResponseBody
	GetRequestId() *string
}

type DeleteIPv6TranslatorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIPv6TranslatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIPv6TranslatorResponseBody) SetRequestId(v string) *DeleteIPv6TranslatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIPv6TranslatorResponseBody) Validate() error {
	return dara.Validate(s)
}
