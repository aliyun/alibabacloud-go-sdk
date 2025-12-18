// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigRuleRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *GetAggregateConfigRuleRequest
	GetConfigRuleId() *string
	SetTag(v []*GetAggregateConfigRuleRequestTag) *GetAggregateConfigRuleRequest
	GetTag() []*GetAggregateConfigRuleRequestTag
}

type GetAggregateConfigRuleRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-7f00626622af0041****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the rule.
	//
	// You can call the [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html) operation to obtain the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	Tag []*GetAggregateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateConfigRuleRequest) GetTag() []*GetAggregateConfigRuleRequestTag {
	return s.Tag
}

func (s *GetAggregateConfigRuleRequest) SetAggregatorId(v string) *GetAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleRequest) SetConfigRuleId(v string) *GetAggregateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleRequest) SetTag(v []*GetAggregateConfigRuleRequestTag) *GetAggregateConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *GetAggregateConfigRuleRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateConfigRuleRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAggregateConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetAggregateConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetAggregateConfigRuleRequestTag) SetKey(v string) *GetAggregateConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *GetAggregateConfigRuleRequestTag) SetValue(v string) *GetAggregateConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *GetAggregateConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}
