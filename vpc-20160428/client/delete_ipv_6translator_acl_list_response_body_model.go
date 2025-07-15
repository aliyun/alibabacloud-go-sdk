// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorAclListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIPv6TranslatorAclListResponseBody
	GetRequestId() *string
}

type DeleteIPv6TranslatorAclListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIPv6TranslatorAclListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorAclListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorAclListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIPv6TranslatorAclListResponseBody) SetRequestId(v string) *DeleteIPv6TranslatorAclListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIPv6TranslatorAclListResponseBody) Validate() error {
	return dara.Validate(s)
}
