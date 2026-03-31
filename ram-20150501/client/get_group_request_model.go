// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetGroupRequest
	GetGroupName() *string
}

type GetGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupRequest) SetGroupName(v string) *GetGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
