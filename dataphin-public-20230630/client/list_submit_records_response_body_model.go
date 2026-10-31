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
	// The error code. A value of OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The query result.
	ListResult *ListSubmitRecordsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
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
	// The list of pending publish records.
	Data []*ListSubmitRecordsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total count.
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
	// The change type. Valid values:
	//
	// - 0: create
	//
	// - 1: update
	//
	// - 2: delete
	//
	// example:
	//
	// 0
	ChangeType *int32 `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The time when the record was created. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the record was updated. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The ID of the pending publish record.
	//
	// example:
	//
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// 1234567
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// Object A
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The object type. Valid values:
	//
	// - MAX_COMPUTE_SQL: MaxCompute SQL node
	//
	// - MAX_COMPUTE_MR: MaxCompute MR node
	//
	// - SPARK_JAR_ON_MAX_COMPUTE: Spark JAR on MaxCompute node
	//
	// - SHELL: Shell node
	//
	// - PYTHON: Python node
	//
	// - PERL: Perl script
	//
	// - CHECK: Check node
	//
	// - DATA_X: Data synchronization node
	//
	// - VIRTUAL: Virtual node
	//
	// - IDE_RESOURCE: Resource
	//
	// - UDF: Function
	//
	// - HIVE_SQL: Hive SQL node
	//
	// - HADOOP_MR: Hadoop MR node
	//
	// - SPARK_JAR_ON_HIVE: Spark JAR on Hive node
	//
	// - FLINK_SQL: Flink SQL node
	//
	// - FLINK_TEMPLATE_SQL: Flink SQL template node
	//
	// - STREAM_TEMPLATE: Stream computing template
	//
	// - META_TABLE: Metatable
	//
	// - STREAM_UDF: Stream computing function
	//
	// - FLINK_DATASTREAM: Real-time Flink DataStream node
	//
	// - STREAM_CUSTOM_DATASOURCE: Real-time custom data source
	//
	// - ADB_FOR_PG: AnalyticDB for PostgreSQL node
	//
	// - INCEPTOR_SQL: Inceptor SQL node
	//
	// - MIRROR_TABLE: Mirror table
	//
	// - MIDDLE_TABLE: Intermediate table
	//
	// - APPLICATION_TABLE: Application table
	//
	// - IMPALA_SQL: Impala SQL node
	//
	// - OFFLINE_PIPELINE: Offline pipeline node
	//
	// - REAL_TIME_PIPELINE: Real-time pipeline node
	//
	// - DIM_LOGICAL_TABLE: Logical dimension table
	//
	// - FCT_LOGICAL_TABLE: Logical fact table
	//
	// - BIZ_CONDITION: Business filter
	//
	// - ATOM_INDEX: Atomic metric
	//
	// - DERIVED_INDEX: Derived metric
	//
	// - CALC_DERIVED_INDEX: Derivative metric
	//
	// - PAI_DESIGNER: PAI Designer node
	//
	// - ARGODB_SQL: ArgoDB SQL node
	//
	// - HOLOGRES_SQL: Hologres SQL node
	//
	// - IMPALA_SQL: Impala SQL node
	//
	// - STARROCKS_SQL: StarRocks SQL node
	//
	// - DATABASE_SQL: Database SQL node
	//
	// - SPARK_SQL: Spark SQL node
	//
	// - TASK_TEMPLATE: Task template
	//
	// - EXTERNAL_TRIGGER: External trigger node
	//
	// - GAUSS_SQL: Gauss SQL node
	//
	// example:
	//
	// 2024-10-10 10:00:00
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The version of the object.
	//
	// example:
	//
	// 1
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1241844456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The submission comment.
	//
	// example:
	//
	// Submission comment
	SubmitComment *string `json:"SubmitComment,omitempty" xml:"SubmitComment,omitempty"`
	// The ID of the submitter.
	//
	// example:
	//
	// 307999999
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
	// The name of the submitter.
	//
	// example:
	//
	// Zhang San
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
