// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMatchedMaliciousNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeMatchedMaliciousNamesRequest
	GetLang() *string
	SetLevels(v string) *DescribeMatchedMaliciousNamesRequest
	GetLevels() *string
}

type DescribeMatchedMaliciousNamesRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severity of the malicious image sample that you want to query. You can enter multiple severities. Separate the severities with commas (,). Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
}

func (s DescribeMatchedMaliciousNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMatchedMaliciousNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMatchedMaliciousNamesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeMatchedMaliciousNamesRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeMatchedMaliciousNamesRequest) SetLang(v string) *DescribeMatchedMaliciousNamesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesRequest) SetLevels(v string) *DescribeMatchedMaliciousNamesRequest {
	s.Levels = &v
	return s
}

func (s *DescribeMatchedMaliciousNamesRequest) Validate() error {
	return dara.Validate(s)
}
