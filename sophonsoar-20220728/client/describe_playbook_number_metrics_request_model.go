// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNumberMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePlaybookNumberMetricsRequest
	GetLang() *string
}

type DescribePlaybookNumberMetricsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePlaybookNumberMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNumberMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookNumberMetricsRequest) SetLang(v string) *DescribePlaybookNumberMetricsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookNumberMetricsRequest) Validate() error {
	return dara.Validate(s)
}
