// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIPv6TranslatorEntryResponseBody
	GetRequestId() *string
}

type DeleteIPv6TranslatorEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B2F5262-6B57-43F2-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIPv6TranslatorEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIPv6TranslatorEntryResponseBody) SetRequestId(v string) *DeleteIPv6TranslatorEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
