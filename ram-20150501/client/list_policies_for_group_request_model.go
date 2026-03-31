// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListPoliciesForGroupRequest
	GetGroupName() *string
}

type ListPoliciesForGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// dev
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListPoliciesForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListPoliciesForGroupRequest) SetGroupName(v string) *ListPoliciesForGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListPoliciesForGroupRequest) Validate() error {
	return dara.Validate(s)
}
