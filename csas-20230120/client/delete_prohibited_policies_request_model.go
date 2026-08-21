// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *DeleteProhibitedPoliciesRequest
	GetPolicyIds() []*string
}

type DeleteProhibitedPoliciesRequest struct {
	// The IDs of the software prohibition policies to delete. Duplicate IDs are not allowed. You can specify up to 100 IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s DeleteProhibitedPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DeleteProhibitedPoliciesRequest) SetPolicyIds(v []*string) *DeleteProhibitedPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteProhibitedPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
