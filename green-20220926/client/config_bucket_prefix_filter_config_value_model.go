// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBucketPrefixFilterConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixFilterType(v string) *ConfigBucketPrefixFilterConfigValue
	GetPrefixFilterType() *string
	SetPrefixFilters(v []*string) *ConfigBucketPrefixFilterConfigValue
	GetPrefixFilters() []*string
}

type ConfigBucketPrefixFilterConfigValue struct {
	PrefixFilterType *string   `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	PrefixFilters    []*string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty" type:"Repeated"`
}

func (s ConfigBucketPrefixFilterConfigValue) String() string {
	return dara.Prettify(s)
}

func (s ConfigBucketPrefixFilterConfigValue) GoString() string {
	return s.String()
}

func (s *ConfigBucketPrefixFilterConfigValue) GetPrefixFilterType() *string {
	return s.PrefixFilterType
}

func (s *ConfigBucketPrefixFilterConfigValue) GetPrefixFilters() []*string {
	return s.PrefixFilters
}

func (s *ConfigBucketPrefixFilterConfigValue) SetPrefixFilterType(v string) *ConfigBucketPrefixFilterConfigValue {
	s.PrefixFilterType = &v
	return s
}

func (s *ConfigBucketPrefixFilterConfigValue) SetPrefixFilters(v []*string) *ConfigBucketPrefixFilterConfigValue {
	s.PrefixFilters = v
	return s
}

func (s *ConfigBucketPrefixFilterConfigValue) Validate() error {
	return dara.Validate(s)
}
