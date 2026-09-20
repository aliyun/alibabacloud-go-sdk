// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillScopeResponseBody
	GetRequestId() *string
}

type UpdateSkillScopeResponseBody struct {
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// F4BFD370-7466-5F56-ACE5-A2D11A26C6BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSkillScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillScopeResponseBody) SetRequestId(v string) *UpdateSkillScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
