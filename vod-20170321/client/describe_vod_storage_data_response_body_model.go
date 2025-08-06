// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStorageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodStorageDataResponseBody
	GetDataInterval() *string
	SetRequestId(v string) *DescribeVodStorageDataResponseBody
	GetRequestId() *string
	SetStorageData(v *DescribeVodStorageDataResponseBodyStorageData) *DescribeVodStorageDataResponseBody
	GetStorageData() *DescribeVodStorageDataResponseBodyStorageData
}

type DescribeVodStorageDataResponseBody struct {
	// The time granularity at which the data was queried. Valid values:
	//
	// 	- **hour**
	//
	// 	- **day**
	//
	// example:
	//
	// day
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-****-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage usage data returned.
	StorageData *DescribeVodStorageDataResponseBodyStorageData `json:"StorageData,omitempty" xml:"StorageData,omitempty" type:"Struct"`
}

func (s DescribeVodStorageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStorageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodStorageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodStorageDataResponseBody) GetStorageData() *DescribeVodStorageDataResponseBodyStorageData {
	return s.StorageData
}

func (s *DescribeVodStorageDataResponseBody) SetDataInterval(v string) *DescribeVodStorageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodStorageDataResponseBody) SetRequestId(v string) *DescribeVodStorageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodStorageDataResponseBody) SetStorageData(v *DescribeVodStorageDataResponseBodyStorageData) *DescribeVodStorageDataResponseBody {
	s.StorageData = v
	return s
}

func (s *DescribeVodStorageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodStorageDataResponseBodyStorageData struct {
	StorageDataItem []*DescribeVodStorageDataResponseBodyStorageDataStorageDataItem `json:"StorageDataItem,omitempty" xml:"StorageDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodStorageDataResponseBodyStorageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStorageDataResponseBodyStorageData) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBodyStorageData) GetStorageDataItem() []*DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	return s.StorageDataItem
}

func (s *DescribeVodStorageDataResponseBodyStorageData) SetStorageDataItem(v []*DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) *DescribeVodStorageDataResponseBodyStorageData {
	s.StorageDataItem = v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodStorageDataResponseBodyStorageDataStorageDataItem struct {
	// The outbound traffic. Unit: bytes. The outbound traffic is generated when videos are directly downloaded or played from OSS buckets without Alibaba Cloud CDN acceleration.
	//
	// example:
	//
	// 111111
	NetworkOut *string `json:"NetworkOut,omitempty" xml:"NetworkOut,omitempty"`
	// The detailed usage data of storage-related resources. Unit: bytes.
	//
	// example:
	//
	// 111111
	StorageUtilization *string `json:"StorageUtilization,omitempty" xml:"StorageUtilization,omitempty"`
	// The timestamp of the returned data. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-01T15:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) GetNetworkOut() *string {
	return s.NetworkOut
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) GetStorageUtilization() *string {
	return s.StorageUtilization
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetNetworkOut(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.NetworkOut = &v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetStorageUtilization(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.StorageUtilization = &v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetTimeStamp(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) Validate() error {
	return dara.Validate(s)
}
