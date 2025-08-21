// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindAliasResponseBody
	GetRequestId() *string
}

type BindAliasResponseBody struct {
	// example:
	//
	// 159E4422-6624-4750-8943-DFD98D34858C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAliasResponseBody) GoString() string {
	return s.String()
}

func (s *BindAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAliasResponseBody) SetRequestId(v string) *BindAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
