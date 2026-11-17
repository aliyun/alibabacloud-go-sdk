// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationCustomRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCustomRuleIds(v []*string) *DeleteApplicationCustomRulesRequest
	GetApplicationCustomRuleIds() []*string
}

type DeleteApplicationCustomRulesRequest struct {
	// The list of IDs of the custom application rules to delete. You can specify up to 100 IDs in a single request. The IDs cannot be empty and must be unique.
	//
	// This parameter is required.
	ApplicationCustomRuleIds []*string `json:"ApplicationCustomRuleIds,omitempty" xml:"ApplicationCustomRuleIds,omitempty" type:"Repeated"`
}

func (s DeleteApplicationCustomRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationCustomRulesRequest) GetApplicationCustomRuleIds() []*string {
	return s.ApplicationCustomRuleIds
}

func (s *DeleteApplicationCustomRulesRequest) SetApplicationCustomRuleIds(v []*string) *DeleteApplicationCustomRulesRequest {
	s.ApplicationCustomRuleIds = v
	return s
}

func (s *DeleteApplicationCustomRulesRequest) Validate() error {
	return dara.Validate(s)
}
