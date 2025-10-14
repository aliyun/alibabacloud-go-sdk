// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody
	GetDataSources() []*ListDataSourcesResponseBodyDataSources
	SetMaxResults(v int32) *ListDataSourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSourcesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDataSourcesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataSourcesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListDataSourcesResponseBody
	GetTotalPage() *int32
}

type ListDataSourcesResponseBody struct {
	DataSources []*ListDataSourcesResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1。
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetDataSources() []*ListDataSourcesResponseBodyDataSources {
	return s.DataSources
}

func (s *ListDataSourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataSourcesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListDataSourcesResponseBody) SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBody) SetMaxResults(v int32) *ListDataSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetNextToken(v string) *ListDataSourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetPageNumber(v int32) *ListDataSourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetPageSize(v int32) *ListDataSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetTotalCount(v int32) *ListDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetTotalPage(v int32) *ListDataSourcesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourcesResponseBodyDataSources struct {
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// center。
	DataSourceFrom *string `json:"DataSourceFrom,omitempty" xml:"DataSourceFrom,omitempty"`
	// example:
	//
	// ds-scpfegri73oyoknbc90c。
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// AD_LOG。
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// true。
	DataSourceRecognizeEnabled *bool `json:"DataSourceRecognizeEnabled,omitempty" xml:"DataSourceRecognizeEnabled,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot。
	DataSourceRecognizer *string                                                       `json:"DataSourceRecognizer,omitempty" xml:"DataSourceRecognizer,omitempty"`
	DataSourceReferences []*ListDataSourcesResponseBodyDataSourcesDataSourceReferences `json:"DataSourceReferences,omitempty" xml:"DataSourceReferences,omitempty" type:"Repeated"`
	DataSourceStatus     *string                                                       `json:"DataSourceStatus,omitempty" xml:"DataSourceStatus,omitempty"`
	DataSourceStores     []*ListDataSourcesResponseBodyDataSourcesDataSourceStores     `json:"DataSourceStores,omitempty" xml:"DataSourceStores,omitempty" type:"Repeated"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log_173326*******。
	DataSourceTemplateId *string `json:"DataSourceTemplateId,omitempty" xml:"DataSourceTemplateId,omitempty"`
	// example:
	//
	// custom。
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// audit-activity。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceFrom() *string {
	return s.DataSourceFrom
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceRecognizeEnabled() *bool {
	return s.DataSourceRecognizeEnabled
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceRecognizer() *string {
	return s.DataSourceRecognizer
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceReferences() []*ListDataSourcesResponseBodyDataSourcesDataSourceReferences {
	return s.DataSourceReferences
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceStatus() *string {
	return s.DataSourceStatus
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceStores() []*ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	return s.DataSourceStores
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceTemplateId() *string {
	return s.DataSourceTemplateId
}

func (s *ListDataSourcesResponseBodyDataSources) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourcesResponseBodyDataSources) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListDataSourcesResponseBodyDataSources) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListDataSourcesResponseBodyDataSources) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDataSourcesResponseBodyDataSources) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *ListDataSourcesResponseBodyDataSources) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSourcesResponseBodyDataSources) SetCreateTime(v int64) *ListDataSourcesResponseBodyDataSources {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceFrom(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceFrom = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceId(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceName(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceRecognizeEnabled(v bool) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceRecognizeEnabled = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceRecognizer(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceRecognizer = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceReferences(v []*ListDataSourcesResponseBodyDataSourcesDataSourceReferences) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceReferences = v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceStatus(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceStatus = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceStores(v []*ListDataSourcesResponseBodyDataSourcesDataSourceStores) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceStores = v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceTemplateId(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceTemplateId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceType(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetLogProjectName(v string) *ListDataSourcesResponseBodyDataSources {
	s.LogProjectName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetLogRegionId(v string) *ListDataSourcesResponseBodyDataSources {
	s.LogRegionId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetLogStoreName(v string) *ListDataSourcesResponseBodyDataSources {
	s.LogStoreName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetLogUserId(v int64) *ListDataSourcesResponseBodyDataSources {
	s.LogUserId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetUpdateTime(v int64) *ListDataSourcesResponseBodyDataSources {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) Validate() error {
	if s.DataSourceReferences != nil {
		for _, item := range s.DataSourceReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataSourceStores != nil {
		for _, item := range s.DataSourceStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourcesResponseBodyDataSourcesDataSourceReferences struct {
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log_173326*******。
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSourcesDataSourceReferences) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSourcesDataSourceReferences) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceReferences) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceReferences) SetDataIngestionId(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceReferences {
	s.DataIngestionId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceReferences) Validate() error {
	return dara.Validate(s)
}

type ListDataSourcesResponseBodyDataSourcesDataSourceStores struct {
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// center。
	DataSourceStoreFrom *string `json:"DataSourceStoreFrom,omitempty" xml:"DataSourceStoreFrom,omitempty"`
	// example:
	//
	// di_xxxx_source_1。
	DataSourceStoreId *string `json:"DataSourceStoreId,omitempty" xml:"DataSourceStoreId,omitempty"`
	// example:
	//
	// normal。
	DataSourceStoreStatus *string `json:"DataSourceStoreStatus,omitempty" xml:"DataSourceStoreStatus,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// audit-activity。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSourcesDataSourceStores) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSourcesDataSourceStores) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetCheckTime() *int64 {
	return s.CheckTime
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetDataSourceStoreFrom() *string {
	return s.DataSourceStoreFrom
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetDataSourceStoreId() *string {
	return s.DataSourceStoreId
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetDataSourceStoreStatus() *string {
	return s.DataSourceStoreStatus
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetCheckTime(v int64) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.CheckTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetCreateTime(v int64) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetDataSourceStoreFrom(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.DataSourceStoreFrom = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetDataSourceStoreId(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.DataSourceStoreId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetDataSourceStoreStatus(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.DataSourceStoreStatus = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetLogProjectName(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.LogProjectName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetLogRegionId(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.LogRegionId = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetLogStoreName(v string) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.LogStoreName = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) SetUpdateTime(v int64) *ListDataSourcesResponseBodyDataSourcesDataSourceStores {
	s.UpdateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSourcesDataSourceStores) Validate() error {
	return dara.Validate(s)
}
