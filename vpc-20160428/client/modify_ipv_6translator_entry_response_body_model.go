// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIPv6TranslatorEntryResponseBody
	GetRequestId() *string
}

type ModifyIPv6TranslatorEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIPv6TranslatorEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIPv6TranslatorEntryResponseBody) SetRequestId(v string) *ModifyIPv6TranslatorEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIPv6TranslatorEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
