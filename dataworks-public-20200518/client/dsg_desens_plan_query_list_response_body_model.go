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
	// The pagination information.
	PageData *DsgDesensPlanQueryListResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Struct"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	return dara.Validate(s)
}

type DsgDesensPlanQueryListResponseBodyPageData struct {
	// The information about the data masking rule.
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
	// The number of data masking rules.
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
	return dara.Validate(s)
}

type DsgDesensPlanQueryListResponseBodyPageDataData struct {
	// Indicates whether a watermark is added. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CheckWatermark *bool `json:"CheckWatermark,omitempty" xml:"CheckWatermark,omitempty"`
	// The sensitive field type.
	//
	// example:
	//
	// phone
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The type of the data masking method.
	//
	// example:
	//
	// HASH
	DesenMode *string `json:"DesenMode,omitempty" xml:"DesenMode,omitempty"`
	// The details of the data masking rule.
	DesensPlan *DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan `json:"DesensPlan,omitempty" xml:"DesensPlan,omitempty" type:"Struct"`
	// The data masking rule.
	//
	// example:
	//
	// HASH
	DesensRule *string `json:"DesensRule,omitempty" xml:"DesensRule,omitempty"`
	// The data masking method.
	//
	// example:
	//
	// HASH
	DesensWay *string `json:"DesensWay,omitempty" xml:"DesensWay,omitempty"`
	// The time when the data masking rule was created.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the data masking rule was modified.
	//
	// example:
	//
	// 2024-05-09 15:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the data masking rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The owner of the data masking rule.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the data masking rule.
	//
	// example:
	//
	// phone_hash
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The code of the level-1 data masking scenario to which the rule belongs. Valid values:
	//
	// 	- dataworks_display_desense_code: masking of displayed data in DataStudio and Data Map
	//
	// 	- maxcompute_desense_code: data masking at the MaxCompute compute engine layer
	//
	// 	- maxcompute_new_desense_code: data masking at the MaxCompute compute engine layer (new)
	//
	// 	- hologres_display_desense_code: data masking at the Hologres compute engine layer
	//
	// 	- dataworks_data_integration_desense_code: static data masking in Data Integration
	//
	// 	- dataworks_analysis_desense_code: masking of displayed data in DataAnalysis
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The name of the level-2 data masking scenario to which the data masking rule belongs.
	//
	// example:
	//
	// test_scene
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DsgDesensPlanQueryListResponseBodyPageDataData) Validate() error {
	return dara.Validate(s)
}

type DsgDesensPlanQueryListResponseBodyPageDataDataDesensPlan struct {
	// The type of the data masking rule.
	//
	// example:
	//
	// hash
	DesensPlanType *string `json:"DesensPlanType,omitempty" xml:"DesensPlanType,omitempty"`
	// The parameters for the data masking rule. For more information about the parameters, see the [DsgDesensPlanAddOrUpdate](https://help.aliyun.com/document_detail/2786295.html) API reference.
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
