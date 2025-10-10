// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeZonesRequest
	GetAcceptLanguage() *string
}

type DescribeZonesRequest struct {
	// The supported language. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English
	//
	// 	- **ja**: Japanese
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
