// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetEngine(v string) *DescribeRegionsRequest
	GetEngine() *string
}

type DescribeRegionsRequest struct {
	// The supported language. Valid values:
	//
	// - **zh-CN**: Chinese (default)
	//
	// - **en-US**: English
	//
	// - **ja**: Japanese.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The data engine type. Valid values:
	//
	// - **hbase**: ApsaraDB for HBase Standard Edition or ApsaraDB for HBase single-node edition.
	//
	// - **hbaseue**: ApsaraDB for HBase Performance-enhanced Edition.
	//
	// - **serverlesshbase**: ApsaraDB for HBase Serverless edition.
	//
	// - **bds**: BDS instance.
	//
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetEngine(v string) *DescribeRegionsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
