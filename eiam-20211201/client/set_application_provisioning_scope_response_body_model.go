// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationProvisioningScopeResponseBody
	GetRequestId() *string
}

type SetApplicationProvisioningScopeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationProvisioningScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationProvisioningScopeResponseBody) SetRequestId(v string) *SetApplicationProvisioningScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationProvisioningScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
