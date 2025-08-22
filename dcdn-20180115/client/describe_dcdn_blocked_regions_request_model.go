// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBlockedRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeDcdnBlockedRegionsRequest
	GetLanguage() *string
}

type DescribeDcdnBlockedRegionsRequest struct {
	// The language. Valid values: zh, en, and jp.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeDcdnBlockedRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDcdnBlockedRegionsRequest) SetLanguage(v string) *DescribeDcdnBlockedRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsRequest) Validate() error {
	return dara.Validate(s)
}
