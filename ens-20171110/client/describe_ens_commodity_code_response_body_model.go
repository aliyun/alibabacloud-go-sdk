// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCodeInfo(v []*DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) *DescribeEnsCommodityCodeResponseBody
	GetCommodityCodeInfo() []*DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo
	SetRequestId(v string) *DescribeEnsCommodityCodeResponseBody
	GetRequestId() *string
}

type DescribeEnsCommodityCodeResponseBody struct {
	CommodityCodeInfo []*DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo `json:"CommodityCodeInfo,omitempty" xml:"CommodityCodeInfo,omitempty" type:"Repeated"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsCommodityCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityCodeResponseBody) GetCommodityCodeInfo() []*DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo {
	return s.CommodityCodeInfo
}

func (s *DescribeEnsCommodityCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsCommodityCodeResponseBody) SetCommodityCodeInfo(v []*DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) *DescribeEnsCommodityCodeResponseBody {
	s.CommodityCodeInfo = v
	return s
}

func (s *DescribeEnsCommodityCodeResponseBody) SetRequestId(v string) *DescribeEnsCommodityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsCommodityCodeResponseBody) Validate() error {
	if s.CommodityCodeInfo != nil {
		for _, item := range s.CommodityCodeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
}

func (s DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) SetCommodityCode(v string) *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) SetCommodityName(v string) *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo {
	s.CommodityName = &v
	return s
}

func (s *DescribeEnsCommodityCodeResponseBodyCommodityCodeInfo) Validate() error {
	return dara.Validate(s)
}
