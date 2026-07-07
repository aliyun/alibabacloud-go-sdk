// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillSecurityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetIdentitySkillSecurityResponseBody
	GetRequestId() *string
}

type SetIdentitySkillSecurityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdentitySkillSecurityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillSecurityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetIdentitySkillSecurityResponseBody) SetRequestId(v string) *SetIdentitySkillSecurityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetIdentitySkillSecurityResponseBody) Validate() error {
	return dara.Validate(s)
}
