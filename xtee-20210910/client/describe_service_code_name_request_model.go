// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceCodeNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeServiceCodeNameRequest
	GetLang() *string
	SetTab(v string) *DescribeServiceCodeNameRequest
	GetTab() *string
}

type DescribeServiceCodeNameRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Scenario.
	//
	// example:
	//
	// FNAENIC
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s DescribeServiceCodeNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCodeNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceCodeNameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeServiceCodeNameRequest) GetTab() *string {
	return s.Tab
}

func (s *DescribeServiceCodeNameRequest) SetLang(v string) *DescribeServiceCodeNameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeServiceCodeNameRequest) SetTab(v string) *DescribeServiceCodeNameRequest {
	s.Tab = &v
	return s
}

func (s *DescribeServiceCodeNameRequest) Validate() error {
	return dara.Validate(s)
}
