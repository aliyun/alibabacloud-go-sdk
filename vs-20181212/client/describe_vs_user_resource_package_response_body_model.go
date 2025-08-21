// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUserResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVsUserResourcePackageResponseBody
	GetRequestId() *string
	SetResourcePackageInfos(v *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) *DescribeVsUserResourcePackageResponseBody
	GetResourcePackageInfos() *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos
}

type DescribeVsUserResourcePackageResponseBody struct {
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePackageInfos *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeVsUserResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsUserResourcePackageResponseBody) GetResourcePackageInfos() *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos {
	return s.ResourcePackageInfos
}

func (s *DescribeVsUserResourcePackageResponseBody) SetRequestId(v string) *DescribeVsUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) *DescribeVsUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) GetResourcePackageInfo() []*DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	return s.ResourcePackageInfo
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrCapacity  *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InitCapacity  *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacity() *string {
	return s.CurrCapacity
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacity() *string {
	return s.InitCapacity
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) Validate() error {
	return dara.Validate(s)
}
