// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceServerScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceServerScopeResponseBody
	GetRequestId() *string
}

type UpdateResourceServerScopeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceServerScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceServerScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceServerScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceServerScopeResponseBody) SetRequestId(v string) *UpdateResourceServerScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceServerScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
