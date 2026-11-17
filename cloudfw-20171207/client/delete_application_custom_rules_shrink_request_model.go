// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationCustomRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCustomRuleIdsShrink(v string) *DeleteApplicationCustomRulesShrinkRequest
	GetApplicationCustomRuleIdsShrink() *string
}

type DeleteApplicationCustomRulesShrinkRequest struct {
	// The list of IDs of the custom application rules to delete. You can specify up to 100 IDs in a single request. The IDs cannot be empty and must be unique.
	//
	// This parameter is required.
	ApplicationCustomRuleIdsShrink *string `json:"ApplicationCustomRuleIds,omitempty" xml:"ApplicationCustomRuleIds,omitempty"`
}

func (s DeleteApplicationCustomRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationCustomRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationCustomRulesShrinkRequest) GetApplicationCustomRuleIdsShrink() *string {
	return s.ApplicationCustomRuleIdsShrink
}

func (s *DeleteApplicationCustomRulesShrinkRequest) SetApplicationCustomRuleIdsShrink(v string) *DeleteApplicationCustomRulesShrinkRequest {
	s.ApplicationCustomRuleIdsShrink = &v
	return s
}

func (s *DeleteApplicationCustomRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
