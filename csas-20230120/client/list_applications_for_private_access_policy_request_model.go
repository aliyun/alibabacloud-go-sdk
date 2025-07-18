// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *ListApplicationsForPrivateAccessPolicyRequest
	GetPolicyIds() []*string
}

type ListApplicationsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApplicationsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListApplicationsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
