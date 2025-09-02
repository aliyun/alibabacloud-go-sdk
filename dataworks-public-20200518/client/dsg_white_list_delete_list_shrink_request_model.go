// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListDeleteListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DsgWhiteListDeleteListShrinkRequest
	GetIdsShrink() *string
}

type DsgWhiteListDeleteListShrinkRequest struct {
	// The IDs of the whitelists.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DsgWhiteListDeleteListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListDeleteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgWhiteListDeleteListShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgWhiteListDeleteListShrinkRequest) SetIdsShrink(v string) *DsgWhiteListDeleteListShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgWhiteListDeleteListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
