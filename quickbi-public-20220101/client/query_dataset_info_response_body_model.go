// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDatasetInfoResponseBody
	GetRequestId() *string
	SetResult(v *QueryDatasetInfoResponseBodyResult) *QueryDatasetInfoResponseBody
	GetResult() *QueryDatasetInfoResponseBodyResult
	SetSuccess(v bool) *QueryDatasetInfoResponseBody
	GetSuccess() *bool
}

type QueryDatasetInfoResponseBody struct {
	// Whether the operation is successfully returned. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The dataset information.
	Result *QueryDatasetInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The unique ID of the dataset.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetInfoResponseBody) GetResult() *QueryDatasetInfoResponseBodyResult {
	return s.Result
}

func (s *QueryDatasetInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDatasetInfoResponseBody) SetRequestId(v string) *QueryDatasetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetInfoResponseBody) SetResult(v *QueryDatasetInfoResponseBodyResult) *QueryDatasetInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetInfoResponseBody) SetSuccess(v bool) *QueryDatasetInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDatasetInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDatasetInfoResponseBodyResult struct {
	// The unique ID of the dataset.
	CubeTableList []*QueryDatasetInfoResponseBodyResultCubeTableList `json:"CubeTableList,omitempty" xml:"CubeTableList,omitempty" type:"Repeated"`
	// The unique ID of the workspace to which the dataset belongs.
	//
	// example:
	//
	// false
	CustimzeSql *bool `json:"CustimzeSql,omitempty" xml:"CustimzeSql,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- mysql
	//
	// 	- odps
	//
	// 	- oracle
	//
	// 	- ... Data source types supported by Quick BI such as
	//
	// example:
	//
	// a201c85c-******
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The user ID of the dataset owner in the Quick BI.
	//
	// example:
	//
	// opds_40
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// A list of all dimensions in the dataset.
	DimensionList []*QueryDatasetInfoResponseBodyResultDimensionList `json:"DimensionList,omitempty" xml:"DimensionList,omitempty" type:"Repeated"`
	// The unique ID of the metric.
	Directory *QueryDatasetInfoResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the data source.
	//
	// example:
	//
	// a201c85c-******
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The time when the dataset was last modified.
	//
	// example:
	//
	// odps
	DsName *string `json:"DsName,omitempty" xml:"DsName,omitempty"`
	// The point in time when the training dataset was created.
	//
	// example:
	//
	// odps
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// Indicates whether to customize SQL statements. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 1629450382000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The information about the dataset.
	//
	// example:
	//
	// 1629450382000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// A list of all measures for the dataset.
	MeasureList []*QueryDatasetInfoResponseBodyResultMeasureList `json:"MeasureList,omitempty" xml:"MeasureList,omitempty" type:"Repeated"`
	// Whether to enable extraction acceleration. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	OpenOfflineAcceleration *bool `json:"OpenOfflineAcceleration,omitempty" xml:"OpenOfflineAcceleration,omitempty"`
	// Test Space
	//
	// example:
	//
	// b8494aab26124*****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unique ID of the data source.
	//
	// example:
	//
	// The name of the dataset owner.
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The name of the training dataset.
	//
	// example:
	//
	// false
	RowLevel *bool `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	// Whether row-level permissions are enabled. Valid values:
	//
	// 	- true: The VIP Netty channel is enabled.
	//
	// 	- false: The VIP Netty channel is disabled.
	//
	// example:
	//
	// 420abef4-a79b-4289-b12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Big Baby
	//
	// example:
	//
	// The name of the workspace in which the dataset resides.
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResult) GetCubeTableList() []*QueryDatasetInfoResponseBodyResultCubeTableList {
	return s.CubeTableList
}

func (s *QueryDatasetInfoResponseBodyResult) GetCustimzeSql() *bool {
	return s.CustimzeSql
}

func (s *QueryDatasetInfoResponseBodyResult) GetDatasetId() *string {
	return s.DatasetId
}

func (s *QueryDatasetInfoResponseBodyResult) GetDatasetName() *string {
	return s.DatasetName
}

func (s *QueryDatasetInfoResponseBodyResult) GetDimensionList() []*QueryDatasetInfoResponseBodyResultDimensionList {
	return s.DimensionList
}

func (s *QueryDatasetInfoResponseBodyResult) GetDirectory() *QueryDatasetInfoResponseBodyResultDirectory {
	return s.Directory
}

func (s *QueryDatasetInfoResponseBodyResult) GetDsId() *string {
	return s.DsId
}

func (s *QueryDatasetInfoResponseBodyResult) GetDsName() *string {
	return s.DsName
}

func (s *QueryDatasetInfoResponseBodyResult) GetDsType() *string {
	return s.DsType
}

func (s *QueryDatasetInfoResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryDatasetInfoResponseBodyResult) GetGmtModify() *string {
	return s.GmtModify
}

func (s *QueryDatasetInfoResponseBodyResult) GetMeasureList() []*QueryDatasetInfoResponseBodyResultMeasureList {
	return s.MeasureList
}

func (s *QueryDatasetInfoResponseBodyResult) GetOpenOfflineAcceleration() *bool {
	return s.OpenOfflineAcceleration
}

func (s *QueryDatasetInfoResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryDatasetInfoResponseBodyResult) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryDatasetInfoResponseBodyResult) GetRowLevel() *bool {
	return s.RowLevel
}

func (s *QueryDatasetInfoResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryDatasetInfoResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *QueryDatasetInfoResponseBodyResult) SetCubeTableList(v []*QueryDatasetInfoResponseBodyResultCubeTableList) *QueryDatasetInfoResponseBodyResult {
	s.CubeTableList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetCustimzeSql(v bool) *QueryDatasetInfoResponseBodyResult {
	s.CustimzeSql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDatasetId(v string) *QueryDatasetInfoResponseBodyResult {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDatasetName(v string) *QueryDatasetInfoResponseBodyResult {
	s.DatasetName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDimensionList(v []*QueryDatasetInfoResponseBodyResultDimensionList) *QueryDatasetInfoResponseBodyResult {
	s.DimensionList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDirectory(v *QueryDatasetInfoResponseBodyResultDirectory) *QueryDatasetInfoResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsId(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsName(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsType(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetGmtCreate(v string) *QueryDatasetInfoResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetGmtModify(v string) *QueryDatasetInfoResponseBodyResult {
	s.GmtModify = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetMeasureList(v []*QueryDatasetInfoResponseBodyResultMeasureList) *QueryDatasetInfoResponseBodyResult {
	s.MeasureList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetOpenOfflineAcceleration(v bool) *QueryDatasetInfoResponseBodyResult {
	s.OpenOfflineAcceleration = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetOwnerId(v string) *QueryDatasetInfoResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetOwnerName(v string) *QueryDatasetInfoResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetRowLevel(v bool) *QueryDatasetInfoResponseBodyResult {
	s.RowLevel = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetWorkspaceId(v string) *QueryDatasetInfoResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetWorkspaceName(v string) *QueryDatasetInfoResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) Validate() error {
	if s.CubeTableList != nil {
		for _, item := range s.CubeTableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DimensionList != nil {
		for _, item := range s.DimensionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	if s.MeasureList != nil {
		for _, item := range s.MeasureList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDatasetInfoResponseBodyResultCubeTableList struct {
	// Indicates whether the data source table is valid. Valid values:
	//
	// 	- true: data source table
	//
	// 	- false: custom table
	//
	// example:
	//
	// odps_40
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The display name of the table.
	//
	// example:
	//
	// false
	Customsql *bool `json:"Customsql,omitempty" xml:"Customsql,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// dfefd7f4-fc6e-42c9-b4******
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// maxcompute
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The unique ID of the table.
	//
	// example:
	//
	// true
	FactTable *bool `json:"FactTable,omitempty" xml:"FactTable,omitempty"`
	// Indicates whether the table is a custom SQL table. Valid values:
	//
	// 	- true: custom SQL table
	//
	// 	- false: non-custom SQL table
	//
	// example:
	//
	// select 	- from ****
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The list of tables used by the dataset.
	//
	// example:
	//
	// viewdasb8494aab2612473cb74992159fe****
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- mysql
	//
	// 	- odps
	//
	// 	- oracle
	//
	// 	- ... and other data source types supported by Quick BI
	//
	// example:
	//
	// 7a62530b36
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultCubeTableList) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultCubeTableList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetCaption() *string {
	return s.Caption
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetCustomsql() *bool {
	return s.Customsql
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetDsType() *string {
	return s.DsType
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetFactTable() *bool {
	return s.FactTable
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetSql() *string {
	return s.Sql
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetTableName() *string {
	return s.TableName
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetCustomsql(v bool) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Customsql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetDatasourceId(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.DatasourceId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetDsType(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.DsType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetFactTable(v bool) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.FactTable = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetSql(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Sql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetTableName(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.TableName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetUniqueId(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.UniqueId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetInfoResponseBodyResultDimensionList struct {
	// The unique ID of the field that is referenced by the group measure. Non-NULL if and only if the metric is a grouping metric.
	//
	// example:
	//
	// city
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// A list of all dimensions in the dataset.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The actual physical field.
	//
	// example:
	//
	// group_dimension
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// Data type; value:
	//
	// 	- string: character
	//
	// 	- number: a number
	//
	// 	- datetime: time
	//
	// example:
	//
	// example_expression
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Expression for the flattened calculation dimensions.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example_expression
	ExpressionV2 *string `json:"ExpressionV2,omitempty" xml:"ExpressionV2,omitempty"`
	// Expression for a calculated dimension; valid only for calculated dimensions.
	//
	// example:
	//
	// city
	FactColumn *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	// The description of the field.
	//
	// example:
	//
	// hhhh
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// The type of the dimension. Valid values:
	//
	// 	- standard_dimension: General Dimension
	//
	// 	- calculate_dimension: calculating dimensions
	//
	// 	- group_dimension: grouping dimensions
	//
	// example:
	//
	// example_granularity
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The granularity.
	//
	// example:
	//
	// 308f7****
	RefUid *string `json:"RefUid,omitempty" xml:"RefUid,omitempty"`
	// The ARN.
	//
	// example:
	//
	// 7a62530***
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	// The display name of the dimension.
	//
	// example:
	//
	// a69774***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultDimensionList) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultDimensionList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetCaption() *string {
	return s.Caption
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetDataType() *string {
	return s.DataType
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetDimensionType() *string {
	return s.DimensionType
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetExpression() *string {
	return s.Expression
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetExpressionV2() *string {
	return s.ExpressionV2
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetFactColumn() *string {
	return s.FactColumn
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetRefUid() *string {
	return s.RefUid
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetTableUniqueId() *string {
	return s.TableUniqueId
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) GetUid() *string {
	return s.Uid
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetDataType(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.DataType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetDimensionType(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.DimensionType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetExpression(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Expression = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetExpressionV2(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.ExpressionV2 = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetFactColumn(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.FactColumn = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetFieldDescription(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.FieldDescription = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetGranularity(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Granularity = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetRefUid(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.RefUid = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetTableUniqueId(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.TableUniqueId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetUid(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Uid = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetInfoResponseBodyResultDirectory struct {
	// Test directory
	//
	// example:
	//
	// a3eecab7-618d-4f9f-*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Test directory
	//
	// example:
	//
	// The name of the directory.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the directory to which the dataset belongs.
	//
	// example:
	//
	// 88b680****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The path of the directory ID, for example, aa/bb/cc/dd.
	//
	// example:
	//
	// The path name of the directory ID, for example, one-level directory /two-level directory.
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultDirectory) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) GetId() *string {
	return s.Id
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) GetName() *string {
	return s.Name
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) GetPathId() *string {
	return s.PathId
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) GetPathName() *string {
	return s.PathName
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetId(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetName(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetPathId(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetPathName(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) Validate() error {
	return dara.Validate(s)
}

type QueryDatasetInfoResponseBodyResultMeasureList struct {
	// The actual physical field.
	//
	// example:
	//
	// profit_amt
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// A list of all measures for the dataset.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Data type; value:
	//
	// 	- string: character
	//
	// 	- number: a number
	//
	// 	- datetime: time
	//
	// example:
	//
	// example_expression
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Expression for flattened computation metrics.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example_expression
	ExpressionV2 *string `json:"ExpressionV2,omitempty" xml:"ExpressionV2,omitempty"`
	// The type of the measure. Valid values:
	//
	// 	- standard_measure: General Metrics
	//
	// 	- calculate_measure: Calculating Measures
	//
	// example:
	//
	// profit_amt
	FactColumn *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	// The description of the field.
	//
	// example:
	//
	// asadsda
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// An expression that calculates a measure; valid only for calculated measures.
	//
	// example:
	//
	// calculate_measure
	MeasureType *string `json:"MeasureType,omitempty" xml:"MeasureType,omitempty"`
	// The display name of the metric.
	//
	// example:
	//
	// 7a62530b36
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	// The unique ID of the table to which the table belongs, which corresponds to the UniqueId of the CubeTypeList.
	//
	// example:
	//
	// 88b680****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultMeasureList) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultMeasureList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetCaption() *string {
	return s.Caption
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetDataType() *string {
	return s.DataType
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetExpression() *string {
	return s.Expression
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetExpressionV2() *string {
	return s.ExpressionV2
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetFactColumn() *string {
	return s.FactColumn
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetMeasureType() *string {
	return s.MeasureType
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetTableUniqueId() *string {
	return s.TableUniqueId
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) GetUid() *string {
	return s.Uid
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetDataType(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.DataType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetExpression(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Expression = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetExpressionV2(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.ExpressionV2 = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetFactColumn(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.FactColumn = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetFieldDescription(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.FieldDescription = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetMeasureType(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.MeasureType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetTableUniqueId(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.TableUniqueId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetUid(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Uid = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) Validate() error {
	return dara.Validate(s)
}
