// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *ListTagsForPrivateAccessPolicyRequest
	GetPolicyIds() []*string
}

type ListTagsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListTagsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListTagsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
