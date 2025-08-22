// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScenes(v string) *DescribeDcdnWafScenesRequest
	GetDefenseScenes() *string
}

type DescribeDcdnWafScenesRequest struct {
	// The types of the protection policies that you want to query. Separate multiple types with commas (,). Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: IP address whitelist
	//
	// 	- ip_blacklist: IP address blacklist
	//
	// 	- region_block: region blacklist
	//
	// 	- bot: bot management
	//
	// > If you do not set this parameter, all types of protection policies are queried.
	//
	// example:
	//
	// waf_group,custom_acl,whitelist
	DefenseScenes *string `json:"DefenseScenes,omitempty" xml:"DefenseScenes,omitempty"`
}

func (s DescribeDcdnWafScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafScenesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafScenesRequest) GetDefenseScenes() *string {
	return s.DefenseScenes
}

func (s *DescribeDcdnWafScenesRequest) SetDefenseScenes(v string) *DescribeDcdnWafScenesRequest {
	s.DefenseScenes = &v
	return s
}

func (s *DescribeDcdnWafScenesRequest) Validate() error {
	return dara.Validate(s)
}
