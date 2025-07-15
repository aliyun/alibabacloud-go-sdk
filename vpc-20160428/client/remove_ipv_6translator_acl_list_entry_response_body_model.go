// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIPv6TranslatorAclListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveIPv6TranslatorAclListEntryResponseBody
	GetRequestId() *string
}

type RemoveIPv6TranslatorAclListEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIPv6TranslatorAclListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveIPv6TranslatorAclListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIPv6TranslatorAclListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveIPv6TranslatorAclListEntryResponseBody) SetRequestId(v string) *RemoveIPv6TranslatorAclListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
