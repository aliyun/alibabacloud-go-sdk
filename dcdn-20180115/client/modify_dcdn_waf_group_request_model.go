// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyDcdnWafGroupRequest
	GetId() *int64
	SetName(v string) *ModifyDcdnWafGroupRequest
	GetName() *string
	SetRules(v string) *ModifyDcdnWafGroupRequest
	GetRules() *string
}

type ModifyDcdnWafGroupRequest struct {
	// The ID of the custom WAF rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000110
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom WAF rule group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The incremental modifications of the rules in the custom WAF rule group. The value is a JSON string.
	//
	// example:
	//
	// {\\"All\\":false,\\"Op\\":\\"del\\",\\"List\\":\\"900109\\"}
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyDcdnWafGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyDcdnWafGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDcdnWafGroupRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyDcdnWafGroupRequest) SetId(v int64) *ModifyDcdnWafGroupRequest {
	s.Id = &v
	return s
}

func (s *ModifyDcdnWafGroupRequest) SetName(v string) *ModifyDcdnWafGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyDcdnWafGroupRequest) SetRules(v string) *ModifyDcdnWafGroupRequest {
	s.Rules = &v
	return s
}

func (s *ModifyDcdnWafGroupRequest) Validate() error {
	return dara.Validate(s)
}
