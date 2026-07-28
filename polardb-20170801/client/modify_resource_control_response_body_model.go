// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourceControlResponseBody
	GetRequestId() *string
}

type ModifyResourceControlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 47921222-0D37-4133-8C0D-017DC3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceControlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceControlResponseBody) SetRequestId(v string) *ModifyResourceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
