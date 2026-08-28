// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAndDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyAttachmentId(v string) *DetachAndDeletePolicyRequest
	GetPolicyAttachmentId() *string
}

type DetachAndDeletePolicyRequest struct {
	// The policy association ID.
	//
	// example:
	//
	// pr-cq7l5s5lhtgi6qasrdc0
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
}

func (s DetachAndDeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachAndDeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachAndDeletePolicyRequest) GetPolicyAttachmentId() *string {
	return s.PolicyAttachmentId
}

func (s *DetachAndDeletePolicyRequest) SetPolicyAttachmentId(v string) *DetachAndDeletePolicyRequest {
	s.PolicyAttachmentId = &v
	return s
}

func (s *DetachAndDeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
