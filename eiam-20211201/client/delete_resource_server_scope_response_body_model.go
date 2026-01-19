// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceServerScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceServerScopeResponseBody
	GetRequestId() *string
}

type DeleteResourceServerScopeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceServerScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceServerScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceServerScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceServerScopeResponseBody) SetRequestId(v string) *DeleteResourceServerScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceServerScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
