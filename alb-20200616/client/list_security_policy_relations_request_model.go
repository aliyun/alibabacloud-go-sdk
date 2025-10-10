// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPolicyRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRelationsRequest
	GetSecurityPolicyIds() []*string
}

type ListSecurityPolicyRelationsRequest struct {
	// The security policy IDs. You can specify up to five IDs.
	//
	// This parameter is required.
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPolicyRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsRequest) GetSecurityPolicyIds() []*string {
	return s.SecurityPolicyIds
}

func (s *ListSecurityPolicyRelationsRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRelationsRequest {
	s.SecurityPolicyIds = v
	return s
}

func (s *ListSecurityPolicyRelationsRequest) Validate() error {
	return dara.Validate(s)
}
