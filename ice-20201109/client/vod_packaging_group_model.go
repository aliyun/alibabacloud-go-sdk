// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVodPackagingGroup interface {
	dara.Model
	String() string
	GoString() string
	SetApproximateAssetCount(v int64) *VodPackagingGroup
	GetApproximateAssetCount() *int64
	SetConfigurationCount(v int64) *VodPackagingGroup
	GetConfigurationCount() *int64
	SetCreateTime(v string) *VodPackagingGroup
	GetCreateTime() *string
	SetDescription(v string) *VodPackagingGroup
	GetDescription() *string
	SetDomainName(v string) *VodPackagingGroup
	GetDomainName() *string
	SetGroupName(v string) *VodPackagingGroup
	GetGroupName() *string
}

type VodPackagingGroup struct {
	ApproximateAssetCount *int64  `json:"ApproximateAssetCount,omitempty" xml:"ApproximateAssetCount,omitempty"`
	ConfigurationCount    *int64  `json:"ConfigurationCount,omitempty" xml:"ConfigurationCount,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupName             *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s VodPackagingGroup) String() string {
	return dara.Prettify(s)
}

func (s VodPackagingGroup) GoString() string {
	return s.String()
}

func (s *VodPackagingGroup) GetApproximateAssetCount() *int64 {
	return s.ApproximateAssetCount
}

func (s *VodPackagingGroup) GetConfigurationCount() *int64 {
	return s.ConfigurationCount
}

func (s *VodPackagingGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *VodPackagingGroup) GetDescription() *string {
	return s.Description
}

func (s *VodPackagingGroup) GetDomainName() *string {
	return s.DomainName
}

func (s *VodPackagingGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *VodPackagingGroup) SetApproximateAssetCount(v int64) *VodPackagingGroup {
	s.ApproximateAssetCount = &v
	return s
}

func (s *VodPackagingGroup) SetConfigurationCount(v int64) *VodPackagingGroup {
	s.ConfigurationCount = &v
	return s
}

func (s *VodPackagingGroup) SetCreateTime(v string) *VodPackagingGroup {
	s.CreateTime = &v
	return s
}

func (s *VodPackagingGroup) SetDescription(v string) *VodPackagingGroup {
	s.Description = &v
	return s
}

func (s *VodPackagingGroup) SetDomainName(v string) *VodPackagingGroup {
	s.DomainName = &v
	return s
}

func (s *VodPackagingGroup) SetGroupName(v string) *VodPackagingGroup {
	s.GroupName = &v
	return s
}

func (s *VodPackagingGroup) Validate() error {
	return dara.Validate(s)
}
