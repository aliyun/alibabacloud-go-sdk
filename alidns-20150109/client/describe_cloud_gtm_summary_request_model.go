// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmSummaryRequest
	GetAcceptLanguage() *string
}

type DescribeCloudGtmSummaryRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeCloudGtmSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSummaryRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmSummaryRequest) SetAcceptLanguage(v string) *DescribeCloudGtmSummaryRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmSummaryRequest) Validate() error {
	return dara.Validate(s)
}
