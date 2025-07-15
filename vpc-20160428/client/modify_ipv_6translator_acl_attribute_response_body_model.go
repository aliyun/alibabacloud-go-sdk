// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIPv6TranslatorAclAttributeResponseBody
	GetRequestId() *string
}

type ModifyIPv6TranslatorAclAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIPv6TranslatorAclAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIPv6TranslatorAclAttributeResponseBody) SetRequestId(v string) *ModifyIPv6TranslatorAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
