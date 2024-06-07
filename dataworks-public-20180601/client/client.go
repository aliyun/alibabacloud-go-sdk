// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateManualDagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-12 00:00:00
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// example:
	//
	// param_k1=param_v1 param_k2=param_v2
	DagPara *string `json:"DagPara,omitempty" xml:"DagPara,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// {"103180025": "test=$[yyyy-mm-dd]"}
	NodePara *string `json:"NodePara,omitempty" xml:"NodePara,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateManualDagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateManualDagRequest) GoString() string {
	return s.String()
}

func (s *CreateManualDagRequest) SetBizdate(v string) *CreateManualDagRequest {
	s.Bizdate = &v
	return s
}

func (s *CreateManualDagRequest) SetDagPara(v string) *CreateManualDagRequest {
	s.DagPara = &v
	return s
}

func (s *CreateManualDagRequest) SetFlowName(v string) *CreateManualDagRequest {
	s.FlowName = &v
	return s
}

func (s *CreateManualDagRequest) SetNodePara(v string) *CreateManualDagRequest {
	s.NodePara = &v
	return s
}

func (s *CreateManualDagRequest) SetProjectName(v string) *CreateManualDagRequest {
	s.ProjectName = &v
	return s
}

type CreateManualDagResponseBody struct {
	// example:
	//
	// 2d9ce-38ef-4923-baf6-391a7e656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// example:
	//
	// test
	ReturnErrorSolution *string `json:"ReturnErrorSolution,omitempty" xml:"ReturnErrorSolution,omitempty"`
	// example:
	//
	// test
	ReturnMessage *string `json:"ReturnMessage,omitempty" xml:"ReturnMessage,omitempty"`
	// example:
	//
	// 1244311235
	ReturnValue *int64 `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty"`
}

func (s CreateManualDagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateManualDagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManualDagResponseBody) SetRequestId(v string) *CreateManualDagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateManualDagResponseBody) SetReturnCode(v string) *CreateManualDagResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *CreateManualDagResponseBody) SetReturnErrorSolution(v string) *CreateManualDagResponseBody {
	s.ReturnErrorSolution = &v
	return s
}

func (s *CreateManualDagResponseBody) SetReturnMessage(v string) *CreateManualDagResponseBody {
	s.ReturnMessage = &v
	return s
}

func (s *CreateManualDagResponseBody) SetReturnValue(v int64) *CreateManualDagResponseBody {
	s.ReturnValue = &v
	return s
}

type CreateManualDagResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateManualDagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateManualDagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateManualDagResponse) GoString() string {
	return s.String()
}

func (s *CreateManualDagResponse) SetHeaders(v map[string]*string) *CreateManualDagResponse {
	s.Headers = v
	return s
}

func (s *CreateManualDagResponse) SetStatusCode(v int32) *CreateManualDagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateManualDagResponse) SetBody(v *CreateManualDagResponseBody) *CreateManualDagResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	// This parameter is required.
	FileId            *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	ProjectId         *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetFileId(v int64) *DeleteFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteFileRequest) SetProjectId(v int64) *DeleteFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFileRequest) SetProjectIdentifier(v string) *DeleteFileRequest {
	s.ProjectIdentifier = &v
	return s
}

type DeleteFileResponseBody struct {
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetHttpStatusCode(v int32) *DeleteFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DescribeEmrHiveTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// pt_table_090901_emr
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeEmrHiveTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmrHiveTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmrHiveTableRequest) SetClusterId(v string) *DescribeEmrHiveTableRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeEmrHiveTableRequest) SetDatabaseName(v string) *DescribeEmrHiveTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeEmrHiveTableRequest) SetTableName(v string) *DescribeEmrHiveTableRequest {
	s.TableName = &v
	return s
}

type DescribeEmrHiveTableResponseBody struct {
	Data *DescribeEmrHiveTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// test
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E6F0DBDD-5AD8-4870-A6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmrHiveTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmrHiveTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmrHiveTableResponseBody) SetData(v *DescribeEmrHiveTableResponseBodyData) *DescribeEmrHiveTableResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEmrHiveTableResponseBody) SetErrorCode(v string) *DescribeEmrHiveTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBody) SetErrorMessage(v string) *DescribeEmrHiveTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBody) SetRequestId(v string) *DescribeEmrHiveTableResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEmrHiveTableResponseBodyData struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterBizId   *string                                        `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ClusterBizName *string                                        `json:"ClusterBizName,omitempty" xml:"ClusterBizName,omitempty"`
	Columns        []*DescribeEmrHiveTableResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 2019-09-09 20:41:28
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2019-09-09 20:41:28
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// org.apache.hadoop.mapred.TextInputFormat
	InputFormat *string `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	// example:
	//
	// false
	IsCompressed *bool `json:"IsCompressed,omitempty" xml:"IsCompressed,omitempty"`
	// example:
	//
	// false
	IsTemporary *bool `json:"IsTemporary,omitempty" xml:"IsTemporary,omitempty"`
	// example:
	//
	// 1970-01-01 08:00:00
	LastAccessTime *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// example:
	//
	// hdfs://emr-header-1.cluster-136574:9000/user/hive/warehouse/pt_table_090901_emr
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// example:
	//
	// root
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1861276710322536
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// USER
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// example:
	//
	// DS,HR,REGION
	PartitionKeys *string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty"`
	// example:
	//
	// org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe
	SerializationLib *string `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
	// example:
	//
	// RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA
	TableComment *string `json:"TableComment,omitempty" xml:"TableComment,omitempty"`
	// example:
	//
	// RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA
	TableDesc *string `json:"TableDesc,omitempty" xml:"TableDesc,omitempty"`
	// example:
	//
	// pt_table_090901_emr
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// {\\"transient_lastDdlTime\\":\\"1568031823\\",\\"bucketing_version\\":\\"2\\",\\"comment\\":\\"RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA\\"}
	TableParameters *string `json:"TableParameters,omitempty" xml:"TableParameters,omitempty"`
	// example:
	//
	// 552
	TableSize *int64 `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeEmrHiveTableResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmrHiveTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEmrHiveTableResponseBodyData) SetClusterBizId(v string) *DescribeEmrHiveTableResponseBodyData {
	s.ClusterBizId = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetClusterBizName(v string) *DescribeEmrHiveTableResponseBodyData {
	s.ClusterBizName = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetColumns(v []*DescribeEmrHiveTableResponseBodyDataColumns) *DescribeEmrHiveTableResponseBodyData {
	s.Columns = v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetDatabaseName(v string) *DescribeEmrHiveTableResponseBodyData {
	s.DatabaseName = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetGmtCreate(v string) *DescribeEmrHiveTableResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetGmtModified(v string) *DescribeEmrHiveTableResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetInputFormat(v string) *DescribeEmrHiveTableResponseBodyData {
	s.InputFormat = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetIsCompressed(v bool) *DescribeEmrHiveTableResponseBodyData {
	s.IsCompressed = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetIsTemporary(v bool) *DescribeEmrHiveTableResponseBodyData {
	s.IsTemporary = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetLastAccessTime(v string) *DescribeEmrHiveTableResponseBodyData {
	s.LastAccessTime = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetLastModifyTime(v string) *DescribeEmrHiveTableResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetLocation(v string) *DescribeEmrHiveTableResponseBodyData {
	s.Location = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetOutputFormat(v string) *DescribeEmrHiveTableResponseBodyData {
	s.OutputFormat = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetOwner(v string) *DescribeEmrHiveTableResponseBodyData {
	s.Owner = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetOwnerId(v string) *DescribeEmrHiveTableResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetOwnerType(v string) *DescribeEmrHiveTableResponseBodyData {
	s.OwnerType = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetPartitionKeys(v string) *DescribeEmrHiveTableResponseBodyData {
	s.PartitionKeys = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetSerializationLib(v string) *DescribeEmrHiveTableResponseBodyData {
	s.SerializationLib = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableComment(v string) *DescribeEmrHiveTableResponseBodyData {
	s.TableComment = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableDesc(v string) *DescribeEmrHiveTableResponseBodyData {
	s.TableDesc = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableName(v string) *DescribeEmrHiveTableResponseBodyData {
	s.TableName = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableParameters(v string) *DescribeEmrHiveTableResponseBodyData {
	s.TableParameters = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableSize(v int64) *DescribeEmrHiveTableResponseBodyData {
	s.TableSize = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyData) SetTableType(v string) *DescribeEmrHiveTableResponseBodyData {
	s.TableType = &v
	return s
}

type DescribeEmrHiveTableResponseBodyDataColumns struct {
	// example:
	//
	// BALANCE FIELD
	ColumnComment *string `json:"ColumnComment,omitempty" xml:"ColumnComment,omitempty"`
	// example:
	//
	// double
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// 1
	ColumnPosition *int32 `json:"ColumnPosition,omitempty" xml:"ColumnPosition,omitempty"`
	// example:
	//
	// double
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// example:
	//
	// BALANCE FIELD
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s DescribeEmrHiveTableResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmrHiveTableResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetColumnComment(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.ColumnComment = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetColumnName(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.ColumnName = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetColumnPosition(v int32) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.ColumnPosition = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetColumnType(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.ColumnType = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetComment(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.Comment = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetGmtCreate(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEmrHiveTableResponseBodyDataColumns) SetGmtModified(v string) *DescribeEmrHiveTableResponseBodyDataColumns {
	s.GmtModified = &v
	return s
}

type DescribeEmrHiveTableResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmrHiveTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEmrHiveTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmrHiveTableResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmrHiveTableResponse) SetHeaders(v map[string]*string) *DescribeEmrHiveTableResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmrHiveTableResponse) SetStatusCode(v int32) *DescribeEmrHiveTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmrHiveTableResponse) SetBody(v *DescribeEmrHiveTableResponseBody) *DescribeEmrHiveTableResponse {
	s.Body = v
	return s
}

type GetDataServiceApiContextRequest struct {
	// apiId
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 1
	ApiStatus *int32 `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// example:
	//
	// abc-124
	CacheKey *string `json:"CacheKey,omitempty" xml:"CacheKey,omitempty"`
	// example:
	//
	// true
	ForPrivateResGroup *bool `json:"ForPrivateResGroup,omitempty" xml:"ForPrivateResGroup,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetDataServiceApiContextRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataServiceApiContextRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiContextRequest) SetApiId(v int64) *GetDataServiceApiContextRequest {
	s.ApiId = &v
	return s
}

func (s *GetDataServiceApiContextRequest) SetApiStatus(v int32) *GetDataServiceApiContextRequest {
	s.ApiStatus = &v
	return s
}

func (s *GetDataServiceApiContextRequest) SetCacheKey(v string) *GetDataServiceApiContextRequest {
	s.CacheKey = &v
	return s
}

func (s *GetDataServiceApiContextRequest) SetForPrivateResGroup(v bool) *GetDataServiceApiContextRequest {
	s.ForPrivateResGroup = &v
	return s
}

func (s *GetDataServiceApiContextRequest) SetVerbose(v bool) *GetDataServiceApiContextRequest {
	s.Verbose = &v
	return s
}

type GetDataServiceApiContextResponseBody struct {
	// example:
	//
	// {"apiId":123}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// success
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataServiceApiContextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataServiceApiContextResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiContextResponseBody) SetData(v string) *GetDataServiceApiContextResponseBody {
	s.Data = &v
	return s
}

func (s *GetDataServiceApiContextResponseBody) SetErrCode(v string) *GetDataServiceApiContextResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataServiceApiContextResponseBody) SetErrMsg(v string) *GetDataServiceApiContextResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetDataServiceApiContextResponseBody) SetRequestId(v string) *GetDataServiceApiContextResponseBody {
	s.RequestId = &v
	return s
}

type GetDataServiceApiContextResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiContextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataServiceApiContextResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiContextResponse) SetHeaders(v map[string]*string) *GetDataServiceApiContextResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiContextResponse) SetStatusCode(v int32) *GetDataServiceApiContextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiContextResponse) SetBody(v *GetDataServiceApiContextResponseBody) *GetDataServiceApiContextResponse {
	s.Body = v
	return s
}

type GetDataServiceContextUpdateEventResponseBody struct {
	// example:
	//
	// {\\"TotalCount\\": 0, \\"CalcEngines\\": []}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// success
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// 8754EE08-4AA2-5F77-ADD7-754DBBDA9F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataServiceContextUpdateEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataServiceContextUpdateEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceContextUpdateEventResponseBody) SetData(v string) *GetDataServiceContextUpdateEventResponseBody {
	s.Data = &v
	return s
}

func (s *GetDataServiceContextUpdateEventResponseBody) SetErrCode(v string) *GetDataServiceContextUpdateEventResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataServiceContextUpdateEventResponseBody) SetErrMsg(v string) *GetDataServiceContextUpdateEventResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetDataServiceContextUpdateEventResponseBody) SetRequestId(v string) *GetDataServiceContextUpdateEventResponseBody {
	s.RequestId = &v
	return s
}

type GetDataServiceContextUpdateEventResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceContextUpdateEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceContextUpdateEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataServiceContextUpdateEventResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceContextUpdateEventResponse) SetHeaders(v map[string]*string) *GetDataServiceContextUpdateEventResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceContextUpdateEventResponse) SetStatusCode(v int32) *GetDataServiceContextUpdateEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceContextUpdateEventResponse) SetBody(v *GetDataServiceContextUpdateEventResponseBody) *GetDataServiceContextUpdateEventResponse {
	s.Body = v
	return s
}

type GetSwitchValueRequest struct {
	SwitchName *string `json:"SwitchName,omitempty" xml:"SwitchName,omitempty"`
}

func (s GetSwitchValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchValueRequest) GoString() string {
	return s.String()
}

func (s *GetSwitchValueRequest) SetSwitchName(v string) *GetSwitchValueRequest {
	s.SwitchName = &v
	return s
}

type GetSwitchValueResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSwitchValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwitchValueResponseBody) SetData(v string) *GetSwitchValueResponseBody {
	s.Data = &v
	return s
}

func (s *GetSwitchValueResponseBody) SetRequestId(v string) *GetSwitchValueResponseBody {
	s.RequestId = &v
	return s
}

type GetSwitchValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSwitchValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSwitchValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchValueResponse) GoString() string {
	return s.String()
}

func (s *GetSwitchValueResponse) SetHeaders(v map[string]*string) *GetSwitchValueResponse {
	s.Headers = v
	return s
}

func (s *GetSwitchValueResponse) SetStatusCode(v int32) *GetSwitchValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwitchValueResponse) SetBody(v *GetSwitchValueResponseBody) *GetSwitchValueResponse {
	s.Body = v
	return s
}

type GetTimeMachineTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTimeMachineTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTimeMachineTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTimeMachineTaskRequest) SetTaskId(v string) *GetTimeMachineTaskRequest {
	s.TaskId = &v
	return s
}

type GetTimeMachineTaskResponseBody struct {
	Data      *GetTimeMachineTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMsg    *string                             `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTimeMachineTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTimeMachineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTimeMachineTaskResponseBody) SetData(v *GetTimeMachineTaskResponseBodyData) *GetTimeMachineTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTimeMachineTaskResponseBody) SetErrCode(v string) *GetTimeMachineTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTimeMachineTaskResponseBody) SetErrMsg(v string) *GetTimeMachineTaskResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetTimeMachineTaskResponseBody) SetRequestId(v string) *GetTimeMachineTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetTimeMachineTaskResponseBodyData struct {
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ObjId       *string `json:"ObjId,omitempty" xml:"ObjId,omitempty"`
	ObjName     *string `json:"ObjName,omitempty" xml:"ObjName,omitempty"`
	OperType    *string `json:"OperType,omitempty" xml:"OperType,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTimeMachineTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTimeMachineTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTimeMachineTaskResponseBodyData) SetGmtCreate(v string) *GetTimeMachineTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetGmtModified(v string) *GetTimeMachineTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetHostName(v string) *GetTimeMachineTaskResponseBodyData {
	s.HostName = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetId(v string) *GetTimeMachineTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetObjId(v string) *GetTimeMachineTaskResponseBodyData {
	s.ObjId = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetObjName(v string) *GetTimeMachineTaskResponseBodyData {
	s.ObjName = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetOperType(v string) *GetTimeMachineTaskResponseBodyData {
	s.OperType = &v
	return s
}

func (s *GetTimeMachineTaskResponseBodyData) SetStatus(v string) *GetTimeMachineTaskResponseBodyData {
	s.Status = &v
	return s
}

type GetTimeMachineTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTimeMachineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTimeMachineTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTimeMachineTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTimeMachineTaskResponse) SetHeaders(v map[string]*string) *GetTimeMachineTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTimeMachineTaskResponse) SetStatusCode(v int32) *GetTimeMachineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTimeMachineTaskResponse) SetBody(v *GetTimeMachineTaskResponseBody) *GetTimeMachineTaskResponse {
	s.Body = v
	return s
}

type ListEmrHiveAuditLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 1586509710
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1586509407
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListEmrHiveAuditLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveAuditLogsRequest) GoString() string {
	return s.String()
}

func (s *ListEmrHiveAuditLogsRequest) SetClusterId(v string) *ListEmrHiveAuditLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetDatabaseName(v string) *ListEmrHiveAuditLogsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetEndTime(v int32) *ListEmrHiveAuditLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetPageNumber(v int32) *ListEmrHiveAuditLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetPageSize(v int32) *ListEmrHiveAuditLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetStartTime(v int32) *ListEmrHiveAuditLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListEmrHiveAuditLogsRequest) SetTableName(v string) *ListEmrHiveAuditLogsRequest {
	s.TableName = &v
	return s
}

type ListEmrHiveAuditLogsResponseBody struct {
	Data *ListEmrHiveAuditLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 58D5334A-B013-430E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEmrHiveAuditLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveAuditLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmrHiveAuditLogsResponseBody) SetData(v *ListEmrHiveAuditLogsResponseBodyData) *ListEmrHiveAuditLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBody) SetErrorCode(v string) *ListEmrHiveAuditLogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBody) SetErrorMessage(v string) *ListEmrHiveAuditLogsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBody) SetRequestId(v string) *ListEmrHiveAuditLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListEmrHiveAuditLogsResponseBodyData struct {
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 2
	PageSize  *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PagedData []*ListEmrHiveAuditLogsResponseBodyDataPagedData `json:"PagedData,omitempty" xml:"PagedData,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEmrHiveAuditLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveAuditLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEmrHiveAuditLogsResponseBodyData) SetPageNumber(v int32) *ListEmrHiveAuditLogsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyData) SetPageSize(v int32) *ListEmrHiveAuditLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyData) SetPagedData(v []*ListEmrHiveAuditLogsResponseBodyDataPagedData) *ListEmrHiveAuditLogsResponseBodyData {
	s.PagedData = v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyData) SetTotalCount(v int32) *ListEmrHiveAuditLogsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListEmrHiveAuditLogsResponseBodyDataPagedData struct {
	// example:
	//
	// default
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// example:
	//
	// 1564019679506
	EventTime *int64    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Groups    []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// example:
	//
	// CREATE_TABLE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// test_table
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListEmrHiveAuditLogsResponseBodyDataPagedData) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveAuditLogsResponseBodyDataPagedData) GoString() string {
	return s.String()
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetDatabase(v string) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.Database = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetEventTime(v int64) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.EventTime = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetGroups(v []*string) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.Groups = v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetOperation(v string) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.Operation = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetTable(v string) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.Table = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponseBodyDataPagedData) SetUser(v string) *ListEmrHiveAuditLogsResponseBodyDataPagedData {
	s.User = &v
	return s
}

type ListEmrHiveAuditLogsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEmrHiveAuditLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEmrHiveAuditLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveAuditLogsResponse) GoString() string {
	return s.String()
}

func (s *ListEmrHiveAuditLogsResponse) SetHeaders(v map[string]*string) *ListEmrHiveAuditLogsResponse {
	s.Headers = v
	return s
}

func (s *ListEmrHiveAuditLogsResponse) SetStatusCode(v int32) *ListEmrHiveAuditLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmrHiveAuditLogsResponse) SetBody(v *ListEmrHiveAuditLogsResponseBody) *ListEmrHiveAuditLogsResponse {
	s.Body = v
	return s
}

type ListEmrHiveDatabasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListEmrHiveDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListEmrHiveDatabasesRequest) SetClusterId(v string) *ListEmrHiveDatabasesRequest {
	s.ClusterId = &v
	return s
}

type ListEmrHiveDatabasesResponseBody struct {
	Data []*ListEmrHiveDatabasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// test
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// D9A61DC0-B922-421B-B706
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEmrHiveDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmrHiveDatabasesResponseBody) SetData(v []*ListEmrHiveDatabasesResponseBodyData) *ListEmrHiveDatabasesResponseBody {
	s.Data = v
	return s
}

func (s *ListEmrHiveDatabasesResponseBody) SetErrorCode(v string) *ListEmrHiveDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBody) SetErrorMessage(v string) *ListEmrHiveDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBody) SetRequestId(v string) *ListEmrHiveDatabasesResponseBody {
	s.RequestId = &v
	return s
}

type ListEmrHiveDatabasesResponseBodyData struct {
	// example:
	//
	// Default Hive database
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1568010630000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1568010630000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// hdfs://emr-header-1.cluster-136574:9000/user/hive/warehouse
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// public
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 18612767
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ROLE
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// -
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// HIVE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEmrHiveDatabasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveDatabasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetComment(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Comment = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetGmtCreate(v int64) *ListEmrHiveDatabasesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetGmtModified(v int64) *ListEmrHiveDatabasesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetLocation(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Location = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetName(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetOwner(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetOwnerId(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetOwnerType(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.OwnerType = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetParameters(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Parameters = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetRegion(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetStatus(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEmrHiveDatabasesResponseBodyData) SetType(v string) *ListEmrHiveDatabasesResponseBodyData {
	s.Type = &v
	return s
}

type ListEmrHiveDatabasesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEmrHiveDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEmrHiveDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListEmrHiveDatabasesResponse) SetHeaders(v map[string]*string) *ListEmrHiveDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListEmrHiveDatabasesResponse) SetStatusCode(v int32) *ListEmrHiveDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmrHiveDatabasesResponse) SetBody(v *ListEmrHiveDatabasesResponseBody) *ListEmrHiveDatabasesResponse {
	s.Body = v
	return s
}

type ListEmrHiveTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEmrHiveTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveTablesRequest) GoString() string {
	return s.String()
}

func (s *ListEmrHiveTablesRequest) SetClusterId(v string) *ListEmrHiveTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEmrHiveTablesRequest) SetDatabaseName(v string) *ListEmrHiveTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListEmrHiveTablesRequest) SetPageNumber(v int32) *ListEmrHiveTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEmrHiveTablesRequest) SetPageSize(v int32) *ListEmrHiveTablesRequest {
	s.PageSize = &v
	return s
}

type ListEmrHiveTablesResponseBody struct {
	Data *ListEmrHiveTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// test
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E6F0DBDD-5AD8-4870-A6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEmrHiveTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmrHiveTablesResponseBody) SetData(v *ListEmrHiveTablesResponseBodyData) *ListEmrHiveTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListEmrHiveTablesResponseBody) SetErrorCode(v string) *ListEmrHiveTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEmrHiveTablesResponseBody) SetErrorMessage(v string) *ListEmrHiveTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEmrHiveTablesResponseBody) SetRequestId(v string) *ListEmrHiveTablesResponseBody {
	s.RequestId = &v
	return s
}

type ListEmrHiveTablesResponseBodyData struct {
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 2
	PageSize  *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PagedData []*ListEmrHiveTablesResponseBodyDataPagedData `json:"PagedData,omitempty" xml:"PagedData,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEmrHiveTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEmrHiveTablesResponseBodyData) SetPageNumber(v int32) *ListEmrHiveTablesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyData) SetPageSize(v int32) *ListEmrHiveTablesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyData) SetPagedData(v []*ListEmrHiveTablesResponseBodyDataPagedData) *ListEmrHiveTablesResponseBodyData {
	s.PagedData = v
	return s
}

func (s *ListEmrHiveTablesResponseBodyData) SetTotalCount(v int32) *ListEmrHiveTablesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListEmrHiveTablesResponseBodyDataPagedData struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterBizId   *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ClusterBizName *string `json:"ClusterBizName,omitempty" xml:"ClusterBizName,omitempty"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// org.apache.hadoop.hive.ql.io.orc.OrcInputFormat
	InputFormat *string `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	// example:
	//
	// false
	IsCompressed *bool `json:"IsCompressed,omitempty" xml:"IsCompressed,omitempty"`
	// example:
	//
	// false
	IsTemporary *bool `json:"IsTemporary,omitempty" xml:"IsTemporary,omitempty"`
	// example:
	//
	// 1970-01-01 08:00:00
	LastAccessTime *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// example:
	//
	// 2019-09-09 20:23:47
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// example:
	//
	// hdfs://emr-header-1.cluster-136574:9000/user/hive/warehouse/pt_table_090901_emr
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// example:
	//
	// root
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 18612767103****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ROLE
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// example:
	//
	// DS,HR,REGION
	PartitionKeys *string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty"`
	// example:
	//
	// org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe
	SerializationLib *string `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
	// example:
	//
	// RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA
	TableComment *string `json:"TableComment,omitempty" xml:"TableComment,omitempty"`
	// example:
	//
	// RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA
	TableDesc *string `json:"TableDesc,omitempty" xml:"TableDesc,omitempty"`
	// example:
	//
	// pt_table_090901_emr_orc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// {\\"transient_lastDdlTime\\":\\"1568031823\\",\\"bucketing_version\\":\\"2\\",\\"comment\\":\\"RANDOM GENERATED TEST DATA BY FUNCTION OF RANDOM_TEST_DATA\\"}
	TableParameters *string `json:"TableParameters,omitempty" xml:"TableParameters,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListEmrHiveTablesResponseBodyDataPagedData) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveTablesResponseBodyDataPagedData) GoString() string {
	return s.String()
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetClusterBizId(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.ClusterBizId = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetClusterBizName(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.ClusterBizName = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetDatabaseName(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.DatabaseName = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetGmtCreate(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.GmtCreate = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetGmtModified(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.GmtModified = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetInputFormat(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.InputFormat = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetIsCompressed(v bool) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.IsCompressed = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetIsTemporary(v bool) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.IsTemporary = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetLastAccessTime(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.LastAccessTime = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetLastModifyTime(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.LastModifyTime = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetLocation(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.Location = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetOutputFormat(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.OutputFormat = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetOwner(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.Owner = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetOwnerId(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.OwnerId = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetOwnerType(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.OwnerType = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetPartitionKeys(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.PartitionKeys = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetSerializationLib(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.SerializationLib = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetTableComment(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.TableComment = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetTableDesc(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.TableDesc = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetTableName(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.TableName = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetTableParameters(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.TableParameters = &v
	return s
}

func (s *ListEmrHiveTablesResponseBodyDataPagedData) SetTableType(v string) *ListEmrHiveTablesResponseBodyDataPagedData {
	s.TableType = &v
	return s
}

type ListEmrHiveTablesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEmrHiveTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEmrHiveTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmrHiveTablesResponse) GoString() string {
	return s.String()
}

func (s *ListEmrHiveTablesResponse) SetHeaders(v map[string]*string) *ListEmrHiveTablesResponse {
	s.Headers = v
	return s
}

func (s *ListEmrHiveTablesResponse) SetStatusCode(v int32) *ListEmrHiveTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmrHiveTablesResponse) SetBody(v *ListEmrHiveTablesResponseBody) *ListEmrHiveTablesResponse {
	s.Body = v
	return s
}

type ListGovernanceIssueDataServiceAPIsRequest struct {
	BizDate    *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListGovernanceIssueDataServiceAPIsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueDataServiceAPIsRequest) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetBizDate(v string) *ListGovernanceIssueDataServiceAPIsRequest {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetOwnerId(v string) *ListGovernanceIssueDataServiceAPIsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetPageNumber(v int32) *ListGovernanceIssueDataServiceAPIsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetPageSize(v int32) *ListGovernanceIssueDataServiceAPIsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetProjectId(v int64) *ListGovernanceIssueDataServiceAPIsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetRuleCategory(v string) *ListGovernanceIssueDataServiceAPIsRequest {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsRequest) SetRuleId(v string) *ListGovernanceIssueDataServiceAPIsRequest {
	s.RuleId = &v
	return s
}

type ListGovernanceIssueDataServiceAPIsResponseBody struct {
	Data                *ListGovernanceIssueDataServiceAPIsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicErrorCode    *string                                             `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	DynamicErrorMessage *string                                             `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
	ErrorCode           *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode      *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGovernanceIssueDataServiceAPIsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueDataServiceAPIsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetData(v *ListGovernanceIssueDataServiceAPIsResponseBodyData) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.Data = v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetDynamicErrorCode(v string) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetDynamicErrorMessage(v string) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetErrorCode(v string) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetErrorMessage(v string) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetHttpStatusCode(v int32) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetRequestId(v string) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBody) SetSuccess(v bool) *ListGovernanceIssueDataServiceAPIsResponseBody {
	s.Success = &v
	return s
}

type ListGovernanceIssueDataServiceAPIsResponseBodyData struct {
	APIs       []*ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs `json:"APIs,omitempty" xml:"APIs,omitempty" type:"Repeated"`
	PageNumber *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGovernanceIssueDataServiceAPIsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueDataServiceAPIsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyData) SetAPIs(v []*ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) *ListGovernanceIssueDataServiceAPIsResponseBodyData {
	s.APIs = v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyData) SetPageNumber(v int32) *ListGovernanceIssueDataServiceAPIsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyData) SetPageSize(v int32) *ListGovernanceIssueDataServiceAPIsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyData) SetTotalCount(v int32) *ListGovernanceIssueDataServiceAPIsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BizDate      *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Properties   *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetApiId(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.ApiId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetApiName(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.ApiName = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetBizDate(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetOwnerId(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetProjectId(v int64) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetProperties(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.Properties = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetRuleCategory(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs) SetRuleId(v string) *ListGovernanceIssueDataServiceAPIsResponseBodyDataAPIs {
	s.RuleId = &v
	return s
}

type ListGovernanceIssueDataServiceAPIsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGovernanceIssueDataServiceAPIsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGovernanceIssueDataServiceAPIsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueDataServiceAPIsResponse) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueDataServiceAPIsResponse) SetHeaders(v map[string]*string) *ListGovernanceIssueDataServiceAPIsResponse {
	s.Headers = v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponse) SetStatusCode(v int32) *ListGovernanceIssueDataServiceAPIsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGovernanceIssueDataServiceAPIsResponse) SetBody(v *ListGovernanceIssueDataServiceAPIsResponseBody) *ListGovernanceIssueDataServiceAPIsResponse {
	s.Body = v
	return s
}

type ListGovernanceIssueTablesRequest struct {
	BizDate    *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListGovernanceIssueTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTablesRequest) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTablesRequest) SetBizDate(v string) *ListGovernanceIssueTablesRequest {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetOwnerId(v string) *ListGovernanceIssueTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetPageNumber(v int32) *ListGovernanceIssueTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetPageSize(v int32) *ListGovernanceIssueTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetProjectId(v int64) *ListGovernanceIssueTablesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetRuleCategory(v string) *ListGovernanceIssueTablesRequest {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueTablesRequest) SetRuleId(v string) *ListGovernanceIssueTablesRequest {
	s.RuleId = &v
	return s
}

type ListGovernanceIssueTablesResponseBody struct {
	Data                *ListGovernanceIssueTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicErrorCode    *string                                    `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	DynamicErrorMessage *string                                    `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
	ErrorCode           *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode      *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId           *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGovernanceIssueTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTablesResponseBody) SetData(v *ListGovernanceIssueTablesResponseBodyData) *ListGovernanceIssueTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetDynamicErrorCode(v string) *ListGovernanceIssueTablesResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetDynamicErrorMessage(v string) *ListGovernanceIssueTablesResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetErrorCode(v string) *ListGovernanceIssueTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetErrorMessage(v string) *ListGovernanceIssueTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetHttpStatusCode(v int32) *ListGovernanceIssueTablesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetRequestId(v string) *ListGovernanceIssueTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBody) SetSuccess(v bool) *ListGovernanceIssueTablesResponseBody {
	s.Success = &v
	return s
}

type ListGovernanceIssueTablesResponseBodyData struct {
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tables     []*ListGovernanceIssueTablesResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGovernanceIssueTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTablesResponseBodyData) SetPageNumber(v int32) *ListGovernanceIssueTablesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyData) SetPageSize(v int32) *ListGovernanceIssueTablesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyData) SetTables(v []*ListGovernanceIssueTablesResponseBodyDataTables) *ListGovernanceIssueTablesResponseBodyData {
	s.Tables = v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyData) SetTotalCount(v int32) *ListGovernanceIssueTablesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGovernanceIssueTablesResponseBodyDataTables struct {
	BizDate                   *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	ClusterId                 *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime                *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName              *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DatasourceType            *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	DownstreamDependencyCount *int32  `json:"DownstreamDependencyCount,omitempty" xml:"DownstreamDependencyCount,omitempty"`
	LastAccessTime            *int64  `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LifeCycle                 *int64  `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	McProjectName             *string `json:"McProjectName,omitempty" xml:"McProjectName,omitempty"`
	OwnerId                   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId                 *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Properties                *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RuleCategory              *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId                    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Schema                    *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	TableGuid                 *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TableName                 *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSize                 *int64  `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
}

func (s ListGovernanceIssueTablesResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTablesResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetBizDate(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetClusterId(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.ClusterId = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetCreateTime(v int64) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.CreateTime = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetDatabaseName(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.DatabaseName = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetDatasourceType(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.DatasourceType = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetDownstreamDependencyCount(v int32) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.DownstreamDependencyCount = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetLastAccessTime(v int64) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.LastAccessTime = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetLifeCycle(v int64) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.LifeCycle = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetMcProjectName(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.McProjectName = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetOwnerId(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetProjectId(v int64) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetProperties(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.Properties = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetRuleCategory(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetRuleId(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.RuleId = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetSchema(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.Schema = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetTableGuid(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.TableGuid = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetTableName(v string) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.TableName = &v
	return s
}

func (s *ListGovernanceIssueTablesResponseBodyDataTables) SetTableSize(v int64) *ListGovernanceIssueTablesResponseBodyDataTables {
	s.TableSize = &v
	return s
}

type ListGovernanceIssueTablesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGovernanceIssueTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGovernanceIssueTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTablesResponse) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTablesResponse) SetHeaders(v map[string]*string) *ListGovernanceIssueTablesResponse {
	s.Headers = v
	return s
}

func (s *ListGovernanceIssueTablesResponse) SetStatusCode(v int32) *ListGovernanceIssueTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGovernanceIssueTablesResponse) SetBody(v *ListGovernanceIssueTablesResponseBody) *ListGovernanceIssueTablesResponse {
	s.Body = v
	return s
}

type ListGovernanceIssueTasksRequest struct {
	BizDate    *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListGovernanceIssueTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTasksRequest) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTasksRequest) SetBizDate(v string) *ListGovernanceIssueTasksRequest {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetOwnerId(v string) *ListGovernanceIssueTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetPageNumber(v int32) *ListGovernanceIssueTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetPageSize(v int32) *ListGovernanceIssueTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetProjectId(v int64) *ListGovernanceIssueTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetRuleCategory(v string) *ListGovernanceIssueTasksRequest {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueTasksRequest) SetRuleId(v string) *ListGovernanceIssueTasksRequest {
	s.RuleId = &v
	return s
}

type ListGovernanceIssueTasksResponseBody struct {
	Data                *ListGovernanceIssueTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicErrorCode    *string                                   `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	DynamicErrorMessage *string                                   `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
	ErrorCode           *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode      *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId           *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGovernanceIssueTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTasksResponseBody) SetData(v *ListGovernanceIssueTasksResponseBodyData) *ListGovernanceIssueTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetDynamicErrorCode(v string) *ListGovernanceIssueTasksResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetDynamicErrorMessage(v string) *ListGovernanceIssueTasksResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetErrorCode(v string) *ListGovernanceIssueTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetErrorMessage(v string) *ListGovernanceIssueTasksResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetHttpStatusCode(v int32) *ListGovernanceIssueTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetRequestId(v string) *ListGovernanceIssueTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBody) SetSuccess(v bool) *ListGovernanceIssueTasksResponseBody {
	s.Success = &v
	return s
}

type ListGovernanceIssueTasksResponseBodyData struct {
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tasks      []*ListGovernanceIssueTasksResponseBodyDataTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGovernanceIssueTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTasksResponseBodyData) SetPageNumber(v int32) *ListGovernanceIssueTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyData) SetPageSize(v int32) *ListGovernanceIssueTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyData) SetTasks(v []*ListGovernanceIssueTasksResponseBodyDataTasks) *ListGovernanceIssueTasksResponseBodyData {
	s.Tasks = v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyData) SetTotalCount(v int32) *ListGovernanceIssueTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGovernanceIssueTasksResponseBodyDataTasks struct {
	BizDate      *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	NodeId       *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Properties   *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	RuleId       *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListGovernanceIssueTasksResponseBodyDataTasks) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTasksResponseBodyDataTasks) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetBizDate(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.BizDate = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetNodeId(v int64) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.NodeId = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetNodeName(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.NodeName = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetNodeType(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.NodeType = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetOwnerId(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.OwnerId = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetProjectId(v int64) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.ProjectId = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetProperties(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.Properties = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetRuleCategory(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.RuleCategory = &v
	return s
}

func (s *ListGovernanceIssueTasksResponseBodyDataTasks) SetRuleId(v string) *ListGovernanceIssueTasksResponseBodyDataTasks {
	s.RuleId = &v
	return s
}

type ListGovernanceIssueTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGovernanceIssueTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGovernanceIssueTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceIssueTasksResponse) GoString() string {
	return s.String()
}

func (s *ListGovernanceIssueTasksResponse) SetHeaders(v map[string]*string) *ListGovernanceIssueTasksResponse {
	s.Headers = v
	return s
}

func (s *ListGovernanceIssueTasksResponse) SetStatusCode(v int32) *ListGovernanceIssueTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGovernanceIssueTasksResponse) SetBody(v *ListGovernanceIssueTasksResponseBody) *ListGovernanceIssueTasksResponse {
	s.Body = v
	return s
}

type ListGovernanceRulesRequest struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	IssueType  *string `json:"IssueType,omitempty" xml:"IssueType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGovernanceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGovernanceRulesRequest) SetCategory(v string) *ListGovernanceRulesRequest {
	s.Category = &v
	return s
}

func (s *ListGovernanceRulesRequest) SetIssueType(v string) *ListGovernanceRulesRequest {
	s.IssueType = &v
	return s
}

func (s *ListGovernanceRulesRequest) SetPageNumber(v int32) *ListGovernanceRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceRulesRequest) SetPageSize(v int32) *ListGovernanceRulesRequest {
	s.PageSize = &v
	return s
}

type ListGovernanceRulesResponseBody struct {
	Data                *ListGovernanceRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicErrorCode    *string                              `json:"DynamicErrorCode,omitempty" xml:"DynamicErrorCode,omitempty"`
	DynamicErrorMessage *string                              `json:"DynamicErrorMessage,omitempty" xml:"DynamicErrorMessage,omitempty"`
	ErrorCode           *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode      *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId           *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGovernanceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGovernanceRulesResponseBody) SetData(v *ListGovernanceRulesResponseBodyData) *ListGovernanceRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetDynamicErrorCode(v string) *ListGovernanceRulesResponseBody {
	s.DynamicErrorCode = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetDynamicErrorMessage(v string) *ListGovernanceRulesResponseBody {
	s.DynamicErrorMessage = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetErrorCode(v string) *ListGovernanceRulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetErrorMessage(v string) *ListGovernanceRulesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetHttpStatusCode(v int32) *ListGovernanceRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetRequestId(v string) *ListGovernanceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGovernanceRulesResponseBody) SetSuccess(v bool) *ListGovernanceRulesResponseBody {
	s.Success = &v
	return s
}

type ListGovernanceRulesResponseBodyData struct {
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules      []*ListGovernanceRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGovernanceRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGovernanceRulesResponseBodyData) SetPageNumber(v int32) *ListGovernanceRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyData) SetPageSize(v int32) *ListGovernanceRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyData) SetRules(v []*ListGovernanceRulesResponseBodyDataRules) *ListGovernanceRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *ListGovernanceRulesResponseBodyData) SetTotalCount(v int32) *ListGovernanceRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGovernanceRulesResponseBodyDataRules struct {
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Guide       *string `json:"Guide,omitempty" xml:"Guide,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IssueType   *string `json:"IssueType,omitempty" xml:"IssueType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s ListGovernanceRulesResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetCategory(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Category = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetDescription(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetGuide(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Guide = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetId(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Id = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetIssueType(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.IssueType = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetName(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetNote(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Note = &v
	return s
}

func (s *ListGovernanceRulesResponseBodyDataRules) SetRule(v string) *ListGovernanceRulesResponseBodyDataRules {
	s.Rule = &v
	return s
}

type ListGovernanceRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGovernanceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGovernanceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGovernanceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGovernanceRulesResponse) SetHeaders(v map[string]*string) *ListGovernanceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGovernanceRulesResponse) SetStatusCode(v int32) *ListGovernanceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGovernanceRulesResponse) SetBody(v *ListGovernanceRulesResponseBody) *ListGovernanceRulesResponse {
	s.Body = v
	return s
}

type ListHiveColumnLineagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// balance
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveColumnLineagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesRequest) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesRequest) SetClusterId(v string) *ListHiveColumnLineagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListHiveColumnLineagesRequest) SetColumnName(v string) *ListHiveColumnLineagesRequest {
	s.ColumnName = &v
	return s
}

func (s *ListHiveColumnLineagesRequest) SetDatabaseName(v string) *ListHiveColumnLineagesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveColumnLineagesRequest) SetTableName(v string) *ListHiveColumnLineagesRequest {
	s.TableName = &v
	return s
}

type ListHiveColumnLineagesResponseBody struct {
	Data *ListHiveColumnLineagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 58D5334A-B013-430E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHiveColumnLineagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesResponseBody) SetData(v *ListHiveColumnLineagesResponseBodyData) *ListHiveColumnLineagesResponseBody {
	s.Data = v
	return s
}

func (s *ListHiveColumnLineagesResponseBody) SetErrorCode(v string) *ListHiveColumnLineagesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBody) SetErrorMessage(v string) *ListHiveColumnLineagesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBody) SetRequestId(v string) *ListHiveColumnLineagesResponseBody {
	s.RequestId = &v
	return s
}

type ListHiveColumnLineagesResponseBodyData struct {
	DownstreamLineages []*ListHiveColumnLineagesResponseBodyDataDownstreamLineages `json:"DownstreamLineages,omitempty" xml:"DownstreamLineages,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	DownstreamNumber *int32                                                    `json:"DownstreamNumber,omitempty" xml:"DownstreamNumber,omitempty"`
	UpstreamLineages []*ListHiveColumnLineagesResponseBodyDataUpstreamLineages `json:"UpstreamLineages,omitempty" xml:"UpstreamLineages,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	UpstreamNumber *int32 `json:"UpstreamNumber,omitempty" xml:"UpstreamNumber,omitempty"`
}

func (s ListHiveColumnLineagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesResponseBodyData) SetDownstreamLineages(v []*ListHiveColumnLineagesResponseBodyDataDownstreamLineages) *ListHiveColumnLineagesResponseBodyData {
	s.DownstreamLineages = v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyData) SetDownstreamNumber(v int32) *ListHiveColumnLineagesResponseBodyData {
	s.DownstreamNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyData) SetUpstreamLineages(v []*ListHiveColumnLineagesResponseBodyDataUpstreamLineages) *ListHiveColumnLineagesResponseBodyData {
	s.UpstreamLineages = v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyData) SetUpstreamNumber(v int32) *ListHiveColumnLineagesResponseBodyData {
	s.UpstreamNumber = &v
	return s
}

type ListHiveColumnLineagesResponseBodyDataDownstreamLineages struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// balance
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// 2019-11-10 11:33:52
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 0
	DirectDownColumnNumber *int32 `json:"DirectDownColumnNumber,omitempty" xml:"DirectDownColumnNumber,omitempty"`
	// example:
	//
	// 0
	DirectDownTableNumber *int32 `json:"DirectDownTableNumber,omitempty" xml:"DirectDownTableNumber,omitempty"`
	// example:
	//
	// 1
	DirectUpperColumnNumber *int32 `json:"DirectUpperColumnNumber,omitempty" xml:"DirectUpperColumnNumber,omitempty"`
	// example:
	//
	// 1
	DirectUpperTableNumber *int32 `json:"DirectUpperTableNumber,omitempty" xml:"DirectUpperTableNumber,omitempty"`
	// example:
	//
	// 2019-11-10 11:33:52
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// HIVE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// pt_table_090901_emr_child
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveColumnLineagesResponseBodyDataDownstreamLineages) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesResponseBodyDataDownstreamLineages) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetClusterId(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.ClusterId = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetColumnName(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.ColumnName = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetCreateTime(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.CreateTime = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetDatabaseName(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetDirectDownColumnNumber(v int32) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.DirectDownColumnNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetDirectDownTableNumber(v int32) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.DirectDownTableNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetDirectUpperColumnNumber(v int32) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.DirectUpperColumnNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetDirectUpperTableNumber(v int32) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.DirectUpperTableNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetModifiedTime(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.ModifiedTime = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetSource(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.Source = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataDownstreamLineages) SetTableName(v string) *ListHiveColumnLineagesResponseBodyDataDownstreamLineages {
	s.TableName = &v
	return s
}

type ListHiveColumnLineagesResponseBodyDataUpstreamLineages struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// balance
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// 2019-11-10 11:33:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// 1
	DirectDownColumnNumber *int32 `json:"DirectDownColumnNumber,omitempty" xml:"DirectDownColumnNumber,omitempty"`
	// example:
	//
	// 0
	DirectDownTableNumber *int32 `json:"DirectDownTableNumber,omitempty" xml:"DirectDownTableNumber,omitempty"`
	// example:
	//
	// 2
	DirectUpperColumnNumber *int32 `json:"DirectUpperColumnNumber,omitempty" xml:"DirectUpperColumnNumber,omitempty"`
	// example:
	//
	// 1
	DirectUpperTableNumber *int32 `json:"DirectUpperTableNumber,omitempty" xml:"DirectUpperTableNumber,omitempty"`
	// example:
	//
	// 2019-11-10 11:33:51
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// HIVE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// pt_table_090901_emr_child
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveColumnLineagesResponseBodyDataUpstreamLineages) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesResponseBodyDataUpstreamLineages) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetClusterId(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.ClusterId = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetColumnName(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.ColumnName = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetCreateTime(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.CreateTime = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetDatabaseName(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetDirectDownColumnNumber(v int32) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.DirectDownColumnNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetDirectDownTableNumber(v int32) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.DirectDownTableNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetDirectUpperColumnNumber(v int32) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.DirectUpperColumnNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetDirectUpperTableNumber(v int32) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.DirectUpperTableNumber = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetModifiedTime(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.ModifiedTime = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetSource(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.Source = &v
	return s
}

func (s *ListHiveColumnLineagesResponseBodyDataUpstreamLineages) SetTableName(v string) *ListHiveColumnLineagesResponseBodyDataUpstreamLineages {
	s.TableName = &v
	return s
}

type ListHiveColumnLineagesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHiveColumnLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHiveColumnLineagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHiveColumnLineagesResponse) GoString() string {
	return s.String()
}

func (s *ListHiveColumnLineagesResponse) SetHeaders(v map[string]*string) *ListHiveColumnLineagesResponse {
	s.Headers = v
	return s
}

func (s *ListHiveColumnLineagesResponse) SetStatusCode(v int32) *ListHiveColumnLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHiveColumnLineagesResponse) SetBody(v *ListHiveColumnLineagesResponseBody) *ListHiveColumnLineagesResponse {
	s.Body = v
	return s
}

type ListHiveTableLineagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pt_table_090901_emr
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveTableLineagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesRequest) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesRequest) SetClusterId(v string) *ListHiveTableLineagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListHiveTableLineagesRequest) SetDatabaseName(v string) *ListHiveTableLineagesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveTableLineagesRequest) SetTableName(v string) *ListHiveTableLineagesRequest {
	s.TableName = &v
	return s
}

type ListHiveTableLineagesResponseBody struct {
	Data *ListHiveTableLineagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// test
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHiveTableLineagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesResponseBody) SetData(v *ListHiveTableLineagesResponseBodyData) *ListHiveTableLineagesResponseBody {
	s.Data = v
	return s
}

func (s *ListHiveTableLineagesResponseBody) SetErrorCode(v string) *ListHiveTableLineagesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListHiveTableLineagesResponseBody) SetErrorMessage(v string) *ListHiveTableLineagesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListHiveTableLineagesResponseBody) SetRequestId(v string) *ListHiveTableLineagesResponseBody {
	s.RequestId = &v
	return s
}

type ListHiveTableLineagesResponseBodyData struct {
	DownstreamLineages []*ListHiveTableLineagesResponseBodyDataDownstreamLineages `json:"DownstreamLineages,omitempty" xml:"DownstreamLineages,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	DownstreamNumber *int32                                                   `json:"DownstreamNumber,omitempty" xml:"DownstreamNumber,omitempty"`
	UpstreamLineages []*ListHiveTableLineagesResponseBodyDataUpstreamLineages `json:"UpstreamLineages,omitempty" xml:"UpstreamLineages,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	UpstreamNumber *int32 `json:"UpstreamNumber,omitempty" xml:"UpstreamNumber,omitempty"`
}

func (s ListHiveTableLineagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesResponseBodyData) SetDownstreamLineages(v []*ListHiveTableLineagesResponseBodyDataDownstreamLineages) *ListHiveTableLineagesResponseBodyData {
	s.DownstreamLineages = v
	return s
}

func (s *ListHiveTableLineagesResponseBodyData) SetDownstreamNumber(v int32) *ListHiveTableLineagesResponseBodyData {
	s.DownstreamNumber = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyData) SetUpstreamLineages(v []*ListHiveTableLineagesResponseBodyDataUpstreamLineages) *ListHiveTableLineagesResponseBodyData {
	s.UpstreamLineages = v
	return s
}

func (s *ListHiveTableLineagesResponseBodyData) SetUpstreamNumber(v int32) *ListHiveTableLineagesResponseBodyData {
	s.UpstreamNumber = &v
	return s
}

type ListHiveTableLineagesResponseBodyDataDownstreamLineages struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2020-01-09 18:16:15
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// mr
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// job_1234567055_0006
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2020-01-09 18:16:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// create table parquet_example_0407 select id as one,name as two,salary as three  from PTtable_0407_emr
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// example:
	//
	// HIVE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// pt_table_090901_emr_orc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveTableLineagesResponseBodyDataDownstreamLineages) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesResponseBodyDataDownstreamLineages) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetClusterId(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.ClusterId = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetCreateTime(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.CreateTime = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetDatabaseName(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetEngine(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.Engine = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetJobId(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.JobId = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetModifiedTime(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.ModifiedTime = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetQueryText(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.QueryText = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetSource(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.Source = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataDownstreamLineages) SetTableName(v string) *ListHiveTableLineagesResponseBodyDataDownstreamLineages {
	s.TableName = &v
	return s
}

type ListHiveTableLineagesResponseBodyDataUpstreamLineages struct {
	// example:
	//
	// C-D033DD5FB82436A6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2020-01-09 18:16:15
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// mr
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// job_1234567055_0006
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 2020-01-09 18:16:37
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// create table parquet_example_0407 select id as one,name as two,salary as three  from PTtable_0407_emr
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// example:
	//
	// HIVE
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// pt_table_090901_emr_orc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListHiveTableLineagesResponseBodyDataUpstreamLineages) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesResponseBodyDataUpstreamLineages) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetClusterId(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.ClusterId = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetCreateTime(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.CreateTime = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetDatabaseName(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.DatabaseName = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetEngine(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.Engine = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetJobId(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.JobId = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetModifiedTime(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.ModifiedTime = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetQueryText(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.QueryText = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetSource(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.Source = &v
	return s
}

func (s *ListHiveTableLineagesResponseBodyDataUpstreamLineages) SetTableName(v string) *ListHiveTableLineagesResponseBodyDataUpstreamLineages {
	s.TableName = &v
	return s
}

type ListHiveTableLineagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHiveTableLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHiveTableLineagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHiveTableLineagesResponse) GoString() string {
	return s.String()
}

func (s *ListHiveTableLineagesResponse) SetHeaders(v map[string]*string) *ListHiveTableLineagesResponse {
	s.Headers = v
	return s
}

func (s *ListHiveTableLineagesResponse) SetStatusCode(v int32) *ListHiveTableLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHiveTableLineagesResponse) SetBody(v *ListHiveTableLineagesResponseBody) *ListHiveTableLineagesResponse {
	s.Body = v
	return s
}

type ListTablePartitionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C-2A51D3826C701234
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// ASC/DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListTablePartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablePartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListTablePartitionsRequest) SetClusterId(v string) *ListTablePartitionsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListTablePartitionsRequest) SetDatabaseName(v string) *ListTablePartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListTablePartitionsRequest) SetOrder(v string) *ListTablePartitionsRequest {
	s.Order = &v
	return s
}

func (s *ListTablePartitionsRequest) SetPageNumber(v int32) *ListTablePartitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablePartitionsRequest) SetPageSize(v int32) *ListTablePartitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablePartitionsRequest) SetTableName(v string) *ListTablePartitionsRequest {
	s.TableName = &v
	return s
}

type ListTablePartitionsResponseBody struct {
	Data *ListTablePartitionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTablePartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablePartitionsResponseBody) SetData(v *ListTablePartitionsResponseBodyData) *ListTablePartitionsResponseBody {
	s.Data = v
	return s
}

func (s *ListTablePartitionsResponseBody) SetErrorCode(v string) *ListTablePartitionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTablePartitionsResponseBody) SetErrorMessage(v string) *ListTablePartitionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTablePartitionsResponseBody) SetRequestId(v string) *ListTablePartitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListTablePartitionsResponseBodyData struct {
	// example:
	//
	// 10
	PageSize  *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PagedData []*ListTablePartitionsResponseBodyDataPagedData `json:"PagedData,omitempty" xml:"PagedData,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablePartitionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTablePartitionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablePartitionsResponseBodyData) SetPageSize(v int32) *ListTablePartitionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTablePartitionsResponseBodyData) SetPagedData(v []*ListTablePartitionsResponseBodyDataPagedData) *ListTablePartitionsResponseBodyData {
	s.PagedData = v
	return s
}

func (s *ListTablePartitionsResponseBodyData) SetTotalCount(v int32) *ListTablePartitionsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTablePartitionsResponseBodyDataPagedData struct {
	// example:
	//
	// 1568032253000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1568032253000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// hdfs://emr-header-1.cluster-136574:9000/user/hive/warehouse/pt_table_090901_emr_child/ds=20190909/hr=20/region=shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// test
	PartitionComment *string `json:"PartitionComment,omitempty" xml:"PartitionComment,omitempty"`
	// example:
	//
	// ds=20190909/hr=20/region=shanghai
	PartitionName *string `json:"PartitionName,omitempty" xml:"PartitionName,omitempty"`
	// example:
	//
	// hdfs://emr-header-1.cluster-136574:9000/user/hive/warehouse/pt_table_090901_emr_child/ds=20190909/hr=20/region=shanghai
	PartitionPath *string `json:"PartitionPath,omitempty" xml:"PartitionPath,omitempty"`
	// example:
	//
	// HIVE
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
}

func (s ListTablePartitionsResponseBodyDataPagedData) String() string {
	return tea.Prettify(s)
}

func (s ListTablePartitionsResponseBodyDataPagedData) GoString() string {
	return s.String()
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetGmtCreate(v int64) *ListTablePartitionsResponseBodyDataPagedData {
	s.GmtCreate = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetGmtModified(v int64) *ListTablePartitionsResponseBodyDataPagedData {
	s.GmtModified = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetLocation(v string) *ListTablePartitionsResponseBodyDataPagedData {
	s.Location = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetPartitionComment(v string) *ListTablePartitionsResponseBodyDataPagedData {
	s.PartitionComment = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetPartitionName(v string) *ListTablePartitionsResponseBodyDataPagedData {
	s.PartitionName = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetPartitionPath(v string) *ListTablePartitionsResponseBodyDataPagedData {
	s.PartitionPath = &v
	return s
}

func (s *ListTablePartitionsResponseBodyDataPagedData) SetPartitionType(v string) *ListTablePartitionsResponseBodyDataPagedData {
	s.PartitionType = &v
	return s
}

type ListTablePartitionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTablePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTablePartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablePartitionsResponse) GoString() string {
	return s.String()
}

func (s *ListTablePartitionsResponse) SetHeaders(v map[string]*string) *ListTablePartitionsResponse {
	s.Headers = v
	return s
}

func (s *ListTablePartitionsResponse) SetStatusCode(v int32) *ListTablePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablePartitionsResponse) SetBody(v *ListTablePartitionsResponseBody) *ListTablePartitionsResponse {
	s.Body = v
	return s
}

type OpenDataWorksStandardServiceRequest struct {
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s OpenDataWorksStandardServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDataWorksStandardServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenDataWorksStandardServiceRequest) SetRegion(v string) *OpenDataWorksStandardServiceRequest {
	s.Region = &v
	return s
}

type OpenDataWorksStandardServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDataWorksStandardServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDataWorksStandardServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDataWorksStandardServiceResponseBody) SetOrderId(v string) *OpenDataWorksStandardServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenDataWorksStandardServiceResponseBody) SetRequestId(v string) *OpenDataWorksStandardServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenDataWorksStandardServiceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDataWorksStandardServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDataWorksStandardServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDataWorksStandardServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenDataWorksStandardServiceResponse) SetHeaders(v map[string]*string) *OpenDataWorksStandardServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenDataWorksStandardServiceResponse) SetStatusCode(v int32) *OpenDataWorksStandardServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDataWorksStandardServiceResponse) SetBody(v *OpenDataWorksStandardServiceResponseBody) *OpenDataWorksStandardServiceResponse {
	s.Body = v
	return s
}

type SearchManualDagNodeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123434234
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_odps_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SearchManualDagNodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchManualDagNodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *SearchManualDagNodeInstanceRequest) SetDagId(v int64) *SearchManualDagNodeInstanceRequest {
	s.DagId = &v
	return s
}

func (s *SearchManualDagNodeInstanceRequest) SetProjectName(v string) *SearchManualDagNodeInstanceRequest {
	s.ProjectName = &v
	return s
}

type SearchManualDagNodeInstanceResponseBody struct {
	Data *SearchManualDagNodeInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// test
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// 2d9ced66-38ef-4923-baf6-391dd3a7e656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchManualDagNodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchManualDagNodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchManualDagNodeInstanceResponseBody) SetData(v *SearchManualDagNodeInstanceResponseBodyData) *SearchManualDagNodeInstanceResponseBody {
	s.Data = v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBody) SetErrCode(v string) *SearchManualDagNodeInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBody) SetErrMsg(v string) *SearchManualDagNodeInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBody) SetRequestId(v string) *SearchManualDagNodeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBody) SetSuccess(v bool) *SearchManualDagNodeInstanceResponseBody {
	s.Success = &v
	return s
}

type SearchManualDagNodeInstanceResponseBodyData struct {
	NodeInsInfo []*SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo `json:"NodeInsInfo,omitempty" xml:"NodeInsInfo,omitempty" type:"Repeated"`
}

func (s SearchManualDagNodeInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchManualDagNodeInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchManualDagNodeInstanceResponseBodyData) SetNodeInsInfo(v []*SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) *SearchManualDagNodeInstanceResponseBodyData {
	s.NodeInsInfo = v
	return s
}

type SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo struct {
	// example:
	//
	// 2018-12-12 00:00:00
	BeginRunningTime *string `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	BeginWaitResTime *string `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	BeginWaitTimeTime *string `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12434232423
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// example:
	//
	// 5
	DagType *int32 `json:"DagType,omitempty" xml:"DagType,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 12322434112
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2018-12-12 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// param_k1=param_v1
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// example:
	//
	// 6
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) GoString() string {
	return s.String()
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetBeginRunningTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.BeginRunningTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetBeginWaitResTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.BeginWaitResTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetBeginWaitTimeTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetBizdate(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.Bizdate = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetCreateTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.CreateTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetDagId(v int64) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.DagId = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetDagType(v int32) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.DagType = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetFinishTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.FinishTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetInstanceId(v int64) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.InstanceId = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetModifyTime(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.ModifyTime = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetNodeName(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.NodeName = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetParaValue(v string) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.ParaValue = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo) SetStatus(v int32) *SearchManualDagNodeInstanceResponseBodyDataNodeInsInfo {
	s.Status = &v
	return s
}

type SearchManualDagNodeInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchManualDagNodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchManualDagNodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchManualDagNodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *SearchManualDagNodeInstanceResponse) SetHeaders(v map[string]*string) *SearchManualDagNodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *SearchManualDagNodeInstanceResponse) SetStatusCode(v int32) *SearchManualDagNodeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchManualDagNodeInstanceResponse) SetBody(v *SearchManualDagNodeInstanceResponseBody) *SearchManualDagNodeInstanceResponse {
	s.Body = v
	return s
}

type SendTaskMetaCallbackRequest struct {
	// This parameter is required.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	ConnectionInfo *string `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty"`
	// This parameter is required.
	EndTime   *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// This parameter is required.
	TaskEnvParam *string `json:"TaskEnvParam,omitempty" xml:"TaskEnvParam,omitempty"`
	// This parameter is required.
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s SendTaskMetaCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTaskMetaCallbackRequest) GoString() string {
	return s.String()
}

func (s *SendTaskMetaCallbackRequest) SetCode(v string) *SendTaskMetaCallbackRequest {
	s.Code = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetConnectionInfo(v string) *SendTaskMetaCallbackRequest {
	s.ConnectionInfo = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetEndTime(v int64) *SendTaskMetaCallbackRequest {
	s.EndTime = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetResources(v []*string) *SendTaskMetaCallbackRequest {
	s.Resources = v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetStartTime(v int64) *SendTaskMetaCallbackRequest {
	s.StartTime = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetSubType(v string) *SendTaskMetaCallbackRequest {
	s.SubType = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetTaskEnvParam(v string) *SendTaskMetaCallbackRequest {
	s.TaskEnvParam = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetTenantId(v int64) *SendTaskMetaCallbackRequest {
	s.TenantId = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetType(v string) *SendTaskMetaCallbackRequest {
	s.Type = &v
	return s
}

func (s *SendTaskMetaCallbackRequest) SetUser(v string) *SendTaskMetaCallbackRequest {
	s.User = &v
	return s
}

type SendTaskMetaCallbackResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ErrorCode *int64  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTaskMetaCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTaskMetaCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SendTaskMetaCallbackResponseBody) SetData(v string) *SendTaskMetaCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *SendTaskMetaCallbackResponseBody) SetErrMsg(v string) *SendTaskMetaCallbackResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SendTaskMetaCallbackResponseBody) SetErrorCode(v int64) *SendTaskMetaCallbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendTaskMetaCallbackResponseBody) SetRequestId(v string) *SendTaskMetaCallbackResponseBody {
	s.RequestId = &v
	return s
}

type SendTaskMetaCallbackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTaskMetaCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTaskMetaCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTaskMetaCallbackResponse) GoString() string {
	return s.String()
}

func (s *SendTaskMetaCallbackResponse) SetHeaders(v map[string]*string) *SendTaskMetaCallbackResponse {
	s.Headers = v
	return s
}

func (s *SendTaskMetaCallbackResponse) SetStatusCode(v int32) *SendTaskMetaCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTaskMetaCallbackResponse) SetBody(v *SendTaskMetaCallbackResponseBody) *SendTaskMetaCallbackResponse {
	s.Body = v
	return s
}

type SetSwitchValueRequest struct {
	SwitchName  *string `json:"SwitchName,omitempty" xml:"SwitchName,omitempty"`
	SwitchValue *string `json:"SwitchValue,omitempty" xml:"SwitchValue,omitempty"`
}

func (s SetSwitchValueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSwitchValueRequest) GoString() string {
	return s.String()
}

func (s *SetSwitchValueRequest) SetSwitchName(v string) *SetSwitchValueRequest {
	s.SwitchName = &v
	return s
}

func (s *SetSwitchValueRequest) SetSwitchValue(v string) *SetSwitchValueRequest {
	s.SwitchValue = &v
	return s
}

type SetSwitchValueResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSwitchValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSwitchValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetSwitchValueResponseBody) SetData(v bool) *SetSwitchValueResponseBody {
	s.Data = &v
	return s
}

func (s *SetSwitchValueResponseBody) SetRequestId(v string) *SetSwitchValueResponseBody {
	s.RequestId = &v
	return s
}

type SetSwitchValueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSwitchValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSwitchValueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSwitchValueResponse) GoString() string {
	return s.String()
}

func (s *SetSwitchValueResponse) SetHeaders(v map[string]*string) *SetSwitchValueResponse {
	s.Headers = v
	return s
}

func (s *SetSwitchValueResponse) SetStatusCode(v int32) *SetSwitchValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSwitchValueResponse) SetBody(v *SetSwitchValueResponseBody) *SetSwitchValueResponse {
	s.Body = v
	return s
}

type StartCollectQualityRequest struct {
	// This parameter is required.
	CallbackResultString *string `json:"CallbackResultString,omitempty" xml:"CallbackResultString,omitempty"`
}

func (s StartCollectQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityRequest) GoString() string {
	return s.String()
}

func (s *StartCollectQualityRequest) SetCallbackResultString(v string) *StartCollectQualityRequest {
	s.CallbackResultString = &v
	return s
}

type StartCollectQualityResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReturnCode  *string                                       `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	ReturnValue []*StartCollectQualityResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Repeated"`
}

func (s StartCollectQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityResponseBody) GoString() string {
	return s.String()
}

func (s *StartCollectQualityResponseBody) SetRequestId(v string) *StartCollectQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCollectQualityResponseBody) SetReturnCode(v string) *StartCollectQualityResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *StartCollectQualityResponseBody) SetReturnValue(v []*StartCollectQualityResponseBodyReturnValue) *StartCollectQualityResponseBody {
	s.ReturnValue = v
	return s
}

type StartCollectQualityResponseBodyReturnValue struct {
	ActualExpression *string                                                      `json:"ActualExpression,omitempty" xml:"ActualExpression,omitempty"`
	BizDate          *string                                                      `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	CallbackUrl      *string                                                      `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Connection       *string                                                      `json:"Connection,omitempty" xml:"Connection,omitempty"`
	EntityId         *int64                                                       `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MatchExpression  *string                                                      `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	PluginName       *string                                                      `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	StrongMethodSet  []*StartCollectQualityResponseBodyReturnValueStrongMethodSet `json:"StrongMethodSet,omitempty" xml:"StrongMethodSet,omitempty" type:"Repeated"`
	TableGuid        *string                                                      `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TaskId           *string                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WeakMethodSet    []*StartCollectQualityResponseBodyReturnValueWeakMethodSet   `json:"WeakMethodSet,omitempty" xml:"WeakMethodSet,omitempty" type:"Repeated"`
}

func (s StartCollectQualityResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *StartCollectQualityResponseBodyReturnValue) SetActualExpression(v string) *StartCollectQualityResponseBodyReturnValue {
	s.ActualExpression = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetBizDate(v string) *StartCollectQualityResponseBodyReturnValue {
	s.BizDate = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetCallbackUrl(v string) *StartCollectQualityResponseBodyReturnValue {
	s.CallbackUrl = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetConnection(v string) *StartCollectQualityResponseBodyReturnValue {
	s.Connection = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetEntityId(v int64) *StartCollectQualityResponseBodyReturnValue {
	s.EntityId = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetMatchExpression(v string) *StartCollectQualityResponseBodyReturnValue {
	s.MatchExpression = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetPluginName(v string) *StartCollectQualityResponseBodyReturnValue {
	s.PluginName = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetStrongMethodSet(v []*StartCollectQualityResponseBodyReturnValueStrongMethodSet) *StartCollectQualityResponseBodyReturnValue {
	s.StrongMethodSet = v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetTableGuid(v string) *StartCollectQualityResponseBodyReturnValue {
	s.TableGuid = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetTaskId(v string) *StartCollectQualityResponseBodyReturnValue {
	s.TaskId = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValue) SetWeakMethodSet(v []*StartCollectQualityResponseBodyReturnValueWeakMethodSet) *StartCollectQualityResponseBodyReturnValue {
	s.WeakMethodSet = v
	return s
}

type StartCollectQualityResponseBodyReturnValueStrongMethodSet struct {
	ColName      *string `json:"ColName,omitempty" xml:"ColName,omitempty"`
	IsColRule    *bool   `json:"IsColRule,omitempty" xml:"IsColRule,omitempty"`
	IsSqlRule    *bool   `json:"IsSqlRule,omitempty" xml:"IsSqlRule,omitempty"`
	IsStrongRule *bool   `json:"IsStrongRule,omitempty" xml:"IsStrongRule,omitempty"`
	MethodName   *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s StartCollectQualityResponseBodyReturnValueStrongMethodSet) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityResponseBodyReturnValueStrongMethodSet) GoString() string {
	return s.String()
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetColName(v string) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.ColName = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetIsColRule(v bool) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.IsColRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetIsSqlRule(v bool) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.IsSqlRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetIsStrongRule(v bool) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.IsStrongRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetMethodName(v string) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.MethodName = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueStrongMethodSet) SetRuleId(v int64) *StartCollectQualityResponseBodyReturnValueStrongMethodSet {
	s.RuleId = &v
	return s
}

type StartCollectQualityResponseBodyReturnValueWeakMethodSet struct {
	ColName      *string `json:"ColName,omitempty" xml:"ColName,omitempty"`
	IsColRule    *bool   `json:"IsColRule,omitempty" xml:"IsColRule,omitempty"`
	IsSqlRule    *bool   `json:"IsSqlRule,omitempty" xml:"IsSqlRule,omitempty"`
	IsStrongRule *bool   `json:"IsStrongRule,omitempty" xml:"IsStrongRule,omitempty"`
	MethodName   *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s StartCollectQualityResponseBodyReturnValueWeakMethodSet) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityResponseBodyReturnValueWeakMethodSet) GoString() string {
	return s.String()
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetColName(v string) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.ColName = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetIsColRule(v bool) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.IsColRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetIsSqlRule(v bool) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.IsSqlRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetIsStrongRule(v bool) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.IsStrongRule = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetMethodName(v string) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.MethodName = &v
	return s
}

func (s *StartCollectQualityResponseBodyReturnValueWeakMethodSet) SetRuleId(v int64) *StartCollectQualityResponseBodyReturnValueWeakMethodSet {
	s.RuleId = &v
	return s
}

type StartCollectQualityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCollectQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCollectQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCollectQualityResponse) GoString() string {
	return s.String()
}

func (s *StartCollectQualityResponse) SetHeaders(v map[string]*string) *StartCollectQualityResponse {
	s.Headers = v
	return s
}

func (s *StartCollectQualityResponse) SetStatusCode(v int32) *StartCollectQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCollectQualityResponse) SetBody(v *StartCollectQualityResponseBody) *StartCollectQualityResponse {
	s.Body = v
	return s
}

type StartDoCheckQualityRequest struct {
	// This parameter is required.
	CallbackResultString *string `json:"CallbackResultString,omitempty" xml:"CallbackResultString,omitempty"`
}

func (s StartDoCheckQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDoCheckQualityRequest) GoString() string {
	return s.String()
}

func (s *StartDoCheckQualityRequest) SetCallbackResultString(v string) *StartDoCheckQualityRequest {
	s.CallbackResultString = &v
	return s
}

type StartDoCheckQualityResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReturnCode  *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	ReturnValue *bool   `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty"`
}

func (s StartDoCheckQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDoCheckQualityResponseBody) GoString() string {
	return s.String()
}

func (s *StartDoCheckQualityResponseBody) SetRequestId(v string) *StartDoCheckQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDoCheckQualityResponseBody) SetReturnCode(v string) *StartDoCheckQualityResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *StartDoCheckQualityResponseBody) SetReturnValue(v bool) *StartDoCheckQualityResponseBody {
	s.ReturnValue = &v
	return s
}

type StartDoCheckQualityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDoCheckQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDoCheckQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDoCheckQualityResponse) GoString() string {
	return s.String()
}

func (s *StartDoCheckQualityResponse) SetHeaders(v map[string]*string) *StartDoCheckQualityResponse {
	s.Headers = v
	return s
}

func (s *StartDoCheckQualityResponse) SetStatusCode(v int32) *StartDoCheckQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDoCheckQualityResponse) SetBody(v *StartDoCheckQualityResponseBody) *StartDoCheckQualityResponse {
	s.Body = v
	return s
}

type StartTaskQualityRequest struct {
	// This parameter is required.
	CallbackResultString *string `json:"CallbackResultString,omitempty" xml:"CallbackResultString,omitempty"`
}

func (s StartTaskQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityRequest) GoString() string {
	return s.String()
}

func (s *StartTaskQualityRequest) SetCallbackResultString(v string) *StartTaskQualityRequest {
	s.CallbackResultString = &v
	return s
}

type StartTaskQualityResponseBody struct {
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReturnCode  *string                                  `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	ReturnValue *StartTaskQualityResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Struct"`
}

func (s StartTaskQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskQualityResponseBody) SetRequestId(v string) *StartTaskQualityResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskQualityResponseBody) SetReturnCode(v string) *StartTaskQualityResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *StartTaskQualityResponseBody) SetReturnValue(v *StartTaskQualityResponseBodyReturnValue) *StartTaskQualityResponseBody {
	s.ReturnValue = v
	return s
}

type StartTaskQualityResponseBodyReturnValue struct {
	ActualExpression *string                                                   `json:"ActualExpression,omitempty" xml:"ActualExpression,omitempty"`
	BizDate          *string                                                   `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	CallbackUrl      *string                                                   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	Connection       *string                                                   `json:"Connection,omitempty" xml:"Connection,omitempty"`
	EntityId         *int64                                                    `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	MatchExpression  *string                                                   `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	PluginName       *string                                                   `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	StatisticsFlag   *int64                                                    `json:"StatisticsFlag,omitempty" xml:"StatisticsFlag,omitempty"`
	StrongMethodSet  []*StartTaskQualityResponseBodyReturnValueStrongMethodSet `json:"StrongMethodSet,omitempty" xml:"StrongMethodSet,omitempty" type:"Repeated"`
	TableGuid        *string                                                   `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TaskId           *string                                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TriggerFlag      *int64                                                    `json:"TriggerFlag,omitempty" xml:"TriggerFlag,omitempty"`
	WeakMethodSet    []*StartTaskQualityResponseBodyReturnValueWeakMethodSet   `json:"WeakMethodSet,omitempty" xml:"WeakMethodSet,omitempty" type:"Repeated"`
}

func (s StartTaskQualityResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *StartTaskQualityResponseBodyReturnValue) SetActualExpression(v string) *StartTaskQualityResponseBodyReturnValue {
	s.ActualExpression = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetBizDate(v string) *StartTaskQualityResponseBodyReturnValue {
	s.BizDate = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetCallbackUrl(v string) *StartTaskQualityResponseBodyReturnValue {
	s.CallbackUrl = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetConnection(v string) *StartTaskQualityResponseBodyReturnValue {
	s.Connection = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetEntityId(v int64) *StartTaskQualityResponseBodyReturnValue {
	s.EntityId = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetMatchExpression(v string) *StartTaskQualityResponseBodyReturnValue {
	s.MatchExpression = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetPluginName(v string) *StartTaskQualityResponseBodyReturnValue {
	s.PluginName = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetStatisticsFlag(v int64) *StartTaskQualityResponseBodyReturnValue {
	s.StatisticsFlag = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetStrongMethodSet(v []*StartTaskQualityResponseBodyReturnValueStrongMethodSet) *StartTaskQualityResponseBodyReturnValue {
	s.StrongMethodSet = v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetTableGuid(v string) *StartTaskQualityResponseBodyReturnValue {
	s.TableGuid = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetTaskId(v string) *StartTaskQualityResponseBodyReturnValue {
	s.TaskId = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetTriggerFlag(v int64) *StartTaskQualityResponseBodyReturnValue {
	s.TriggerFlag = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValue) SetWeakMethodSet(v []*StartTaskQualityResponseBodyReturnValueWeakMethodSet) *StartTaskQualityResponseBodyReturnValue {
	s.WeakMethodSet = v
	return s
}

type StartTaskQualityResponseBodyReturnValueStrongMethodSet struct {
	ColName      *string `json:"ColName,omitempty" xml:"ColName,omitempty"`
	IsColRule    *bool   `json:"IsColRule,omitempty" xml:"IsColRule,omitempty"`
	IsSqlRule    *bool   `json:"IsSqlRule,omitempty" xml:"IsSqlRule,omitempty"`
	IsStrongRule *bool   `json:"IsStrongRule,omitempty" xml:"IsStrongRule,omitempty"`
	MethodName   *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s StartTaskQualityResponseBodyReturnValueStrongMethodSet) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityResponseBodyReturnValueStrongMethodSet) GoString() string {
	return s.String()
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetColName(v string) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.ColName = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetIsColRule(v bool) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.IsColRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetIsSqlRule(v bool) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.IsSqlRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetIsStrongRule(v bool) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.IsStrongRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetMethodName(v string) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.MethodName = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueStrongMethodSet) SetRuleId(v int64) *StartTaskQualityResponseBodyReturnValueStrongMethodSet {
	s.RuleId = &v
	return s
}

type StartTaskQualityResponseBodyReturnValueWeakMethodSet struct {
	ColName      *string `json:"ColName,omitempty" xml:"ColName,omitempty"`
	IsColRule    *bool   `json:"IsColRule,omitempty" xml:"IsColRule,omitempty"`
	IsSqlRule    *bool   `json:"IsSqlRule,omitempty" xml:"IsSqlRule,omitempty"`
	IsStrongRule *bool   `json:"IsStrongRule,omitempty" xml:"IsStrongRule,omitempty"`
	MethodName   *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s StartTaskQualityResponseBodyReturnValueWeakMethodSet) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityResponseBodyReturnValueWeakMethodSet) GoString() string {
	return s.String()
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetColName(v string) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.ColName = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetIsColRule(v bool) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.IsColRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetIsSqlRule(v bool) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.IsSqlRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetIsStrongRule(v bool) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.IsStrongRule = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetMethodName(v string) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.MethodName = &v
	return s
}

func (s *StartTaskQualityResponseBodyReturnValueWeakMethodSet) SetRuleId(v int64) *StartTaskQualityResponseBodyReturnValueWeakMethodSet {
	s.RuleId = &v
	return s
}

type StartTaskQualityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTaskQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTaskQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTaskQualityResponse) GoString() string {
	return s.String()
}

func (s *StartTaskQualityResponse) SetHeaders(v map[string]*string) *StartTaskQualityResponse {
	s.Headers = v
	return s
}

func (s *StartTaskQualityResponse) SetStatusCode(v int32) *StartTaskQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTaskQualityResponse) SetBody(v *StartTaskQualityResponseBody) *StartTaskQualityResponse {
	s.Body = v
	return s
}

type TriggerDataLoaderResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerDataLoaderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerDataLoaderResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerDataLoaderResponseBody) SetData(v bool) *TriggerDataLoaderResponseBody {
	s.Data = &v
	return s
}

func (s *TriggerDataLoaderResponseBody) SetRequestId(v string) *TriggerDataLoaderResponseBody {
	s.RequestId = &v
	return s
}

type TriggerDataLoaderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerDataLoaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerDataLoaderResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerDataLoaderResponse) GoString() string {
	return s.String()
}

func (s *TriggerDataLoaderResponse) SetHeaders(v map[string]*string) *TriggerDataLoaderResponse {
	s.Headers = v
	return s
}

func (s *TriggerDataLoaderResponse) SetStatusCode(v int32) *TriggerDataLoaderResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerDataLoaderResponse) SetBody(v *TriggerDataLoaderResponseBody) *TriggerDataLoaderResponse {
	s.Body = v
	return s
}

type TriggerTimeMachineTaskResponseBody struct {
	Data      *TriggerTimeMachineTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                                 `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMsg    *string                                 `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerTimeMachineTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerTimeMachineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerTimeMachineTaskResponseBody) SetData(v *TriggerTimeMachineTaskResponseBodyData) *TriggerTimeMachineTaskResponseBody {
	s.Data = v
	return s
}

func (s *TriggerTimeMachineTaskResponseBody) SetErrCode(v string) *TriggerTimeMachineTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBody) SetErrMsg(v string) *TriggerTimeMachineTaskResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBody) SetRequestId(v string) *TriggerTimeMachineTaskResponseBody {
	s.RequestId = &v
	return s
}

type TriggerTimeMachineTaskResponseBodyData struct {
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ObjId       *string `json:"ObjId,omitempty" xml:"ObjId,omitempty"`
	ObjName     *string `json:"ObjName,omitempty" xml:"ObjName,omitempty"`
	OperType    *string `json:"OperType,omitempty" xml:"OperType,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s TriggerTimeMachineTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TriggerTimeMachineTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetGmtCreate(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetGmtModified(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetHostName(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.HostName = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetId(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetObjId(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.ObjId = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetObjName(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.ObjName = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetOperType(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.OperType = &v
	return s
}

func (s *TriggerTimeMachineTaskResponseBodyData) SetStatus(v string) *TriggerTimeMachineTaskResponseBodyData {
	s.Status = &v
	return s
}

type TriggerTimeMachineTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerTimeMachineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerTimeMachineTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerTimeMachineTaskResponse) GoString() string {
	return s.String()
}

func (s *TriggerTimeMachineTaskResponse) SetHeaders(v map[string]*string) *TriggerTimeMachineTaskResponse {
	s.Headers = v
	return s
}

func (s *TriggerTimeMachineTaskResponse) SetStatusCode(v int32) *TriggerTimeMachineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerTimeMachineTaskResponse) SetBody(v *TriggerTimeMachineTaskResponseBody) *TriggerTimeMachineTaskResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            tea.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           tea.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dataworks.aliyuncs.com"),
		"eu-central-1":          tea.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateManualDagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateManualDagResponse
func (client *Client) CreateManualDagWithOptions(request *CreateManualDagRequest, runtime *util.RuntimeOptions) (_result *CreateManualDagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bizdate)) {
		query["Bizdate"] = request.Bizdate
	}

	if !tea.BoolValue(util.IsUnset(request.DagPara)) {
		query["DagPara"] = request.DagPara
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		query["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.NodePara)) {
		query["NodePara"] = request.NodePara
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateManualDag"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateManualDagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateManualDagRequest
//
// @return CreateManualDagResponse
func (client *Client) CreateManualDag(request *CreateManualDagRequest) (_result *CreateManualDagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateManualDagResponse{}
	_body, _err := client.CreateManualDagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectIdentifier)) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeEmrHiveTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEmrHiveTableResponse
func (client *Client) DescribeEmrHiveTableWithOptions(request *DescribeEmrHiveTableRequest, runtime *util.RuntimeOptions) (_result *DescribeEmrHiveTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEmrHiveTable"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEmrHiveTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeEmrHiveTableRequest
//
// @return DescribeEmrHiveTableResponse
func (client *Client) DescribeEmrHiveTable(request *DescribeEmrHiveTableRequest) (_result *DescribeEmrHiveTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEmrHiveTableResponse{}
	_body, _err := client.DescribeEmrHiveTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询apiContext接口
//
// @param request - GetDataServiceApiContextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiContextResponse
func (client *Client) GetDataServiceApiContextWithOptions(request *GetDataServiceApiContextRequest, runtime *util.RuntimeOptions) (_result *GetDataServiceApiContextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataServiceApiContext"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataServiceApiContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询apiContext接口
//
// @param request - GetDataServiceApiContextRequest
//
// @return GetDataServiceApiContextResponse
func (client *Client) GetDataServiceApiContext(request *GetDataServiceApiContextRequest) (_result *GetDataServiceApiContextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataServiceApiContextResponse{}
	_body, _err := client.GetDataServiceApiContextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询apiContext更新事件接口
//
// @param request - GetDataServiceContextUpdateEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceContextUpdateEventResponse
func (client *Client) GetDataServiceContextUpdateEventWithOptions(runtime *util.RuntimeOptions) (_result *GetDataServiceContextUpdateEventResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetDataServiceContextUpdateEvent"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataServiceContextUpdateEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询apiContext更新事件接口
//
// @return GetDataServiceContextUpdateEventResponse
func (client *Client) GetDataServiceContextUpdateEvent() (_result *GetDataServiceContextUpdateEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataServiceContextUpdateEventResponse{}
	_body, _err := client.GetDataServiceContextUpdateEventWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据Switch名称获取值
//
// @param request - GetSwitchValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSwitchValueResponse
func (client *Client) GetSwitchValueWithOptions(request *GetSwitchValueRequest, runtime *util.RuntimeOptions) (_result *GetSwitchValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SwitchName)) {
		query["SwitchName"] = request.SwitchName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSwitchValue"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSwitchValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据Switch名称获取值
//
// @param request - GetSwitchValueRequest
//
// @return GetSwitchValueResponse
func (client *Client) GetSwitchValue(request *GetSwitchValueRequest) (_result *GetSwitchValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSwitchValueResponse{}
	_body, _err := client.GetSwitchValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询timeMachine任务详情
//
// @param request - GetTimeMachineTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTimeMachineTaskResponse
func (client *Client) GetTimeMachineTaskWithOptions(request *GetTimeMachineTaskRequest, runtime *util.RuntimeOptions) (_result *GetTimeMachineTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTimeMachineTask"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTimeMachineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询timeMachine任务详情
//
// @param request - GetTimeMachineTaskRequest
//
// @return GetTimeMachineTaskResponse
func (client *Client) GetTimeMachineTask(request *GetTimeMachineTaskRequest) (_result *GetTimeMachineTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTimeMachineTaskResponse{}
	_body, _err := client.GetTimeMachineTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListEmrHiveAuditLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEmrHiveAuditLogsResponse
func (client *Client) ListEmrHiveAuditLogsWithOptions(request *ListEmrHiveAuditLogsRequest, runtime *util.RuntimeOptions) (_result *ListEmrHiveAuditLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEmrHiveAuditLogs"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmrHiveAuditLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEmrHiveAuditLogsRequest
//
// @return ListEmrHiveAuditLogsResponse
func (client *Client) ListEmrHiveAuditLogs(request *ListEmrHiveAuditLogsRequest) (_result *ListEmrHiveAuditLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEmrHiveAuditLogsResponse{}
	_body, _err := client.ListEmrHiveAuditLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListEmrHiveDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEmrHiveDatabasesResponse
func (client *Client) ListEmrHiveDatabasesWithOptions(request *ListEmrHiveDatabasesRequest, runtime *util.RuntimeOptions) (_result *ListEmrHiveDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEmrHiveDatabases"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmrHiveDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEmrHiveDatabasesRequest
//
// @return ListEmrHiveDatabasesResponse
func (client *Client) ListEmrHiveDatabases(request *ListEmrHiveDatabasesRequest) (_result *ListEmrHiveDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEmrHiveDatabasesResponse{}
	_body, _err := client.ListEmrHiveDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListEmrHiveTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEmrHiveTablesResponse
func (client *Client) ListEmrHiveTablesWithOptions(request *ListEmrHiveTablesRequest, runtime *util.RuntimeOptions) (_result *ListEmrHiveTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEmrHiveTables"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmrHiveTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEmrHiveTablesRequest
//
// @return ListEmrHiveTablesResponse
func (client *Client) ListEmrHiveTables(request *ListEmrHiveTablesRequest) (_result *ListEmrHiveTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEmrHiveTablesResponse{}
	_body, _err := client.ListEmrHiveTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据服务API
//
// @param request - ListGovernanceIssueDataServiceAPIsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGovernanceIssueDataServiceAPIsResponse
func (client *Client) ListGovernanceIssueDataServiceAPIsWithOptions(request *ListGovernanceIssueDataServiceAPIsRequest, runtime *util.RuntimeOptions) (_result *ListGovernanceIssueDataServiceAPIsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizDate)) {
		body["BizDate"] = request.BizDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCategory)) {
		body["RuleCategory"] = request.RuleCategory
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGovernanceIssueDataServiceAPIs"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGovernanceIssueDataServiceAPIsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据服务API
//
// @param request - ListGovernanceIssueDataServiceAPIsRequest
//
// @return ListGovernanceIssueDataServiceAPIsResponse
func (client *Client) ListGovernanceIssueDataServiceAPIs(request *ListGovernanceIssueDataServiceAPIsRequest) (_result *ListGovernanceIssueDataServiceAPIsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGovernanceIssueDataServiceAPIsResponse{}
	_body, _err := client.ListGovernanceIssueDataServiceAPIsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询治理项问题详情
//
// @param request - ListGovernanceIssueTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGovernanceIssueTablesResponse
func (client *Client) ListGovernanceIssueTablesWithOptions(request *ListGovernanceIssueTablesRequest, runtime *util.RuntimeOptions) (_result *ListGovernanceIssueTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizDate)) {
		body["BizDate"] = request.BizDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCategory)) {
		body["RuleCategory"] = request.RuleCategory
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGovernanceIssueTables"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGovernanceIssueTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询治理项问题详情
//
// @param request - ListGovernanceIssueTablesRequest
//
// @return ListGovernanceIssueTablesResponse
func (client *Client) ListGovernanceIssueTables(request *ListGovernanceIssueTablesRequest) (_result *ListGovernanceIssueTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGovernanceIssueTablesResponse{}
	_body, _err := client.ListGovernanceIssueTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询治理项-任务问题详情
//
// @param request - ListGovernanceIssueTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGovernanceIssueTasksResponse
func (client *Client) ListGovernanceIssueTasksWithOptions(request *ListGovernanceIssueTasksRequest, runtime *util.RuntimeOptions) (_result *ListGovernanceIssueTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizDate)) {
		body["BizDate"] = request.BizDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCategory)) {
		body["RuleCategory"] = request.RuleCategory
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGovernanceIssueTasks"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGovernanceIssueTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询治理项-任务问题详情
//
// @param request - ListGovernanceIssueTasksRequest
//
// @return ListGovernanceIssueTasksResponse
func (client *Client) ListGovernanceIssueTasks(request *ListGovernanceIssueTasksRequest) (_result *ListGovernanceIssueTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGovernanceIssueTasksResponse{}
	_body, _err := client.ListGovernanceIssueTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询治理项定义信息
//
// @param request - ListGovernanceRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGovernanceRulesResponse
func (client *Client) ListGovernanceRulesWithOptions(request *ListGovernanceRulesRequest, runtime *util.RuntimeOptions) (_result *ListGovernanceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.IssueType)) {
		body["IssueType"] = request.IssueType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGovernanceRules"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGovernanceRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询治理项定义信息
//
// @param request - ListGovernanceRulesRequest
//
// @return ListGovernanceRulesResponse
func (client *Client) ListGovernanceRules(request *ListGovernanceRulesRequest) (_result *ListGovernanceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGovernanceRulesResponse{}
	_body, _err := client.ListGovernanceRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListHiveColumnLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHiveColumnLineagesResponse
func (client *Client) ListHiveColumnLineagesWithOptions(request *ListHiveColumnLineagesRequest, runtime *util.RuntimeOptions) (_result *ListHiveColumnLineagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnName)) {
		query["ColumnName"] = request.ColumnName
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHiveColumnLineages"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHiveColumnLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListHiveColumnLineagesRequest
//
// @return ListHiveColumnLineagesResponse
func (client *Client) ListHiveColumnLineages(request *ListHiveColumnLineagesRequest) (_result *ListHiveColumnLineagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHiveColumnLineagesResponse{}
	_body, _err := client.ListHiveColumnLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListHiveTableLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHiveTableLineagesResponse
func (client *Client) ListHiveTableLineagesWithOptions(request *ListHiveTableLineagesRequest, runtime *util.RuntimeOptions) (_result *ListHiveTableLineagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHiveTableLineages"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHiveTableLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListHiveTableLineagesRequest
//
// @return ListHiveTableLineagesResponse
func (client *Client) ListHiveTableLineages(request *ListHiveTableLineagesRequest) (_result *ListHiveTableLineagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHiveTableLineagesResponse{}
	_body, _err := client.ListHiveTableLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTablePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablePartitionsResponse
func (client *Client) ListTablePartitionsWithOptions(request *ListTablePartitionsRequest, runtime *util.RuntimeOptions) (_result *ListTablePartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTablePartitions"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTablePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTablePartitionsRequest
//
// @return ListTablePartitionsResponse
func (client *Client) ListTablePartitions(request *ListTablePartitionsRequest) (_result *ListTablePartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTablePartitionsResponse{}
	_body, _err := client.ListTablePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenDataWorksStandardServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDataWorksStandardServiceResponse
func (client *Client) OpenDataWorksStandardServiceWithOptions(request *OpenDataWorksStandardServiceRequest, runtime *util.RuntimeOptions) (_result *OpenDataWorksStandardServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDataWorksStandardService"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenDataWorksStandardServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - OpenDataWorksStandardServiceRequest
//
// @return OpenDataWorksStandardServiceResponse
func (client *Client) OpenDataWorksStandardService(request *OpenDataWorksStandardServiceRequest) (_result *OpenDataWorksStandardServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDataWorksStandardServiceResponse{}
	_body, _err := client.OpenDataWorksStandardServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SearchManualDagNodeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchManualDagNodeInstanceResponse
func (client *Client) SearchManualDagNodeInstanceWithOptions(request *SearchManualDagNodeInstanceRequest, runtime *util.RuntimeOptions) (_result *SearchManualDagNodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DagId)) {
		query["DagId"] = request.DagId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchManualDagNodeInstance"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchManualDagNodeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SearchManualDagNodeInstanceRequest
//
// @return SearchManualDagNodeInstanceResponse
func (client *Client) SearchManualDagNodeInstance(request *SearchManualDagNodeInstanceRequest) (_result *SearchManualDagNodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchManualDagNodeInstanceResponse{}
	_body, _err := client.SearchManualDagNodeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SendTaskMetaCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTaskMetaCallbackResponse
func (client *Client) SendTaskMetaCallbackWithOptions(request *SendTaskMetaCallbackRequest, runtime *util.RuntimeOptions) (_result *SendTaskMetaCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionInfo)) {
		body["ConnectionInfo"] = request.ConnectionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		body["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		body["SubType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEnvParam)) {
		body["TaskEnvParam"] = request.TaskEnvParam
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		body["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTaskMetaCallback"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTaskMetaCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SendTaskMetaCallbackRequest
//
// @return SendTaskMetaCallbackResponse
func (client *Client) SendTaskMetaCallback(request *SendTaskMetaCallbackRequest) (_result *SendTaskMetaCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTaskMetaCallbackResponse{}
	_body, _err := client.SendTaskMetaCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Switch的值
//
// @param request - SetSwitchValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSwitchValueResponse
func (client *Client) SetSwitchValueWithOptions(request *SetSwitchValueRequest, runtime *util.RuntimeOptions) (_result *SetSwitchValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SwitchName)) {
		query["SwitchName"] = request.SwitchName
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchValue)) {
		query["SwitchValue"] = request.SwitchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSwitchValue"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSwitchValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Switch的值
//
// @param request - SetSwitchValueRequest
//
// @return SetSwitchValueResponse
func (client *Client) SetSwitchValue(request *SetSwitchValueRequest) (_result *SetSwitchValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSwitchValueResponse{}
	_body, _err := client.SetSwitchValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// startCollect
//
// @param request - StartCollectQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCollectQualityResponse
func (client *Client) StartCollectQualityWithOptions(request *StartCollectQualityRequest, runtime *util.RuntimeOptions) (_result *StartCollectQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackResultString)) {
		body["CallbackResultString"] = request.CallbackResultString
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCollectQuality"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCollectQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// startCollect
//
// @param request - StartCollectQualityRequest
//
// @return StartCollectQualityResponse
func (client *Client) StartCollectQuality(request *StartCollectQualityRequest) (_result *StartCollectQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCollectQualityResponse{}
	_body, _err := client.StartCollectQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartDoCheckQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDoCheckQualityResponse
func (client *Client) StartDoCheckQualityWithOptions(request *StartDoCheckQualityRequest, runtime *util.RuntimeOptions) (_result *StartDoCheckQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackResultString)) {
		body["CallbackResultString"] = request.CallbackResultString
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDoCheckQuality"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDoCheckQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartDoCheckQualityRequest
//
// @return StartDoCheckQualityResponse
func (client *Client) StartDoCheckQuality(request *StartDoCheckQualityRequest) (_result *StartDoCheckQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDoCheckQualityResponse{}
	_body, _err := client.StartDoCheckQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartTaskQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskQualityResponse
func (client *Client) StartTaskQualityWithOptions(request *StartTaskQualityRequest, runtime *util.RuntimeOptions) (_result *StartTaskQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackResultString)) {
		body["CallbackResultString"] = request.CallbackResultString
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTaskQuality"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTaskQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartTaskQualityRequest
//
// @return StartTaskQualityResponse
func (client *Client) StartTaskQuality(request *StartTaskQualityRequest) (_result *StartTaskQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTaskQualityResponse{}
	_body, _err := client.StartTaskQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 触发元数据的Merge操作
//
// @param request - TriggerDataLoaderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerDataLoaderResponse
func (client *Client) TriggerDataLoaderWithOptions(runtime *util.RuntimeOptions) (_result *TriggerDataLoaderResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("TriggerDataLoader"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerDataLoaderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发元数据的Merge操作
//
// @return TriggerDataLoaderResponse
func (client *Client) TriggerDataLoader() (_result *TriggerDataLoaderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerDataLoaderResponse{}
	_body, _err := client.TriggerDataLoaderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 触发timeMachine任务
//
// @param request - TriggerTimeMachineTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerTimeMachineTaskResponse
func (client *Client) TriggerTimeMachineTaskWithOptions(runtime *util.RuntimeOptions) (_result *TriggerTimeMachineTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("TriggerTimeMachineTask"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerTimeMachineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发timeMachine任务
//
// @return TriggerTimeMachineTaskResponse
func (client *Client) TriggerTimeMachineTask() (_result *TriggerTimeMachineTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerTimeMachineTaskResponse{}
	_body, _err := client.TriggerTimeMachineTaskWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
