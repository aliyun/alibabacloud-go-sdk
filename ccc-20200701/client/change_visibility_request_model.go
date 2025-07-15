// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ChangeVisibilityRequest
	GetInstanceId() *string
	SetInvisible(v bool) *ChangeVisibilityRequest
	GetInvisible() *bool
	SetUserId(v string) *ChangeVisibilityRequest
	GetUserId() *string
}

type ChangeVisibilityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// szpczf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Invisible *bool `json:"Invisible,omitempty" xml:"Invisible,omitempty"`
	// example:
	//
	// sam@szpczf
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ChangeVisibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityRequest) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeVisibilityRequest) GetInvisible() *bool {
	return s.Invisible
}

func (s *ChangeVisibilityRequest) GetUserId() *string {
	return s.UserId
}

func (s *ChangeVisibilityRequest) SetInstanceId(v string) *ChangeVisibilityRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeVisibilityRequest) SetInvisible(v bool) *ChangeVisibilityRequest {
	s.Invisible = &v
	return s
}

func (s *ChangeVisibilityRequest) SetUserId(v string) *ChangeVisibilityRequest {
	s.UserId = &v
	return s
}

func (s *ChangeVisibilityRequest) Validate() error {
	return dara.Validate(s)
}
