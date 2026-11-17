// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApplicationStatisticsRequest
	GetLang() *string
}

type DescribeApplicationStatisticsRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh**: Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// This parameter determines whether the category names and technology type names are returned in Chinese or English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeApplicationStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApplicationStatisticsRequest) SetLang(v string) *DescribeApplicationStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApplicationStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
