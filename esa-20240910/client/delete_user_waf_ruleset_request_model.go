// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteUserWafRulesetRequest
	GetId() *int64
	SetInstanceId(v string) *DeleteUserWafRulesetRequest
	GetInstanceId() *string
}

type DeleteUserWafRulesetRequest struct {
	// WAF ruleset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteUserWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteUserWafRulesetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserWafRulesetRequest) SetId(v int64) *DeleteUserWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *DeleteUserWafRulesetRequest) SetInstanceId(v string) *DeleteUserWafRulesetRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
