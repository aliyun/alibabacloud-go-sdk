// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupDeleteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DsgUserGroupDeleteShrinkRequest
	GetIdsShrink() *string
}

type DsgUserGroupDeleteShrinkRequest struct {
	// The information about the user group.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DsgUserGroupDeleteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupDeleteShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgUserGroupDeleteShrinkRequest) SetIdsShrink(v string) *DsgUserGroupDeleteShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgUserGroupDeleteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
