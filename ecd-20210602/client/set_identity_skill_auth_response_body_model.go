// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetIdentitySkillAuthResponseBody
	GetRequestId() *string
}

type SetIdentitySkillAuthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A87DBB05-653A-5E4B-B72B-5F4A1E07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdentitySkillAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillAuthResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetIdentitySkillAuthResponseBody) SetRequestId(v string) *SetIdentitySkillAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetIdentitySkillAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
