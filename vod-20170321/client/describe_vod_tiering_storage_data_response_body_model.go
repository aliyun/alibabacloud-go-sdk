// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodTieringStorageDataResponseBody
	GetRequestId() *string
	SetStorageData(v []*DescribeVodTieringStorageDataResponseBodyStorageData) *DescribeVodTieringStorageDataResponseBody
	GetStorageData() []*DescribeVodTieringStorageDataResponseBodyStorageData
}

type DescribeVodTieringStorageDataResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage usage data returned.
	StorageData []*DescribeVodTieringStorageDataResponseBodyStorageData `json:"StorageData,omitempty" xml:"StorageData,omitempty" type:"Repeated"`
}

func (s DescribeVodTieringStorageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodTieringStorageDataResponseBody) GetStorageData() []*DescribeVodTieringStorageDataResponseBodyStorageData {
	return s.StorageData
}

func (s *DescribeVodTieringStorageDataResponseBody) SetRequestId(v string) *DescribeVodTieringStorageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBody) SetStorageData(v []*DescribeVodTieringStorageDataResponseBodyStorageData) *DescribeVodTieringStorageDataResponseBody {
	s.StorageData = v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBody) Validate() error {
	if s.StorageData != nil {
		for _, item := range s.StorageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodTieringStorageDataResponseBodyStorageData struct {
	// The data that is stored less than a month. Unit: bytes.
	//
	// example:
	//
	// 123
	LessthanMonthDatasize *int64 `json:"LessthanMonthDatasize,omitempty" xml:"LessthanMonthDatasize,omitempty"`
	// The region in which data is queried.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The storage type.
	//
	// example:
	//
	// IA
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The storage usage. Unit: bytes.
	//
	// example:
	//
	// 1234
	StorageUtilization *int64 `json:"StorageUtilization,omitempty" xml:"StorageUtilization,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-29T01:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodTieringStorageDataResponseBodyStorageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageDataResponseBodyStorageData) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) GetLessthanMonthDatasize() *int64 {
	return s.LessthanMonthDatasize
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) GetStorageUtilization() *int64 {
	return s.StorageUtilization
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) SetLessthanMonthDatasize(v int64) *DescribeVodTieringStorageDataResponseBodyStorageData {
	s.LessthanMonthDatasize = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) SetRegion(v string) *DescribeVodTieringStorageDataResponseBodyStorageData {
	s.Region = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) SetStorageClass(v string) *DescribeVodTieringStorageDataResponseBodyStorageData {
	s.StorageClass = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) SetStorageUtilization(v int64) *DescribeVodTieringStorageDataResponseBodyStorageData {
	s.StorageUtilization = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) SetTimeStamp(v string) *DescribeVodTieringStorageDataResponseBodyStorageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodTieringStorageDataResponseBodyStorageData) Validate() error {
	return dara.Validate(s)
}
