// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUpdateDesensStatusListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesensStatus(v int32) *DsgUpdateDesensStatusListShrinkRequest
	GetDesensStatus() *int32
	SetIdsShrink(v string) *DsgUpdateDesensStatusListShrinkRequest
	GetIdsShrink() *string
}

type DsgUpdateDesensStatusListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesensStatus *int32 `json:"DesensStatus,omitempty" xml:"DesensStatus,omitempty"`
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DsgUpdateDesensStatusListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUpdateDesensStatusListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgUpdateDesensStatusListShrinkRequest) GetDesensStatus() *int32 {
	return s.DesensStatus
}

func (s *DsgUpdateDesensStatusListShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgUpdateDesensStatusListShrinkRequest) SetDesensStatus(v int32) *DsgUpdateDesensStatusListShrinkRequest {
	s.DesensStatus = &v
	return s
}

func (s *DsgUpdateDesensStatusListShrinkRequest) SetIdsShrink(v string) *DsgUpdateDesensStatusListShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgUpdateDesensStatusListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
