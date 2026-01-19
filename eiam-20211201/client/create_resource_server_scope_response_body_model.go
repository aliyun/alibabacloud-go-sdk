// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceServerScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceServerScopeResponseBody
	GetRequestId() *string
	SetResourceServerScopeId(v string) *CreateResourceServerScopeResponseBody
	GetResourceServerScopeId() *string
}

type CreateResourceServerScopeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ress_neg35flu6byysxwutaxu3dxxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
}

func (s CreateResourceServerScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceServerScopeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceServerScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceServerScopeResponseBody) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *CreateResourceServerScopeResponseBody) SetRequestId(v string) *CreateResourceServerScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceServerScopeResponseBody) SetResourceServerScopeId(v string) *CreateResourceServerScopeResponseBody {
	s.ResourceServerScopeId = &v
	return s
}

func (s *CreateResourceServerScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
