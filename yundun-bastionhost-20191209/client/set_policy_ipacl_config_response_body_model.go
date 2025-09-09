// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyIPAclConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyIPAclConfigResponseBody
	GetRequestId() *string
}

type SetPolicyIPAclConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyIPAclConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyIPAclConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyIPAclConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyIPAclConfigResponseBody) SetRequestId(v string) *SetPolicyIPAclConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyIPAclConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
