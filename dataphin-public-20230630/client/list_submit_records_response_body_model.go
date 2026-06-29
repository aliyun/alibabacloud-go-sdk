// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubmitRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSubmitRecordsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSubmitRecordsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListSubmitRecordsResponseBodyListResult) *ListSubmitRecordsResponseBody
	GetListResult() *ListSubmitRecordsResponseBodyListResult
	SetMessage(v string) *ListSubmitRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubmitRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSubmitRecordsResponseBody
	GetSuccess() *bool
}

type ListSubmitRecordsResponseBody struct {
	// Error code. OK indicates a successful request.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Query result.
	ListResult *ListSubmitRecordsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// Error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSubmitRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSubmitRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSubmitRecordsResponseBody) GetListResult() *ListSubmitRecordsResponseBodyListResult {
	return s.ListResult
}

func (s *ListSubmitRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubmitRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubmitRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubmitRecordsResponseBody) SetCode(v string) *ListSubmitRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetHttpStatusCode(v int32) *ListSubmitRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetListResult(v *ListSubmitRecordsResponseBodyListResult) *ListSubmitRecordsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetMessage(v string) *ListSubmitRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetRequestId(v string) *ListSubmitRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) SetSuccess(v bool) *ListSubmitRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubmitRecordsResponseBody) Validate() error {
	if s.ListResult != nil {
		if err := s.ListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubmitRecordsResponseBodyListResult struct {
	// List of pending deployment records.
	Data []*ListSubmitRecordsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Total count.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubmitRecordsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBodyListResult) GetData() []*ListSubmitRecordsResponseBodyListResultData {
	return s.Data
}

func (s *ListSubmitRecordsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubmitRecordsResponseBodyListResult) SetData(v []*ListSubmitRecordsResponseBodyListResultData) *ListSubmitRecordsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResult) SetTotalCount(v int32) *ListSubmitRecordsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResult) Validate() error {
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

type ListSubmitRecordsResponseBodyListResultData struct {
	// Change type. 0: Create / 1: Update / 2: Delete.
	//
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// Creation time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Modification time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// Pending deployment record ID.
	//
	// example:
	//
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Node ID.
	//
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Object ID.
	//
	// example:
	//
	// 1234567
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// Object name.
	//
	// example:
	//
	// 对象A
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// Object type. Valid values:
	//
	// - MaxCompute SQL task: MAX_COMPUTE_SQL
	//
	// - MaxCompute MR task: MAX_COMPUTE_MR
	//
	// - Spark JAR on MaxCompute: SPARK_JAR_ON_MAX_COMPUTE
	//
	// - Shell task: SHELL
	//
	// - Python task: PYTHON
	//
	// - Perl script: PERL
	//
	// - Check: CHECK
	//
	// - Sync task: DATA_X
	//
	// - Virtual node: VIRTUAL
	//
	// - Resource: IDE_RESOURCE
	//
	// - Function: UDF
	//
	// - Hive SQL task: HIVE_SQL
	//
	// - Hadoop MR task: HADOOP_MR
	//
	// - Spark JAR on Hive task: SPARK_JAR_ON_HIVE
	//
	// - Flink SQL task: FLINK_SQL
	//
	// - Flink SQL template task: FLINK_TEMPLATE_SQL
	//
	// - Stream computing template: STREAM_TEMPLATE
	//
	// - Metatable: META_TABLE
	//
	// - Stream computing function: STREAM_UDF
	//
	// - Real-time Flink DataStream: FLINK_DATASTREAM
	//
	// - Real-time custom data source: STREAM_CUSTOM_DATASOURCE
	//
	// - AnalyticDB for PostgreSQL task: ADB_FOR_PG
	//
	// - TDH SQL task: INCEPTOR_SQL
	//
	// - Mirror table: MIRROR_TABLE
	//
	// - Intermediate table: MIDDLE_TABLE
	//
	// - Application table: APPLICATION_TABLE
	//
	// - Impala SQL task: IMPALA_SQL
	//
	// - Offline pipeline task: OFFLINE_PIPELINE
	//
	// - Real-time pipeline task: REAL_TIME_PIPELINE
	//
	// - Dimension logical table: DIM_LOGICAL_TABLE
	//
	// - Fact logical table: FCT_LOGICAL_TABLE
	//
	// - Business condition: BIZ_CONDITION
	//
	// - Atomic metric: ATOM_INDEX
	//
	// - Derived metric: DERIVED_INDEX
	//
	// - Calculated derived metric: CALC_DERIVED_INDEX
	//
	// - PAI task: PAI_DESIGNER
	//
	// - ArgoDB SQL task: ARGODB_SQL
	//
	// - Hologres SQL task: HOLOGRES_SQL
	//
	// - Impala SQL task: IMPALA_SQL
	//
	// - StarRocks SQL task: STARROCKS_SQL
	//
	// - Database SQL task: DATABASE_SQL
	//
	// - Spark SQL task: SPARK_SQL
	//
	// - Compute template: TASK_TEMPLATE
	//
	// - External trigger node: EXTERNAL_TRIGGER
	//
	// - Gauss SQL task: GAUSS_SQL
	//
	// example:
	//
	// 2024-10-10 10:00:00
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// Object version.
	//
	// example:
	//
	// 1
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// Project ID.
	//
	// example:
	//
	// 1241844456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Submission comment.
	//
	// example:
	//
	// 提交信息
	SubmitComment *string `json:"SubmitComment,omitempty" xml:"SubmitComment,omitempty"`
	// Submitter ID.
	//
	// example:
	//
	// 307999999
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
	// Submitter name.
	//
	// example:
	//
	// 张三
	SubmitterName *string `json:"SubmitterName,omitempty" xml:"SubmitterName,omitempty"`
}

func (s ListSubmitRecordsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetId() *int64 {
	return s.Id
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitComment() *string {
	return s.SubmitComment
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitter() *string {
	return s.Submitter
}

func (s *ListSubmitRecordsResponseBodyListResultData) GetSubmitterName() *string {
	return s.SubmitterName
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetChangeType(v int32) *ListSubmitRecordsResponseBodyListResultData {
	s.ChangeType = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetGmtCreate(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetGmtModify(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.GmtModify = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetId(v int64) *ListSubmitRecordsResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetNodeId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.NodeId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectName(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectName = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectType(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectType = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetObjectVersion(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ObjectVersion = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetProjectId(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.ProjectId = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitComment(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.SubmitComment = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitter(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.Submitter = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) SetSubmitterName(v string) *ListSubmitRecordsResponseBodyListResultData {
	s.SubmitterName = &v
	return s
}

func (s *ListSubmitRecordsResponseBodyListResultData) Validate() error {
	return dara.Validate(s)
}
