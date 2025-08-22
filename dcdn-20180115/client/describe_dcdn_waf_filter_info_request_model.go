// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafFilterInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScenes(v string) *DescribeDcdnWafFilterInfoRequest
	GetDefenseScenes() *string
	SetLanguage(v string) *DescribeDcdnWafFilterInfoRequest
	GetLanguage() *string
}

type DescribeDcdnWafFilterInfoRequest struct {
	// The type of the protection policy. Separate multiple types with commas (,). Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: IP address whitelist
	//
	// >If you do not specify this parameter, all types are returned.
	//
	// example:
	//
	// custom_acl
	DefenseScenes *string `json:"DefenseScenes,omitempty" xml:"DefenseScenes,omitempty"`
	// The language of the returned information. Valid values:
	//
	// 	- en: English
	//
	// 	- cn: Simplified Chinese
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeDcdnWafFilterInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoRequest) GetDefenseScenes() *string {
	return s.DefenseScenes
}

func (s *DescribeDcdnWafFilterInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDcdnWafFilterInfoRequest) SetDefenseScenes(v string) *DescribeDcdnWafFilterInfoRequest {
	s.DefenseScenes = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoRequest) SetLanguage(v string) *DescribeDcdnWafFilterInfoRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoRequest) Validate() error {
	return dara.Validate(s)
}
