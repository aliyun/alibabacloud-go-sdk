// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodUserResourcePackageResponseBody
	GetRequestId() *string
	SetResourcePackageInfos(v *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) *DescribeVodUserResourcePackageResponseBody
	GetResourcePackageInfos() *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos
}

type DescribeVodUserResourcePackageResponseBody struct {
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePackageInfos *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeVodUserResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserResourcePackageResponseBody) GetResourcePackageInfos() *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos {
	return s.ResourcePackageInfos
}

func (s *DescribeVodUserResourcePackageResponseBody) SetRequestId(v string) *DescribeVodUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) *DescribeVodUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) GetResourcePackageInfo() []*DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	return s.ResourcePackageInfo
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrCapacity  *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InitCapacity  *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacity() *string {
	return s.CurrCapacity
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacity() *string {
	return s.InitCapacity
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetEndTime(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStartTime(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) Validate() error {
	return dara.Validate(s)
}
