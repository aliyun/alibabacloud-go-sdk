// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeRegionsRequest
	GetLanguage() *string
}

type DescribeRegionsRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese. This is the default value.
	//
	// - en: English.
	//
	// - ja: Japanese.
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeRegionsRequest) SetLanguage(v string) *DescribeRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
