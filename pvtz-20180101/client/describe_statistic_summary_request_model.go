// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeStatisticSummaryRequest
	GetLang() *string
	SetUserClientIp(v string) *DescribeStatisticSummaryRequest
	GetUserClientIp() *string
}

type DescribeStatisticSummaryRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeStatisticSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeStatisticSummaryRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeStatisticSummaryRequest) SetLang(v string) *DescribeStatisticSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStatisticSummaryRequest) SetUserClientIp(v string) *DescribeStatisticSummaryRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeStatisticSummaryRequest) Validate() error {
	return dara.Validate(s)
}
