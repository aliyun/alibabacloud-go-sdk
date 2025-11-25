// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSdlStatisticResponseBody
	GetRequestId() *string
	SetSdlStatisticResp(v *DescribeSdlStatisticResponseBodySdlStatisticResp) *DescribeSdlStatisticResponseBody
	GetSdlStatisticResp() *DescribeSdlStatisticResponseBodySdlStatisticResp
}

type DescribeSdlStatisticResponseBody struct {
	// example:
	//
	// 337A4DBA-8A01-5E9C-99CA-84293E13****
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdlStatisticResp *DescribeSdlStatisticResponseBodySdlStatisticResp `json:"SdlStatisticResp,omitempty" xml:"SdlStatisticResp,omitempty" type:"Struct"`
}

func (s DescribeSdlStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlStatisticResponseBody) GetSdlStatisticResp() *DescribeSdlStatisticResponseBodySdlStatisticResp {
	return s.SdlStatisticResp
}

func (s *DescribeSdlStatisticResponseBody) SetRequestId(v string) *DescribeSdlStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlStatisticResponseBody) SetSdlStatisticResp(v *DescribeSdlStatisticResponseBodySdlStatisticResp) *DescribeSdlStatisticResponseBody {
	s.SdlStatisticResp = v
	return s
}

func (s *DescribeSdlStatisticResponseBody) Validate() error {
	if s.SdlStatisticResp != nil {
		if err := s.SdlStatisticResp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSdlStatisticResponseBodySdlStatisticResp struct {
	SdlAssetTopList       []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList       `json:"SdlAssetTopList,omitempty" xml:"SdlAssetTopList,omitempty" type:"Repeated"`
	SdlDstTopList         []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList         `json:"SdlDstTopList,omitempty" xml:"SdlDstTopList,omitempty" type:"Repeated"`
	SdlEventTypeCountList []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList `json:"SdlEventTypeCountList,omitempty" xml:"SdlEventTypeCountList,omitempty" type:"Repeated"`
}

func (s DescribeSdlStatisticResponseBodySdlStatisticResp) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponseBodySdlStatisticResp) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) GetSdlAssetTopList() []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList {
	return s.SdlAssetTopList
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) GetSdlDstTopList() []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList {
	return s.SdlDstTopList
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) GetSdlEventTypeCountList() []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList {
	return s.SdlEventTypeCountList
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) SetSdlAssetTopList(v []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) *DescribeSdlStatisticResponseBodySdlStatisticResp {
	s.SdlAssetTopList = v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) SetSdlDstTopList(v []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) *DescribeSdlStatisticResponseBodySdlStatisticResp {
	s.SdlDstTopList = v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) SetSdlEventTypeCountList(v []*DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) *DescribeSdlStatisticResponseBodySdlStatisticResp {
	s.SdlEventTypeCountList = v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticResp) Validate() error {
	if s.SdlAssetTopList != nil {
		for _, item := range s.SdlAssetTopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SdlDstTopList != nil {
		for _, item := range s.SdlDstTopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SdlEventTypeCountList != nil {
		for _, item := range s.SdlEventTypeCountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList struct {
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// 116.62.66.XXX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// 0
	TrafficBytes *int64 `json:"TrafficBytes,omitempty" xml:"TrafficBytes,omitempty"`
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) GetTrafficBytes() *int64 {
	return s.TrafficBytes
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) SetAssetType(v string) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList {
	s.AssetType = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) SetPublicIp(v string) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList {
	s.PublicIp = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) SetTrafficBytes(v int64) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList {
	s.TrafficBytes = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlAssetTopList) Validate() error {
	return dara.Validate(s)
}

type DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList struct {
	// example:
	//
	// 47.101.68.XXX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// 0
	TrafficBytes *int64 `json:"TrafficBytes,omitempty" xml:"TrafficBytes,omitempty"`
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) GetTrafficBytes() *int64 {
	return s.TrafficBytes
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) SetPublicIp(v string) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList {
	s.PublicIp = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) SetTrafficBytes(v int64) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList {
	s.TrafficBytes = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlDstTopList) Validate() error {
	return dara.Validate(s)
}

type DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList struct {
	// example:
	//
	// 9
	Count     *string `json:"Count,omitempty" xml:"Count,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) GetCount() *string {
	return s.Count
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) SetCount(v string) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList {
	s.Count = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) SetEventType(v string) *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList {
	s.EventType = &v
	return s
}

func (s *DescribeSdlStatisticResponseBodySdlStatisticRespSdlEventTypeCountList) Validate() error {
	return dara.Validate(s)
}
