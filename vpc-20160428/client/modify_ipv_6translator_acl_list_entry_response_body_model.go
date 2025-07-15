// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIPv6TranslatorAclListEntryResponseBody
	GetRequestId() *string
}

type ModifyIPv6TranslatorAclListEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIPv6TranslatorAclListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIPv6TranslatorAclListEntryResponseBody) SetRequestId(v string) *ModifyIPv6TranslatorAclListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
