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
	// The language of the content within the request. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
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
