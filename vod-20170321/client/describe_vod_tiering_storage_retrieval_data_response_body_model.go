// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTieringStorageRetrievalDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodTieringStorageRetrievalDataResponseBody
	GetRequestId() *string
	SetRetrievalData(v []*DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) *DescribeVodTieringStorageRetrievalDataResponseBody
	GetRetrievalData() []*DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData
}

type DescribeVodTieringStorageRetrievalDataResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D73936****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data retrieval information.
	RetrievalData []*DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData `json:"RetrievalData,omitempty" xml:"RetrievalData,omitempty" type:"Repeated"`
}

func (s DescribeVodTieringStorageRetrievalDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageRetrievalDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBody) GetRetrievalData() []*DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	return s.RetrievalData
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBody) SetRequestId(v string) *DescribeVodTieringStorageRetrievalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBody) SetRetrievalData(v []*DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) *DescribeVodTieringStorageRetrievalDataResponseBody {
	s.RetrievalData = v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData struct {
	// The retrieved Cold Archive data in the bulk mode.
	//
	// example:
	//
	// 123
	CABulkRetrievalData *int64 `json:"CABulkRetrievalData,omitempty" xml:"CABulkRetrievalData,omitempty"`
	// The retrieved Cold Archive data in the expedited mode.
	//
	// example:
	//
	// 123
	CAHighPriorRetrievalData *int64 `json:"CAHighPriorRetrievalData,omitempty" xml:"CAHighPriorRetrievalData,omitempty"`
	// The retrieved Cold Archive data in the standard mode.
	//
	// example:
	//
	// 123
	CAStdRetrievalData *int64 `json:"CAStdRetrievalData,omitempty" xml:"CAStdRetrievalData,omitempty"`
	// The storage region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The data retrieval information.
	//
	// example:
	//
	// 1234
	RetrievalData *int64 `json:"RetrievalData,omitempty" xml:"RetrievalData,omitempty"`
	// The storage type.
	//
	// example:
	//
	// IA
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The timestamp of the returned data. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-02T10:20:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GoString() string {
	return s.String()
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetCABulkRetrievalData() *int64 {
	return s.CABulkRetrievalData
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetCAHighPriorRetrievalData() *int64 {
	return s.CAHighPriorRetrievalData
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetCAStdRetrievalData() *int64 {
	return s.CAStdRetrievalData
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetRetrievalData() *int64 {
	return s.RetrievalData
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetCABulkRetrievalData(v int64) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.CABulkRetrievalData = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetCAHighPriorRetrievalData(v int64) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.CAHighPriorRetrievalData = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetCAStdRetrievalData(v int64) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.CAStdRetrievalData = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetRegion(v string) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.Region = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetRetrievalData(v int64) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.RetrievalData = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetStorageClass(v string) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.StorageClass = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) SetTimeStamp(v string) *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodTieringStorageRetrievalDataResponseBodyRetrievalData) Validate() error {
	return dara.Validate(s)
}
