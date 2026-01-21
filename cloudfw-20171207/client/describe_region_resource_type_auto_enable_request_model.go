// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceTypeAutoEnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRegionResourceTypeAutoEnableRequest
	GetLang() *string
}

type DescribeRegionResourceTypeAutoEnableRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeRegionResourceTypeAutoEnableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceTypeAutoEnableRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceTypeAutoEnableRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRegionResourceTypeAutoEnableRequest) SetLang(v string) *DescribeRegionResourceTypeAutoEnableRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableRequest) Validate() error {
	return dara.Validate(s)
}
