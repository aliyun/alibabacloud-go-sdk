// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetGroupRequest
	GetGroupId() *string
}

type GetGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b1
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
