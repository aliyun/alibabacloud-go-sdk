// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebFilterRuleIdsShrink(v string) *DeleteWebFilterRulesShrinkRequest
	GetWebFilterRuleIdsShrink() *string
}

type DeleteWebFilterRulesShrinkRequest struct {
	// The list of web filtering rule IDs to delete. You can specify up to 100 IDs in a single request. The IDs cannot be empty and must be unique.
	//
	// This parameter is required.
	WebFilterRuleIdsShrink *string `json:"WebFilterRuleIds,omitempty" xml:"WebFilterRuleIds,omitempty"`
}

func (s DeleteWebFilterRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterRulesShrinkRequest) GetWebFilterRuleIdsShrink() *string {
	return s.WebFilterRuleIdsShrink
}

func (s *DeleteWebFilterRulesShrinkRequest) SetWebFilterRuleIdsShrink(v string) *DeleteWebFilterRulesShrinkRequest {
	s.WebFilterRuleIdsShrink = &v
	return s
}

func (s *DeleteWebFilterRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
