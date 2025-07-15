// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntranetDomainPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIntranetDomainPolicyResponseBody
	GetRequestId() *string
}

type ModifyIntranetDomainPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIntranetDomainPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntranetDomainPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIntranetDomainPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIntranetDomainPolicyResponseBody) SetRequestId(v string) *ModifyIntranetDomainPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIntranetDomainPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
