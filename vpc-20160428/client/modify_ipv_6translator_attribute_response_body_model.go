// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIPv6TranslatorAttributeResponseBody
	GetRequestId() *string
}

type ModifyIPv6TranslatorAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B2F5262-6B57-43F2-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIPv6TranslatorAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIPv6TranslatorAttributeResponseBody) SetRequestId(v string) *ModifyIPv6TranslatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
