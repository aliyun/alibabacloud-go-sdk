// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleDataListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSampleDataListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleDataListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSampleDataListResponseBodyResultObject) *DescribeSampleDataListResponseBody
	GetResultObject() []*DescribeSampleDataListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSampleDataListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSampleDataListResponseBody
	GetTotalPage() *int32
}

type DescribeSampleDataListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeSampleDataListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSampleDataListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDataListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleDataListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleDataListResponseBody) GetResultObject() []*DescribeSampleDataListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleDataListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSampleDataListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSampleDataListResponseBody) SetRequestId(v string) *DescribeSampleDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDataListResponseBody) SetCurrentPage(v int32) *DescribeSampleDataListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleDataListResponseBody) SetPageSize(v int32) *DescribeSampleDataListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleDataListResponseBody) SetResultObject(v []*DescribeSampleDataListResponseBodyResultObject) *DescribeSampleDataListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleDataListResponseBody) SetTotalItem(v int32) *DescribeSampleDataListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSampleDataListResponseBody) SetTotalPage(v int32) *DescribeSampleDataListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSampleDataListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSampleDataListResponseBodyResultObject struct {
	// Classification type, binary or multi-class.
	//
	// example:
	//
	// 二分类
	ClassificationType *string `json:"classificationType,omitempty" xml:"classificationType,omitempty"`
	// Criterion value for sample data calculation
	//
	// example:
	//
	// {"正样本":"1"，"负样本":"1"}
	DataDistributed *string `json:"dataDistributed,omitempty" xml:"dataDistributed,omitempty"`
	// First row of sample data. Used to define the values of each column.
	//
	// example:
	//
	// 17700000000
	DataTitle *string `json:"dataTitle,omitempty" xml:"dataTitle,omitempty"`
	// Deletion tag.
	//
	// example:
	//
	// DELETE
	DeleteTag *string `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Name
	//
	// example:
	//
	// 注册样本
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Number of normal samples
	//
	// example:
	//
	// 999
	NormalSize *int64 `json:"normalSize,omitempty" xml:"normalSize,omitempty"`
	// Recall configuration
	//
	// example:
	//
	// {"variables":"a,b,c"}
	RecallConfig *string `json:"recallConfig,omitempty" xml:"recallConfig,omitempty"`
	// Number of risk samples
	//
	// example:
	//
	// 1
	RiskSize *int64 `json:"riskSize,omitempty" xml:"riskSize,omitempty"`
	// Specified risk value
	//
	// example:
	//
	// black
	RiskValue *string `json:"riskValue,omitempty" xml:"riskValue,omitempty"`
	// Sample label details
	//
	// example:
	//
	// [{"type":"positive","size":"2000","value":1},{"type":"negative","size":1900,"value":0}]
	SampleLabelDetail *string `json:"sampleLabelDetail,omitempty" xml:"sampleLabelDetail,omitempty"`
	// Sample size
	//
	// example:
	//
	// 1000
	SampleSize *int64 `json:"sampleSize,omitempty" xml:"sampleSize,omitempty"`
	// Scene code
	//
	// example:
	//
	// account_abuse_detection
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// Status.
	//
	// example:
	//
	// CREATE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Storage path
	//
	// example:
	//
	// saf/de/sample/3dc2spspHKq4G3YI9d08
	StorePath *string `json:"storePath,omitempty" xml:"storePath,omitempty"`
	// Storage type
	//
	// example:
	//
	// OSS
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
	// Whether recall is supported
	//
	// example:
	//
	// true
	SupportRecall *string `json:"supportRecall,omitempty" xml:"supportRecall,omitempty"`
	// User UID
	//
	// example:
	//
	// 1519714049632764
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSampleDataListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetClassificationType() *string {
	return s.ClassificationType
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetDataDistributed() *string {
	return s.DataDistributed
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetDataTitle() *string {
	return s.DataTitle
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetDeleteTag() *string {
	return s.DeleteTag
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetNormalSize() *int64 {
	return s.NormalSize
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetRecallConfig() *string {
	return s.RecallConfig
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetRiskSize() *int64 {
	return s.RiskSize
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetRiskValue() *string {
	return s.RiskValue
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetSampleLabelDetail() *string {
	return s.SampleLabelDetail
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetSampleSize() *int64 {
	return s.SampleSize
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetScene() *string {
	return s.Scene
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetStorePath() *string {
	return s.StorePath
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetStoreType() *string {
	return s.StoreType
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetSupportRecall() *string {
	return s.SupportRecall
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSampleDataListResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetClassificationType(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.ClassificationType = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetDataDistributed(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.DataDistributed = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetDataTitle(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.DataTitle = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetDeleteTag(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.DeleteTag = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetDescription(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetGmtModified(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetId(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetName(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetNormalSize(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.NormalSize = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetRecallConfig(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.RecallConfig = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetRiskSize(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.RiskSize = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetRiskValue(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.RiskValue = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetSampleLabelDetail(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.SampleLabelDetail = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetSampleSize(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.SampleSize = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetScene(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.Scene = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetStatus(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetStorePath(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.StorePath = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetStoreType(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.StoreType = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetSupportRecall(v string) *DescribeSampleDataListResponseBodyResultObject {
	s.SupportRecall = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetUserId(v int64) *DescribeSampleDataListResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) SetVersion(v int32) *DescribeSampleDataListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSampleDataListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
