// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAssetStatisticRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAssetStatisticRequest
	GetSourceIp() *string
}

type DescribeAssetStatisticRequest struct {
	// Language for the request and response messages. Valid values:- **zh**: Chinese- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Source IP address of the requester.
	//
	// example:
	//
	// 112.239.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAssetStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAssetStatisticRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAssetStatisticRequest) SetLang(v string) *DescribeAssetStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetStatisticRequest) SetSourceIp(v string) *DescribeAssetStatisticRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetStatisticRequest) Validate() error {
	return dara.Validate(s)
}
