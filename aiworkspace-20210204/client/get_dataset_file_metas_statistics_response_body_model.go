// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetasStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetasStats(v []*DatasetFileMetasStat) *GetDatasetFileMetasStatisticsResponseBody
	GetDatasetFileMetasStats() []*DatasetFileMetasStat
	SetTotalCount(v int32) *GetDatasetFileMetasStatisticsResponseBody
	GetTotalCount() *int32
	SetRequestId(v string) *GetDatasetFileMetasStatisticsResponseBody
	GetRequestId() *string
}

type GetDatasetFileMetasStatisticsResponseBody struct {
	// The details of the returned aggregation list, including the number of each aggregate item. The list is by default sorted in descending order based on the count number.
	DatasetFileMetasStats []*DatasetFileMetasStat `json:"DatasetFileMetasStats,omitempty" xml:"DatasetFileMetasStats,omitempty" type:"Repeated"`
	// The returned number. Example: the number of metadata records or the number of user-defined tags.
	//
	// example:
	//
	// 73
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDatasetFileMetasStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetasStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetasStatisticsResponseBody) GetDatasetFileMetasStats() []*DatasetFileMetasStat {
	return s.DatasetFileMetasStats
}

func (s *GetDatasetFileMetasStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetDatasetFileMetasStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetFileMetasStatisticsResponseBody) SetDatasetFileMetasStats(v []*DatasetFileMetasStat) *GetDatasetFileMetasStatisticsResponseBody {
	s.DatasetFileMetasStats = v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponseBody) SetTotalCount(v int32) *GetDatasetFileMetasStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponseBody) SetRequestId(v string) *GetDatasetFileMetasStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
