// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *GetConfigRuleRequest
	GetConfigRuleId() *string
	SetTag(v []*GetConfigRuleRequestTag) *GetConfigRuleRequest
	GetTag() []*GetConfigRuleRequestTag
}

type GetConfigRuleRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
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
	Tag []*GetConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetConfigRuleRequest) GetTag() []*GetConfigRuleRequestTag {
	return s.Tag
}

func (s *GetConfigRuleRequest) SetConfigRuleId(v string) *GetConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleRequest) SetTag(v []*GetConfigRuleRequestTag) *GetConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *GetConfigRuleRequest) Validate() error {
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

type GetConfigRuleRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *GetConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetConfigRuleRequestTag) SetKey(v string) *GetConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *GetConfigRuleRequestTag) SetValue(v string) *GetConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *GetConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}
