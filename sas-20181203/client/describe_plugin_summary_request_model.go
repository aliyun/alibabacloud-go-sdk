// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePluginSummaryRequest
	GetLang() *string
	SetPluginName(v string) *DescribePluginSummaryRequest
	GetPluginName() *string
}

type DescribePluginSummaryRequest struct {
	// The language of the content within the request and response.***	- Valid values:
	//
	// 	- **zh*	- (default)
	//
	// 	- **en**
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the plug-in. Valid values:
	//
	// 	- alinet: AliNet.
	//
	// 	- alisecguard: client protection.
	//
	// 	- alihips: AliHips.
	//
	// example:
	//
	// alihips
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
}

func (s DescribePluginSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePluginSummaryRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginSummaryRequest) SetLang(v string) *DescribePluginSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribePluginSummaryRequest) SetPluginName(v string) *DescribePluginSummaryRequest {
	s.PluginName = &v
	return s
}

func (s *DescribePluginSummaryRequest) Validate() error {
	return dara.Validate(s)
}
