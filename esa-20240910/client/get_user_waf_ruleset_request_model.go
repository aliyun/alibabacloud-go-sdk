// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserWafRulesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetUserWafRulesetRequest
	GetId() *int64
	SetInstanceId(v string) *GetUserWafRulesetRequest
	GetInstanceId() *string
}

type GetUserWafRulesetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esa-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetUserWafRulesetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserWafRulesetRequest) GoString() string {
	return s.String()
}

func (s *GetUserWafRulesetRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUserWafRulesetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserWafRulesetRequest) SetId(v int64) *GetUserWafRulesetRequest {
	s.Id = &v
	return s
}

func (s *GetUserWafRulesetRequest) SetInstanceId(v string) *GetUserWafRulesetRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserWafRulesetRequest) Validate() error {
	return dara.Validate(s)
}
