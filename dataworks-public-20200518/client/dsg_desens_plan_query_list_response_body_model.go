// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanQueryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DsgDesensPlanQueryListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgDesensPlanQueryListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgDesensPlanQueryListResponseBody
	GetHttpStatusCode() *int32
	SetPageData(v *DsgDesensPlanQueryListResponseBodyPageData) *DsgDesensPlanQueryListResponseBody
	GetPageData() *DsgDesensPlanQueryListResponseBodyPageData
	SetRequestId(v string) *DsgDesensPlanQueryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgDesensPlanQueryListResponseBody
	GetSuccess() *bool
}

type DsgDesensPlanQueryListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The paginated data.
	PageData *DsgDesensPlanQueryListResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Struct"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s DsgDesensPlanQueryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgDesensPlanQueryListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgDesensPlanQueryListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgDesensPlanQueryListResponseBody) GetPageData() *DsgDesensPlanQueryListResponseBodyPageData {
	return s.PageData
}

func (s *DsgDesensPlanQueryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgDesensPlanQueryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgDesensPlanQueryListResponseBody) SetErrorCode(v string) *DsgDesensPlanQueryListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) SetErrorMessage(v string) *DsgDesensPlanQueryListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) SetHttpStatusCode(v int32) *DsgDesensPlanQueryListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) SetPageData(v *DsgDesensPlanQueryListResponseBodyPageData) *DsgDesensPlanQueryListResponseBody {
	s.PageData = v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) SetRequestId(v string) *DsgDesensPlanQueryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) SetSuccess(v bool) *DsgDesensPlanQueryListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBody) Validate() error {
	if s.PageData != nil {
		if err := s.PageData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DsgDesensPlanQueryListResponseBodyPageData struct {
	// The details of the desensitization rules.
	Data []*DsgDesensPlanQueryListResponseBodyPageDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of matching desensitization rules.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DsgDesensPlanQueryListResponseBodyPageData) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponseBodyPageData) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) GetData() []*DsgDesensPlanQueryListResponseBodyPageDataData {
	return s.Data
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) SetData(v []*DsgDesensPlanQueryListResponseBodyPageDataData) *DsgDesensPlanQueryListResponseBodyPageData {
	s.Data = v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) SetPageNumber(v int32) *DsgDesensPlanQueryListResponseBodyPageData {
	s.PageNumber = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) SetPageSize(v int32) *DsgDesensPlanQueryListResponseBodyPageData {
	s.PageSize = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) SetTotalCount(v int32) *DsgDesensPlanQueryListResponseBodyPageData {
	s.TotalCount = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageData) Validate() error {
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

type DsgDesensPlanQueryListResponseBodyPageDataData struct {
	// Indicates whether to add a watermark. Valid values:
	//
	// - true: A watermark is added.
	//
	// - false: No watermark is added.
	//
	// example:
	//
	// true
	CheckWatermark *bool `json:"CheckWatermark,omitempty" xml:"CheckWatermark,omitempty"`
	// The sensitive data type.
	//
	// example:
	//
	// phone
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The desensitization method.
	//
	// example:
	//
	// HASH
	DesenMode *string `json:"DesenMode,omitempty" xml:"DesenMode,omitempty"`
	// The details of the desensitization plan.
	DesensPlan *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan `json:"DesensPlan,omitempty" xml:"DesensPlan,omitempty" type:"Struct"`
	// The desensitization rule.
	//
	// example:
	//
	// HASH
	DesensRule *string `json:"DesensRule,omitempty" xml:"DesensRule,omitempty"`
	// The desensitization method.
	//
	// example:
	//
	// HASH
	DesensWay *string `json:"DesensWay,omitempty" xml:"DesensWay,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the rule was last modified.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the desensitization rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The owner of the desensitization rule.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the desensitization rule.
	//
	// example:
	//
	// phone_hash
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The level-1 desensitization scene code. Valid values:
	//
	// - Desensitization for display in Data Development and Data Map: dataworks_display_desense_code
	//
	// - Desensitization at the MaxCompute engine layer: maxcompute_desense_code
	//
	// - Desensitization at the MaxCompute engine layer (New): maxcompute_new_desense_code
	//
	// - Desensitization at the Hologres engine layer: hologres_display_desense_code
	//
	// - Static desensitization in Data Integration: dataworks_data_integration_desense_code
	//
	// - Desensitization for display in Data Analysis: dataworks_analysis_desense_code
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The name of the level-2 desensitization scene.
	//
	// example:
	//
	// test_scene
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the rule. Valid values:
	//
	// - 0: Inactive.
	//
	// - 1: Active.
	//
	// example:
	//
	// 1
	Status         *int32                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Columns        []*DsgDesensPlanQueryListResponseBodyPageDataDataColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	EmptyNotDesesn *bool                                                    `json:"emptyNotDesesn,omitempty" xml:"emptyNotDesesn,omitempty"`
}

func (s DsgDesensPlanQueryListResponseBodyPageDataData) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponseBodyPageDataData) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetCheckWatermark() *bool {
	return s.CheckWatermark
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetDataType() *string {
	return s.DataType
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetDesenMode() *string {
	return s.DesenMode
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetDesensPlan() *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan {
	return s.DesensPlan
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetDesensRule() *string {
	return s.DesensRule
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetDesensWay() *string {
	return s.DesensWay
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetId() *int64 {
	return s.Id
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetOwner() *string {
	return s.Owner
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetSceneName() *string {
	return s.SceneName
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetColumns() []*DsgDesensPlanQueryListResponseBodyPageDataDataColumns {
	return s.Columns
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) GetEmptyNotDesesn() *bool {
	return s.EmptyNotDesesn
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetCheckWatermark(v bool) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.CheckWatermark = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetDataType(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.DataType = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetDesenMode(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.DesenMode = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetDesensPlan(v *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.DesensPlan = v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetDesensRule(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.DesensRule = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetDesensWay(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.DesensWay = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetGmtCreate(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.GmtCreate = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetGmtModified(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.GmtModified = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetId(v int64) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.Id = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetOwner(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.Owner = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetRuleName(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.RuleName = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetSceneCode(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.SceneCode = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetSceneName(v string) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.SceneName = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetStatus(v int32) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetColumns(v []*DsgDesensPlanQueryListResponseBodyPageDataDataColumns) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.Columns = v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) SetEmptyNotDesesn(v bool) *DsgDesensPlanQueryListResponseBodyPageDataData {
	s.EmptyNotDesesn = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) Validate() error {
	if s.DesensPlan != nil {
		if err := s.DesensPlan.Validate(); err != nil {
			return err
		}
	}
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan struct {
	// The type of the desensitization plan.
	//
	// example:
	//
	// hash
	DesensPlanType *string `json:"DesensPlanType,omitempty" xml:"DesensPlanType,omitempty"`
	// The parameters for the desensitization rule. For details, see the [DsgDesensPlanAddOrUpdate](https://help.aliyun.com/document_detail/2786295.html) operation.
	ExtParam map[string]interface{} `json:"ExtParam,omitempty" xml:"ExtParam,omitempty"`
}

func (s DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) GetDesensPlanType() *string {
	return s.DesensPlanType
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) GetExtParam() map[string]interface{} {
	return s.ExtParam
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) SetDesensPlanType(v string) *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan {
	s.DesensPlanType = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) SetExtParam(v map[string]interface{}) *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan {
	s.ExtParam = v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan) Validate() error {
	return dara.Validate(s)
}

type DsgDesensPlanQueryListResponseBodyPageDataDataColumns struct {
	Column  *string `json:"column,omitempty" xml:"column,omitempty"`
	DbType  *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	Table   *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s DsgDesensPlanQueryListResponseBodyPageDataDataColumns) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListResponseBodyPageDataDataColumns) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) GetColumn() *string {
	return s.Column
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) GetDbType() *string {
	return s.DbType
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) GetProject() *string {
	return s.Project
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) GetTable() *string {
	return s.Table
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) SetColumn(v string) *DsgDesensPlanQueryListResponseBodyPageDataDataColumns {
	s.Column = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) SetDbType(v string) *DsgDesensPlanQueryListResponseBodyPageDataDataColumns {
	s.DbType = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) SetProject(v string) *DsgDesensPlanQueryListResponseBodyPageDataDataColumns {
	s.Project = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) SetTable(v string) *DsgDesensPlanQueryListResponseBodyPageDataDataColumns {
	s.Table = &v
	return s
}

func (s *DsgDesensPlanQueryListResponseBodyPageDataDataColumns) Validate() error {
	return dara.Validate(s)
}
