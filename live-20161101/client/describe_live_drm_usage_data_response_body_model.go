// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDrmUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrmUsageData(v *DescribeLiveDrmUsageDataResponseBodyDrmUsageData) *DescribeLiveDrmUsageDataResponseBody
	GetDrmUsageData() *DescribeLiveDrmUsageDataResponseBodyDrmUsageData
	SetRequestId(v string) *DescribeLiveDrmUsageDataResponseBody
	GetRequestId() *string
}

type DescribeLiveDrmUsageDataResponseBody struct {
	// The usage of the DRM encryption service at each time interval.
	DrmUsageData *DescribeLiveDrmUsageDataResponseBodyDrmUsageData `json:"DrmUsageData,omitempty" xml:"DrmUsageData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 91FC2D9D-B042-4634-8A5C-7B8E7482C22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDrmUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDrmUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDrmUsageDataResponseBody) GetDrmUsageData() *DescribeLiveDrmUsageDataResponseBodyDrmUsageData {
	return s.DrmUsageData
}

func (s *DescribeLiveDrmUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDrmUsageDataResponseBody) SetDrmUsageData(v *DescribeLiveDrmUsageDataResponseBodyDrmUsageData) *DescribeLiveDrmUsageDataResponseBody {
	s.DrmUsageData = v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBody) SetRequestId(v string) *DescribeLiveDrmUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBody) Validate() error {
	if s.DrmUsageData != nil {
		if err := s.DrmUsageData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDrmUsageDataResponseBodyDrmUsageData struct {
	DataModule []*DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDrmUsageDataResponseBodyDrmUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDrmUsageDataResponseBodyDrmUsageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageData) GetDataModule() []*DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	return s.DataModule
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageData) SetDataModule(v []*DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) *DescribeLiveDrmUsageDataResponseBodyDrmUsageData {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageData) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule struct {
	// The number of times DRM-encrypted live streams are requested.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name. If the value of SplitBy includes domain, the returned data is grouped by domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The DRM type. If the value of SplitBy includes drm_type, the returned data is grouped by DRM type.
	//
	// example:
	//
	// Widevine
	DrmType *string `json:"DrmType,omitempty" xml:"DrmType,omitempty"`
	// The ID of the region. If the value of SplitBy includes region, the returned data is grouped by region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-05-01T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GetDrmType() *string {
	return s.DrmType
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) SetCount(v int64) *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	s.Count = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) SetDomain(v string) *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) SetDrmType(v string) *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	s.DrmType = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) SetRegion(v string) *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	s.Region = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) SetTimeStamp(v string) *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDrmUsageDataResponseBodyDrmUsageDataDataModule) Validate() error {
	return dara.Validate(s)
}
