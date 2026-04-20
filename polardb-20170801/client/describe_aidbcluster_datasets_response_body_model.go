// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetContinuationToken() *string
	SetDataServiceId(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetDataServiceId() *string
	SetDatasetId(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetDatasetId() *string
	SetDatasetMode(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetDatasetMode() *string
	SetDatasetType(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetDatasetType() *string
	SetDatasets(v []*DescribeAIDBClusterDatasetsResponseBodyDatasets) *DescribeAIDBClusterDatasetsResponseBody
	GetDatasets() []*DescribeAIDBClusterDatasetsResponseBodyDatasets
	SetFileCount(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetFileCount() *string
	SetIsTruncated(v bool) *DescribeAIDBClusterDatasetsResponseBody
	GetIsTruncated() *bool
	SetNextContinuationToken(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetNextContinuationToken() *string
	SetPageNumber(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetPageSize() *string
	SetRelativeDBClusterId(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetRelativeDBClusterId() *string
	SetRequestId(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetTotalCount() *string
	SetTotalRecords(v string) *DescribeAIDBClusterDatasetsResponseBody
	GetTotalRecords() *string
}

type DescribeAIDBClusterDatasetsResponseBody struct {
	// example:
	//
	// EFSDF-DF-***
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// example:
	//
	// pcs-2zeei***
	DataServiceId *string `json:"DataServiceId,omitempty" xml:"DataServiceId,omitempty"`
	// example:
	//
	// pds-2ze88***
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// stf
	DatasetMode *string `json:"DatasetMode,omitempty" xml:"DatasetMode,omitempty"`
	// example:
	//
	// train
	DatasetType *string                                            `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Datasets    []*DescribeAIDBClusterDatasetsResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	FileCount *string `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// EFSDF-DF-***
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pc-2ze88***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 6
	TotalRecords *string `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeAIDBClusterDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetDataServiceId() *string {
	return s.DataServiceId
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetDatasetMode() *string {
	return s.DatasetMode
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetDatasets() []*DescribeAIDBClusterDatasetsResponseBodyDatasets {
	return s.Datasets
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetFileCount() *string {
	return s.FileCount
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAIDBClusterDatasetsResponseBody) GetTotalRecords() *string {
	return s.TotalRecords
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetContinuationToken(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.ContinuationToken = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetDataServiceId(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.DataServiceId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetDatasetId(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.DatasetId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetDatasetMode(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.DatasetMode = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetDatasetType(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.DatasetType = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetDatasets(v []*DescribeAIDBClusterDatasetsResponseBodyDatasets) *DescribeAIDBClusterDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetFileCount(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.FileCount = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetIsTruncated(v bool) *DescribeAIDBClusterDatasetsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetNextContinuationToken(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.NextContinuationToken = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetPageNumber(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetPageSize(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetRelativeDBClusterId(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetRequestId(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetTotalCount(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) SetTotalRecords(v string) *DescribeAIDBClusterDatasetsResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBody) Validate() error {
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterDatasetsResponseBodyDatasets struct {
	// example:
	//
	// pcs-2ze22***-q7***
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 2845904
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// 2025-11-06T06:50:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// pds-2ze88***
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// train
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// example:
	//
	// train-***.json
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 2025-11-06T06:50:43Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// example:
	//
	// polardb_ai/datasets/train/sft/dataset01/train-***.json
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// Standard
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// stf
	TrainMode *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
}

func (s DescribeAIDBClusterDatasetsResponseBodyDatasets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterDatasetsResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetFileName() *string {
	return s.FileName
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetLastModified() *string {
	return s.LastModified
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetPath() *string {
	return s.Path
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) GetTrainMode() *string {
	return s.TrainMode
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetBucketName(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.BucketName = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetCapacity(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.Capacity = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetCreationTime(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.CreationTime = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetDatasetId(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.DatasetId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetDatasetType(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.DatasetType = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetFileName(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.FileName = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetLastModified(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.LastModified = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetPath(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.Path = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetStorageType(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.StorageType = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) SetTrainMode(v string) *DescribeAIDBClusterDatasetsResponseBodyDatasets {
	s.TrainMode = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponseBodyDatasets) Validate() error {
	return dara.Validate(s)
}
