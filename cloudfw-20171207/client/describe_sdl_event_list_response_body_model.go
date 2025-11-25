// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSdlEventListResponseBody
	GetRequestId() *string
	SetSdlEventDetailList(v []*DescribeSdlEventListResponseBodySdlEventDetailList) *DescribeSdlEventListResponseBody
	GetSdlEventDetailList() []*DescribeSdlEventListResponseBodySdlEventDetailList
	SetTotalCount(v int64) *DescribeSdlEventListResponseBody
	GetTotalCount() *int64
}

type DescribeSdlEventListResponseBody struct {
	// example:
	//
	// F06DE24D-6EB9-5F55-B588-7BB946DF****
	RequestId          *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdlEventDetailList []*DescribeSdlEventListResponseBodySdlEventDetailList `json:"SdlEventDetailList,omitempty" xml:"SdlEventDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSdlEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlEventListResponseBody) GetSdlEventDetailList() []*DescribeSdlEventListResponseBodySdlEventDetailList {
	return s.SdlEventDetailList
}

func (s *DescribeSdlEventListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSdlEventListResponseBody) SetRequestId(v string) *DescribeSdlEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlEventListResponseBody) SetSdlEventDetailList(v []*DescribeSdlEventListResponseBodySdlEventDetailList) *DescribeSdlEventListResponseBody {
	s.SdlEventDetailList = v
	return s
}

func (s *DescribeSdlEventListResponseBody) SetTotalCount(v int64) *DescribeSdlEventListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSdlEventListResponseBody) Validate() error {
	if s.SdlEventDetailList != nil {
		for _, item := range s.SdlEventDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSdlEventListResponseBodySdlEventDetailList struct {
	// example:
	//
	// test
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// example:
	//
	// 47.100.102.XXX
	AssetPrivateIp *string `json:"AssetPrivateIp,omitempty" xml:"AssetPrivateIp,omitempty"`
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// Trusted
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// example:
	//
	// Trusted
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 000
	CityId *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// example:
	//
	// cn
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// example:
	//
	// 106.14.74.XXX
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// example:
	//
	// 22
	DstPortList *string `json:"DstPortList,omitempty" xml:"DstPortList,omitempty"`
	// example:
	//
	// 1
	EventCnt *int64 `json:"EventCnt,omitempty" xml:"EventCnt,omitempty"`
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 1735697768
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 1738636157
	LastTime     *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// 3082002f02010004067075626c6963a082002002044c33a756020100020100308200103082000c06082b060102010105000500
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// example:
	//
	// TCP
	ProtoList *string `json:"ProtoList,omitempty" xml:"ProtoList,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ce347a98f41e849188aa51c56b02a****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 0
	ResourceIdType *int32 `json:"ResourceIdType,omitempty" xml:"ResourceIdType,omitempty"`
	// example:
	//
	// 10
	SensitiveDataCnt  *int64    `json:"SensitiveDataCnt,omitempty" xml:"SensitiveDataCnt,omitempty"`
	SensitiveDataList []*string `json:"SensitiveDataList,omitempty" xml:"SensitiveDataList,omitempty" type:"Repeated"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	SensitiveType  *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// example:
	//
	// 104.28.226.XX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// example:
	//
	// 443
	SrcPortList *string `json:"SrcPortList,omitempty" xml:"SrcPortList,omitempty"`
	// example:
	//
	// 0
	TrafficBytes *int64 `json:"TrafficBytes,omitempty" xml:"TrafficBytes,omitempty"`
	// example:
	//
	// b91035dc-8be4-411d-bec5-e6320af9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSdlEventListResponseBodySdlEventDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventListResponseBodySdlEventDetailList) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetAssetPrivateIp() *string {
	return s.AssetPrivateIp
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetCategoryClassId() *string {
	return s.CategoryClassId
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetCityId() *string {
	return s.CityId
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetCountryId() *string {
	return s.CountryId
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetDstPortList() *string {
	return s.DstPortList
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetEventCnt() *int64 {
	return s.EventCnt
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetLocationName() *string {
	return s.LocationName
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetPayload() *string {
	return s.Payload
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetProtoList() *string {
	return s.ProtoList
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetResourceIdType() *int32 {
	return s.ResourceIdType
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSensitiveDataCnt() *int64 {
	return s.SensitiveDataCnt
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSensitiveDataList() []*string {
	return s.SensitiveDataList
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetSrcPortList() *string {
	return s.SrcPortList
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetTrafficBytes() *int64 {
	return s.TrafficBytes
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetAssetName(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.AssetName = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetAssetPrivateIp(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.AssetPrivateIp = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetAssetType(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.AssetType = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetCategoryClassId(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetCategoryName(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.CategoryName = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetCityId(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.CityId = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetCountryId(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.CountryId = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetDstIp(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.DstIp = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetDstPortList(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.DstPortList = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetEventCnt(v int64) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.EventCnt = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetEventLevel(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.EventLevel = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetEventName(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.EventName = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetFirstTime(v int64) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.FirstTime = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetLastTime(v int64) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.LastTime = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetLocationName(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.LocationName = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetPayload(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.Payload = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetProtoList(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.ProtoList = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetRegionId(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.RegionId = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetResourceId(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.ResourceId = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetResourceIdType(v int32) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.ResourceIdType = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSensitiveDataCnt(v int64) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SensitiveDataCnt = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSensitiveDataList(v []*string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SensitiveDataList = v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSensitiveLevel(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSensitiveType(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SensitiveType = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSrcIp(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetSrcPortList(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.SrcPortList = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetTrafficBytes(v int64) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.TrafficBytes = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) SetUuid(v string) *DescribeSdlEventListResponseBodySdlEventDetailList {
	s.Uuid = &v
	return s
}

func (s *DescribeSdlEventListResponseBodySdlEventDetailList) Validate() error {
	return dara.Validate(s)
}
