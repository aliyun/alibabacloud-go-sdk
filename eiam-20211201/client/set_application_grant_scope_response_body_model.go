// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationGrantScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationGrantScopeResponseBody
	GetRequestId() *string
}

type SetApplicationGrantScopeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationGrantScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationGrantScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationGrantScopeResponseBody) SetRequestId(v string) *SetApplicationGrantScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationGrantScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
