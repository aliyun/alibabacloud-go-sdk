// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDataServiceListResponseBody
	GetRequestId() *string
	SetResult(v *QueryDataServiceListResponseBodyResult) *QueryDataServiceListResponseBody
	GetResult() *QueryDataServiceListResponseBodyResult
	SetSuccess(v bool) *QueryDataServiceListResponseBody
	GetSuccess() *bool
}

type QueryDataServiceListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *QueryDataServiceListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataServiceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataServiceListResponseBody) GetResult() *QueryDataServiceListResponseBodyResult {
	return s.Result
}

func (s *QueryDataServiceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDataServiceListResponseBody) SetRequestId(v string) *QueryDataServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataServiceListResponseBody) SetResult(v *QueryDataServiceListResponseBodyResult) *QueryDataServiceListResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataServiceListResponseBody) SetSuccess(v bool) *QueryDataServiceListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDataServiceListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataServiceListResponseBodyResult struct {
	// The data service information.
	Data []*QueryDataServiceListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows.
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

func (s QueryDataServiceListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResult) GetData() []*QueryDataServiceListResponseBodyResultData {
	return s.Data
}

func (s *QueryDataServiceListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDataServiceListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDataServiceListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *QueryDataServiceListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryDataServiceListResponseBodyResult) SetData(v []*QueryDataServiceListResponseBodyResultData) *QueryDataServiceListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetPageNum(v int32) *QueryDataServiceListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetPageSize(v int32) *QueryDataServiceListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetTotalNum(v int32) *QueryDataServiceListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetTotalPages(v int32) *QueryDataServiceListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) Validate() error {
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

type QueryDataServiceListResponseBodyResultData struct {
	// The model of the data service in JSON format.
	Content *QueryDataServiceListResponseBodyResultDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the creator.
	//
	// example:
	//
	// 7cb94cd48701
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// Zhang Zilong
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The cube ID.
	//
	// example:
	//
	// d14e7448-0eb3-40d3-9375-4afef8de29fd
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// Test dataset
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// The description.
	//
	// example:
	//
	// For testing
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-05-18 14:00:02.0
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-03-21 18:02:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID of the modifier.
	//
	// example:
	//
	// 7cb94cd48701
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The name of the modifier.
	//
	// example:
	//
	// Zhang Zilong
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// The service name.
	//
	// example:
	//
	// Test report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner.
	//
	// example:
	//
	// 862801339
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// Li Si
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The unique ID of the data service.
	//
	// example:
	//
	// dtsuq3i31f5j8v848b
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// Test workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultData) GetContent() *QueryDataServiceListResponseBodyResultDataContent {
	return s.Content
}

func (s *QueryDataServiceListResponseBodyResultData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryDataServiceListResponseBodyResultData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *QueryDataServiceListResponseBodyResultData) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryDataServiceListResponseBodyResultData) GetCubeName() *string {
	return s.CubeName
}

func (s *QueryDataServiceListResponseBodyResultData) GetDesc() *string {
	return s.Desc
}

func (s *QueryDataServiceListResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryDataServiceListResponseBodyResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryDataServiceListResponseBodyResultData) GetModifierId() *string {
	return s.ModifierId
}

func (s *QueryDataServiceListResponseBodyResultData) GetModifierName() *string {
	return s.ModifierName
}

func (s *QueryDataServiceListResponseBodyResultData) GetName() *string {
	return s.Name
}

func (s *QueryDataServiceListResponseBodyResultData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryDataServiceListResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryDataServiceListResponseBodyResultData) GetSid() *string {
	return s.Sid
}

func (s *QueryDataServiceListResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryDataServiceListResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryDataServiceListResponseBodyResultData) SetContent(v *QueryDataServiceListResponseBodyResultDataContent) *QueryDataServiceListResponseBodyResultData {
	s.Content = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCreatorId(v string) *QueryDataServiceListResponseBodyResultData {
	s.CreatorId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCreatorName(v string) *QueryDataServiceListResponseBodyResultData {
	s.CreatorName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCubeId(v string) *QueryDataServiceListResponseBodyResultData {
	s.CubeId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCubeName(v string) *QueryDataServiceListResponseBodyResultData {
	s.CubeName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetDesc(v string) *QueryDataServiceListResponseBodyResultData {
	s.Desc = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetGmtCreate(v string) *QueryDataServiceListResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetGmtModified(v string) *QueryDataServiceListResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetModifierId(v string) *QueryDataServiceListResponseBodyResultData {
	s.ModifierId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetModifierName(v string) *QueryDataServiceListResponseBodyResultData {
	s.ModifierName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetName(v string) *QueryDataServiceListResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetOwnerId(v string) *QueryDataServiceListResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetOwnerName(v string) *QueryDataServiceListResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetSid(v string) *QueryDataServiceListResponseBodyResultData {
	s.Sid = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetWorkspaceId(v string) *QueryDataServiceListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetWorkspaceName(v string) *QueryDataServiceListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataServiceListResponseBodyResultDataContent struct {
	// The cube ID.
	//
	// example:
	//
	// 56f9f34a-bdba-496a-91a3-a18b1ff73a80
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// Test dataset
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// The detail or aggregate type.
	//
	// example:
	//
	// true
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The request parameter information.
	Filter *QueryDataServiceListResponseBodyResultDataContentFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The returned information.
	ReturnFields []*QueryDataServiceListResponseBodyResultDataContentReturnFields `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty" type:"Repeated"`
}

func (s QueryDataServiceListResponseBodyResultDataContent) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContent) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContent) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryDataServiceListResponseBodyResultDataContent) GetCubeName() *string {
	return s.CubeName
}

func (s *QueryDataServiceListResponseBodyResultDataContent) GetDetail() *bool {
	return s.Detail
}

func (s *QueryDataServiceListResponseBodyResultDataContent) GetFilter() *QueryDataServiceListResponseBodyResultDataContentFilter {
	return s.Filter
}

func (s *QueryDataServiceListResponseBodyResultDataContent) GetReturnFields() []*QueryDataServiceListResponseBodyResultDataContentReturnFields {
	return s.ReturnFields
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetCubeId(v string) *QueryDataServiceListResponseBodyResultDataContent {
	s.CubeId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetCubeName(v string) *QueryDataServiceListResponseBodyResultDataContent {
	s.CubeName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetDetail(v bool) *QueryDataServiceListResponseBodyResultDataContent {
	s.Detail = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetFilter(v *QueryDataServiceListResponseBodyResultDataContentFilter) *QueryDataServiceListResponseBodyResultDataContent {
	s.Filter = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetReturnFields(v []*QueryDataServiceListResponseBodyResultDataContentReturnFields) *QueryDataServiceListResponseBodyResultDataContent {
	s.ReturnFields = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.ReturnFields != nil {
		for _, item := range s.ReturnFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDataServiceListResponseBodyResultDataContentFilter struct {
	// The combined conditions.
	//
	// example:
	//
	// [{
	//
	// 	"field": {
	//
	// 		"fid": "1c1f88cb7d",
	//
	// 		"granularity": "yearRegion",
	//
	// 		"dataType": "datetime",
	//
	// 		"column": "s_date",
	//
	// 		"caption": "Date (year)",
	//
	// 		"type": "dimension"
	//
	// 	},
	//
	// 	"defaultValue": ["2024"],
	//
	// 	"alias": "s_date",
	//
	// 	"type": "basic",
	//
	// 	"operator": "equal-to",
	//
	// 	"essential": true
	//
	// }]
	Filters []map[string]interface{} `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The logical relationship between multiple SQL text keywords.
	//
	// - **or**: OR.
	//
	// - **and**: AND.
	//
	// example:
	//
	// and
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// The type.
	//
	// - basic: Basic.
	//
	// - combined: Complex.
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentFilter) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) GetFilters() []map[string]interface{} {
	return s.Filters
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) GetType() *string {
	return s.Type
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetFilters(v []map[string]interface{}) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.Filters = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetLogicalOperator(v string) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.LogicalOperator = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetType(v string) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.Type = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) Validate() error {
	return dara.Validate(s)
}

type QueryDataServiceListResponseBodyResultDataContentReturnFields struct {
	// The aggregate operation (aggregation), such as SUM, AVG, and MAX.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The field parameter name.
	//
	// example:
	//
	// s_number
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The description of the returned field.
	//
	// example:
	//
	// The theme configuration already exists
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The corresponding cube field information.
	Field *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField `json:"Field,omitempty" xml:"Field,omitempty" type:"Struct"`
	// The sort order.
	//
	// - asc: Ascending.
	//
	// - desc: Descending.
	//
	// - no: Unsorted.
	//
	// example:
	//
	// no
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFields) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFields) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) GetAggregator() *string {
	return s.Aggregator
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) GetAlias() *string {
	return s.Alias
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) GetDesc() *string {
	return s.Desc
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) GetField() *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	return s.Field
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) GetOrderby() *string {
	return s.Orderby
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetAggregator(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Aggregator = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetAlias(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Alias = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetDesc(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Desc = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetField(v *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Field = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetOrderby(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Orderby = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) Validate() error {
	if s.Field != nil {
		if err := s.Field.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataServiceListResponseBodyResultDataContentReturnFieldsField struct {
	// The display name in the cube model, which can be in Chinese or English.
	//
	// example:
	//
	// Date (year)
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The corresponding physical field name.
	//
	// example:
	//
	// shid_star
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The data type.
	//
	// - number: Number.
	//
	// - string: String.
	//
	// - date: Date.
	//
	// - datetime: Datetime.
	//
	// - time: Time.
	//
	// - geographic: Geographic.
	//
	// - boolean: Boolean.
	//
	// - url: URL.
	//
	// example:
	//
	// datetime
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The unique ID of the original field.
	//
	// example:
	//
	// 1c1f88cb7d
	Fid *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	// The supported granularity. This property is included in date and geographic dimensions.
	//
	// example:
	//
	// yearRegion
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The unique name of the cube field, which is primarily used to uniquely locate the returned result.
	//
	// example:
	//
	// sss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type.
	//
	// - Dimension: Dimension.
	//
	// - Measure: Measure.
	//
	// example:
	//
	// dimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetCaption() *string {
	return s.Caption
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetColumn() *string {
	return s.Column
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetDataType() *string {
	return s.DataType
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetFid() *string {
	return s.Fid
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetName() *string {
	return s.Name
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GetType() *string {
	return s.Type
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetCaption(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Caption = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetColumn(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Column = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetDataType(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.DataType = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetFid(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Fid = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetGranularity(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Granularity = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetName(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetType(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Type = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) Validate() error {
	return dara.Validate(s)
}
