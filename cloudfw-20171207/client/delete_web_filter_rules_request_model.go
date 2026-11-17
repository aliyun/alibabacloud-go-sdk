// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebFilterRuleIds(v []*string) *DeleteWebFilterRulesRequest
	GetWebFilterRuleIds() []*string
}

type DeleteWebFilterRulesRequest struct {
	// The list of web filtering rule IDs to delete. You can specify up to 100 IDs in a single request. The IDs cannot be empty and must be unique.
	//
	// This parameter is required.
	WebFilterRuleIds []*string `json:"WebFilterRuleIds,omitempty" xml:"WebFilterRuleIds,omitempty" type:"Repeated"`
}

func (s DeleteWebFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterRulesRequest) GetWebFilterRuleIds() []*string {
	return s.WebFilterRuleIds
}

func (s *DeleteWebFilterRulesRequest) SetWebFilterRuleIds(v []*string) *DeleteWebFilterRulesRequest {
	s.WebFilterRuleIds = v
	return s
}

func (s *DeleteWebFilterRulesRequest) Validate() error {
	return dara.Validate(s)
}
