// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyRuleStatusRequest
	GetId() *int64
	SetIds(v string) *ModifyRuleStatusRequest
	GetIds() *string
	SetLang(v string) *ModifyRuleStatusRequest
	GetLang() *string
	SetStatus(v int32) *ModifyRuleStatusRequest
	GetStatus() *int32
}

type ModifyRuleStatusRequest struct {
	// The ID of the sensitive data detection rule.
	//
	// > You can query the ID of the sensitive data detection rule by calling the **DescribeRules*	- operation.
	//
	// example:
	//
	// 12341
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the sensitive data detection rule. Separate multiple IDs with commas (,).
	//
	// > You can query the ID of the sensitive data detection rule by calling the **DescribeRules*	- operation.
	//
	// example:
	//
	// 1,2,3,4
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable or disable the sensitive data detection rule. Valid values:
	//
	// 	- **0**: disables the sensitive data detection rule.
	//
	// 	- **1**: enables the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyRuleStatusRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyRuleStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyRuleStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyRuleStatusRequest) SetId(v int64) *ModifyRuleStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetIds(v string) *ModifyRuleStatusRequest {
	s.Ids = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetLang(v string) *ModifyRuleStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetStatus(v int32) *ModifyRuleStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
