// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApmFilterConfig interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ApmFilterConfig
	GetKey() *string
	SetType(v string) *ApmFilterConfig
	GetType() *string
	SetValue(v string) *ApmFilterConfig
	GetValue() *string
}

type ApmFilterConfig struct {
	// The tag key to filter by. For example, to filter traces by region, set this parameter to `RegionId`.
	//
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The comparison operator used to match the tag\\"s value. Valid values: `EQUAL` and `NOT_EQUAL`.
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value to compare against the tag\\"s value. Used with the `key` and `type` parameters to form a complete filter condition.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ApmFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s ApmFilterConfig) GoString() string {
	return s.String()
}

func (s *ApmFilterConfig) GetKey() *string {
	return s.Key
}

func (s *ApmFilterConfig) GetType() *string {
	return s.Type
}

func (s *ApmFilterConfig) GetValue() *string {
	return s.Value
}

func (s *ApmFilterConfig) SetKey(v string) *ApmFilterConfig {
	s.Key = &v
	return s
}

func (s *ApmFilterConfig) SetType(v string) *ApmFilterConfig {
	s.Type = &v
	return s
}

func (s *ApmFilterConfig) SetValue(v string) *ApmFilterConfig {
	s.Value = &v
	return s
}

func (s *ApmFilterConfig) Validate() error {
	return dara.Validate(s)
}
