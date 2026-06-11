// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDatasetListResponseBody
	GetRequestId() *string
	SetResult(v *QueryDatasetListResponseBodyResult) *QueryDatasetListResponseBody
	GetResult() *QueryDatasetListResponseBodyResult
	SetSuccess(v bool) *QueryDatasetListResponseBody
	GetSuccess() *bool
}

type QueryDatasetListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The paged results of the dataset list. The details of the datasets are returned in the Data parameter.
	Result *QueryDatasetListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetListResponseBody) GetResult() *QueryDatasetListResponseBodyResult {
	return s.Result
}

func (s *QueryDatasetListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDatasetListResponseBody) SetRequestId(v string) *QueryDatasetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetListResponseBody) SetResult(v *QueryDatasetListResponseBodyResult) *QueryDatasetListResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetListResponseBody) SetSuccess(v bool) *QueryDatasetListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDatasetListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDatasetListResponseBodyResult struct {
	// The details of the dataset list.
	Data []*QueryDatasetListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page specified in the request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryDatasetListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResult) GetData() []*QueryDatasetListResponseBodyResultData {
	return s.Data
}

func (s *QueryDatasetListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDatasetListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDatasetListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryDatasetListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryDatasetListResponseBodyResult) SetData(v []*QueryDatasetListResponseBodyResultData) *QueryDatasetListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetPageNum(v int32) *QueryDatasetListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetPageSize(v int32) *QueryDatasetListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetTotalNum(v int32) *QueryDatasetListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetTotalPages(v int32) *QueryDatasetListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDatasetListResponseBodyResultData struct {
	// The time when the dataset was created.
	//
	// example:
	//
	// 2020-11-02 10:36:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data source to which the dataset belongs.
	DataSource *QueryDatasetListResponseBodyResultDataDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The dataset ID.
	//
	// example:
	//
	// 5820f58c-c734-4d8a-baf1-7979af4f****
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// company_sales_record_copy12
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description of the dataset.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the folder in which the dataset is located.
	Directory *QueryDatasetListResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The time when the dataset was last modified.
	//
	// example:
	//
	// 2020-11-02 10:36:05
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Indicates whether extraction-based acceleration is enabled. Valid values:
	//
	// - true: Enabled
	//
	// - false: Disabled
	//
	// example:
	//
	// true
	OpenOfflineAcceleration *bool `json:"OpenOfflineAcceleration,omitempty" xml:"OpenOfflineAcceleration,omitempty"`
	// The user ID of the dataset owner in Quick BI.
	//
	// example:
	//
	// 136516262323****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The nickname of the dataset owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Indicates whether row-level permissions are enabled. Valid values:
	//
	// - true: Enabled
	//
	// - false: Disabled
	//
	// example:
	//
	// true
	RowLevel *bool `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad06adf
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDatasetListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryDatasetListResponseBodyResultData) GetDataSource() *QueryDatasetListResponseBodyResultDataDataSource {
	return s.DataSource
}

func (s *QueryDatasetListResponseBodyResultData) GetDatasetId() *string {
	return s.DatasetId
}

func (s *QueryDatasetListResponseBodyResultData) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryDatasetListResponseBodyResultData) GetDescription() *string {
	return s.Description
}

func (s *QueryDatasetListResponseBodyResultData) GetDirectory() *QueryDatasetListResponseBodyResultDataDirectory {
	return s.Directory
}

func (s *QueryDatasetListResponseBodyResultData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *QueryDatasetListResponseBodyResultData) GetOpenOfflineAcceleration() *bool {
	return s.OpenOfflineAcceleration
}

func (s *QueryDatasetListResponseBodyResultData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryDatasetListResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryDatasetListResponseBodyResultData) GetRowLevel() *bool {
	return s.RowLevel
}

func (s *QueryDatasetListResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryDatasetListResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryDatasetListResponseBodyResultData) SetCreateTime(v string) *QueryDatasetListResponseBodyResultData {
	s.CreateTime = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDataSource(v *QueryDatasetListResponseBodyResultDataDataSource) *QueryDatasetListResponseBodyResultData {
	s.DataSource = v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDatasetId(v string) *QueryDatasetListResponseBodyResultData {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDatasetName(v string) *QueryDatasetListResponseBodyResultData {
	s.DatasetName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDescription(v string) *QueryDatasetListResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDirectory(v *QueryDatasetListResponseBodyResultDataDirectory) *QueryDatasetListResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetModifyTime(v string) *QueryDatasetListResponseBodyResultData {
	s.ModifyTime = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetOpenOfflineAcceleration(v bool) *QueryDatasetListResponseBodyResultData {
	s.OpenOfflineAcceleration = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetOwnerId(v string) *QueryDatasetListResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetOwnerName(v string) *QueryDatasetListResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetRowLevel(v bool) *QueryDatasetListResponseBodyResultData {
	s.RowLevel = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetWorkspaceId(v string) *QueryDatasetListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetWorkspaceName(v string) *QueryDatasetListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDatasetListResponseBodyResultDataDataSource struct {
	// The data source ID.
	//
	// example:
	//
	// 261b252d-c3c3-498a-a0a7-5d1ec6cd****
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// test
	DsName *string `json:"DsName,omitempty" xml:"DsName,omitempty"`
	// The data source type.
	//
	// example:
	//
	// mysql
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
}

func (s QueryDatasetListResponseBodyResultDataDataSource) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultDataDataSource) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) GetDsId() *string {
	return s.DsId
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) GetDsName() *string {
	return s.DsName
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) GetDsType() *string {
	return s.DsType
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsId(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsName(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsType(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsType = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetListResponseBodyResultDataDirectory struct {
	// The folder ID.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The folder name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the folder path.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The name of the folder path.
	//
	// example:
	//
	// test
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryDatasetListResponseBodyResultDataDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) GetId() *string {
	return s.Id
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) GetName() *string {
	return s.Name
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetId(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetName(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetPathId(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetPathName(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) Validate() error {
	return dara.Validate(s)
}
