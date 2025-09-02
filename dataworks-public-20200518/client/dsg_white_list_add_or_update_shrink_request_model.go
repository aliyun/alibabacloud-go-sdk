// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListAddOrUpdateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWhiteListsShrink(v string) *DsgWhiteListAddOrUpdateShrinkRequest
	GetWhiteListsShrink() *string
}

type DsgWhiteListAddOrUpdateShrinkRequest struct {
	// A collection of whitelists.
	//
	// This parameter is required.
	WhiteListsShrink *string `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty"`
}

func (s DsgWhiteListAddOrUpdateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListAddOrUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgWhiteListAddOrUpdateShrinkRequest) GetWhiteListsShrink() *string {
	return s.WhiteListsShrink
}

func (s *DsgWhiteListAddOrUpdateShrinkRequest) SetWhiteListsShrink(v string) *DsgWhiteListAddOrUpdateShrinkRequest {
	s.WhiteListsShrink = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
