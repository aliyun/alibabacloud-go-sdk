// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountAuthorityResponseBody
	GetRequestId() *string
}

type ModifyAccountAuthorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountAuthorityResponseBody) SetRequestId(v string) *ModifyAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
