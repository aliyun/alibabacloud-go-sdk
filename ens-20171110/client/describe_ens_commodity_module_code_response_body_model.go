// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsCommodityModuleCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCodesInfo(v []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) *DescribeEnsCommodityModuleCodeResponseBody
	GetCommodityCodesInfo() []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo
	SetRequestId(v string) *DescribeEnsCommodityModuleCodeResponseBody
	GetRequestId() *string
}

type DescribeEnsCommodityModuleCodeResponseBody struct {
	CommodityCodesInfo []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo `json:"CommodityCodesInfo,omitempty" xml:"CommodityCodesInfo,omitempty" type:"Repeated"`
	RequestId          *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsCommodityModuleCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityModuleCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityModuleCodeResponseBody) GetCommodityCodesInfo() []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo {
	return s.CommodityCodesInfo
}

func (s *DescribeEnsCommodityModuleCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsCommodityModuleCodeResponseBody) SetCommodityCodesInfo(v []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) *DescribeEnsCommodityModuleCodeResponseBody {
	s.CommodityCodesInfo = v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBody) SetRequestId(v string) *DescribeEnsCommodityModuleCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo struct {
	CommodityCode   *string                                                                        `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ModuleCodesInfo []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo `json:"ModuleCodesInfo,omitempty" xml:"ModuleCodesInfo,omitempty" type:"Repeated"`
}

func (s DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) GetModuleCodesInfo() []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo {
	return s.ModuleCodesInfo
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) SetCommodityCode(v string) *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) SetModuleCodesInfo(v []*DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo {
	s.ModuleCodesInfo = v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) GoString() string {
	return s.String()
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) SetModuleCode(v string) *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) SetModuleName(v string) *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo {
	s.ModuleName = &v
	return s
}

func (s *DescribeEnsCommodityModuleCodeResponseBodyCommodityCodesInfoModuleCodesInfo) Validate() error {
	return dara.Validate(s)
}
