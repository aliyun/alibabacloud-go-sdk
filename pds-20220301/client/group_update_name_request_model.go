// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUpdateNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GroupUpdateNameRequest
	GetGroupId() *string
	SetName(v string) *GroupUpdateNameRequest
	GetName() *string
}

type GroupUpdateNameRequest struct {
	// This parameter is required.
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GroupUpdateNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupUpdateNameRequest) GoString() string {
	return s.String()
}

func (s *GroupUpdateNameRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GroupUpdateNameRequest) GetName() *string {
	return s.Name
}

func (s *GroupUpdateNameRequest) SetGroupId(v string) *GroupUpdateNameRequest {
	s.GroupId = &v
	return s
}

func (s *GroupUpdateNameRequest) SetName(v string) *GroupUpdateNameRequest {
	s.Name = &v
	return s
}

func (s *GroupUpdateNameRequest) Validate() error {
	return dara.Validate(s)
}
