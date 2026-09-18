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
	// The dimension key name of the APM query filter condition. Specifies which dimension to filter by, such as hostname or service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// host.name
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The matching type of the APM query filter condition. Valid values:
	//
	// - ALL: Matches all values.
	//
	// - EQ: Exact match.
	//
	// - NE: Not equal to.
	//
	// - DISABLED: Disables the filter condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// EQ
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The filter value. Can be empty when type is set to ALL or DISABLED.
	//
	// example:
	//
	// cn-hangzhou
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
