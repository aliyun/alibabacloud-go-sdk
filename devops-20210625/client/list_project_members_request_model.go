// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetType(v string) *ListProjectMembersRequest
	GetTargetType() *string
}

type ListProjectMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Space
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListProjectMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListProjectMembersRequest) SetTargetType(v string) *ListProjectMembersRequest {
	s.TargetType = &v
	return s
}

func (s *ListProjectMembersRequest) Validate() error {
	return dara.Validate(s)
}
