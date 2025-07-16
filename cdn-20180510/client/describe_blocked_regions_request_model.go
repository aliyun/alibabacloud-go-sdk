// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockedRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeBlockedRegionsRequest
	GetLanguage() *string
}

type DescribeBlockedRegionsRequest struct {
	// The language. Valid values:
	//
	// 	- **zh**: simplified Chinese
	//
	// 	- **en**: English
	//
	// 	- **jp**: Japanese
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeBlockedRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockedRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlockedRegionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeBlockedRegionsRequest) SetLanguage(v string) *DescribeBlockedRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeBlockedRegionsRequest) Validate() error {
	return dara.Validate(s)
}
