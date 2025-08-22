// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGeoInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeDcdnWafGeoInfoRequest
	GetLanguage() *string
}

type DescribeDcdnWafGeoInfoRequest struct {
	// The language of the information to return. Valid values:
	//
	// 	- cn: Chinese
	//
	// 	- en: English
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeDcdnWafGeoInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDcdnWafGeoInfoRequest) SetLanguage(v string) *DescribeDcdnWafGeoInfoRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoRequest) Validate() error {
	return dara.Validate(s)
}
