// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetQualityEntityResponseBodyData) *GetQualityEntityResponseBody
	GetData() []*GetQualityEntityResponseBodyData
	SetErrorCode(v string) *GetQualityEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetQualityEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetQualityEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetQualityEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityEntityResponseBody
	GetSuccess() *bool
}

type GetQualityEntityResponseBody struct {
	// The information about the partition filter expression.
	Data []*GetQualityEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned.
	//
	// example:
	//
	// 401
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6d739ef6-098a-47****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityEntityResponseBody) GetData() []*GetQualityEntityResponseBodyData {
	return s.Data
}

func (s *GetQualityEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQualityEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetQualityEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityEntityResponseBody) SetData(v []*GetQualityEntityResponseBodyData) *GetQualityEntityResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityEntityResponseBody) SetErrorCode(v string) *GetQualityEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQualityEntityResponseBody) SetErrorMessage(v string) *GetQualityEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetQualityEntityResponseBody) SetHttpStatusCode(v int32) *GetQualityEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityEntityResponseBody) SetRequestId(v string) *GetQualityEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityEntityResponseBody) SetSuccess(v bool) *GetQualityEntityResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQualityEntityResponseBodyData struct {
	// The time when the partition filter expression was created.
	//
	// example:
	//
	// 1593964800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The level of the partition filter expression. Valid values:
	//
	// 	- 0: The partition filter expression is at the SQL level. This indicates that the system checks data quality after each SQL statement is executed.
	//
	// 	- 1: The partition filter expression is at the node level. This indicates that the system checks data quality after all the SQL statements for a node are executed.
	//
	// example:
	//
	// 0
	EntityLevel *int32 `json:"EntityLevel,omitempty" xml:"EntityLevel,omitempty"`
	// The type of the compute engine instance or data source.
	//
	// example:
	//
	// odps
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the Alibaba Cloud account that is used to receive alert notifications.
	//
	// example:
	//
	// 1822931****
	Followers *string `json:"Followers,omitempty" xml:"Followers,omitempty"`
	// Indicates whether the partition filter expression is associated with a node. Valid values:
	//
	// 	- true: The partition filter expression is associated with a node.
	//
	// 	- false: The partition filter expression is not associated with a node.
	//
	// example:
	//
	// true
	HasRelativeNode *bool `json:"HasRelativeNode,omitempty" xml:"HasRelativeNode,omitempty"`
	// The ID of the partition filter expression.
	//
	// example:
	//
	// 4003918
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The partition filter expression.
	//
	// example:
	//
	// dt=$[yyyymmdd-1]
	MatchExpression *string `json:"MatchExpression,omitempty" xml:"MatchExpression,omitempty"`
	// The time when the partition filter expression was modified.
	//
	// example:
	//
	// 1593964800000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to modify the partition filter expression.
	//
	// example:
	//
	// 1822931****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the Alibaba Cloud account that is used to configure the partition filter expression.
	//
	// example:
	//
	// 1822931****
	OnDuty *string `json:"OnDuty,omitempty" xml:"OnDuty,omitempty"`
	// The name of the Alibaba Cloud account that is used to configure the partition filter expression.
	//
	// example:
	//
	// test
	OnDutyAccountName *string `json:"OnDutyAccountName,omitempty" xml:"OnDutyAccountName,omitempty"`
	// The name of the compute engine instance or data source.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The information about the node with which the partition filter expression is associated. The information includes the following items:
	//
	// 	- ProjectName: the name of the workspace to which the node belongs.
	//
	// 	- NodeID: the ID of the node.
	//
	// example:
	//
	// [{"projectName":"xc_DP****","nodeId":7000026****}]
	RelativeNode *string `json:"RelativeNode,omitempty" xml:"RelativeNode,omitempty"`
	// Indicates that the partition filter expression is at the SQL level.
	//
	// example:
	//
	// 0
	Sql *int32 `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The name of the partitioned table.
	//
	// example:
	//
	// test_dqc_de****
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The node.
	//
	// example:
	//
	// 0
	Task *int32 `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s GetQualityEntityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityEntityResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQualityEntityResponseBodyData) GetEntityLevel() *int32 {
	return s.EntityLevel
}

func (s *GetQualityEntityResponseBodyData) GetEnvType() *string {
	return s.EnvType
}

func (s *GetQualityEntityResponseBodyData) GetFollowers() *string {
	return s.Followers
}

func (s *GetQualityEntityResponseBodyData) GetHasRelativeNode() *bool {
	return s.HasRelativeNode
}

func (s *GetQualityEntityResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetQualityEntityResponseBodyData) GetMatchExpression() *string {
	return s.MatchExpression
}

func (s *GetQualityEntityResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetQualityEntityResponseBodyData) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetQualityEntityResponseBodyData) GetOnDuty() *string {
	return s.OnDuty
}

func (s *GetQualityEntityResponseBodyData) GetOnDutyAccountName() *string {
	return s.OnDutyAccountName
}

func (s *GetQualityEntityResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityEntityResponseBodyData) GetRelativeNode() *string {
	return s.RelativeNode
}

func (s *GetQualityEntityResponseBodyData) GetSql() *int32 {
	return s.Sql
}

func (s *GetQualityEntityResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetQualityEntityResponseBodyData) GetTask() *int32 {
	return s.Task
}

func (s *GetQualityEntityResponseBodyData) SetCreateTime(v int64) *GetQualityEntityResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetEntityLevel(v int32) *GetQualityEntityResponseBodyData {
	s.EntityLevel = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetEnvType(v string) *GetQualityEntityResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetFollowers(v string) *GetQualityEntityResponseBodyData {
	s.Followers = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetHasRelativeNode(v bool) *GetQualityEntityResponseBodyData {
	s.HasRelativeNode = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetId(v int64) *GetQualityEntityResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetMatchExpression(v string) *GetQualityEntityResponseBodyData {
	s.MatchExpression = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetModifyTime(v int64) *GetQualityEntityResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetModifyUser(v string) *GetQualityEntityResponseBodyData {
	s.ModifyUser = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetOnDuty(v string) *GetQualityEntityResponseBodyData {
	s.OnDuty = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetOnDutyAccountName(v string) *GetQualityEntityResponseBodyData {
	s.OnDutyAccountName = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetProjectName(v string) *GetQualityEntityResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetRelativeNode(v string) *GetQualityEntityResponseBodyData {
	s.RelativeNode = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetSql(v int32) *GetQualityEntityResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetTableName(v string) *GetQualityEntityResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) SetTask(v int32) *GetQualityEntityResponseBodyData {
	s.Task = &v
	return s
}

func (s *GetQualityEntityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
