// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceLimitDetails interface {
	dara.Model
	String() string
	GoString() string
	SetGCLevel(v string) *ResourceLimitDetails
	GetGCLevel() *string
	SetResourceLimit(v map[string]interface{}) *ResourceLimitDetails
	GetResourceLimit() map[string]interface{}
	SetShouldIgnoreResourceCheck(v bool) *ResourceLimitDetails
	GetShouldIgnoreResourceCheck() *bool
}

type ResourceLimitDetails struct {
	GCLevel                   *string                `json:"GCLevel,omitempty" xml:"GCLevel,omitempty"`
	ResourceLimit             map[string]interface{} `json:"ResourceLimit,omitempty" xml:"ResourceLimit,omitempty"`
	ShouldIgnoreResourceCheck *bool                  `json:"ShouldIgnoreResourceCheck,omitempty" xml:"ShouldIgnoreResourceCheck,omitempty"`
}

func (s ResourceLimitDetails) String() string {
	return dara.Prettify(s)
}

func (s ResourceLimitDetails) GoString() string {
	return s.String()
}

func (s *ResourceLimitDetails) GetGCLevel() *string {
	return s.GCLevel
}

func (s *ResourceLimitDetails) GetResourceLimit() map[string]interface{} {
	return s.ResourceLimit
}

func (s *ResourceLimitDetails) GetShouldIgnoreResourceCheck() *bool {
	return s.ShouldIgnoreResourceCheck
}

func (s *ResourceLimitDetails) SetGCLevel(v string) *ResourceLimitDetails {
	s.GCLevel = &v
	return s
}

func (s *ResourceLimitDetails) SetResourceLimit(v map[string]interface{}) *ResourceLimitDetails {
	s.ResourceLimit = v
	return s
}

func (s *ResourceLimitDetails) SetShouldIgnoreResourceCheck(v bool) *ResourceLimitDetails {
	s.ShouldIgnoreResourceCheck = &v
	return s
}

func (s *ResourceLimitDetails) Validate() error {
	return dara.Validate(s)
}
