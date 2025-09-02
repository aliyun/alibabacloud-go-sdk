// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityFollowerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetQualityFollowerResponseBodyData) *GetQualityFollowerResponseBody
	GetData() []*GetQualityFollowerResponseBodyData
	SetErrorCode(v string) *GetQualityFollowerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetQualityFollowerResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetQualityFollowerResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetQualityFollowerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityFollowerResponseBody
	GetSuccess() *bool
}

type GetQualityFollowerResponseBody struct {
	// The information about the subscription relationship.
	Data []*GetQualityFollowerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You have no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP return code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 38cbdef0-f6cf-49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityFollowerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityFollowerResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityFollowerResponseBody) GetData() []*GetQualityFollowerResponseBodyData {
	return s.Data
}

func (s *GetQualityFollowerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQualityFollowerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetQualityFollowerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityFollowerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityFollowerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityFollowerResponseBody) SetData(v []*GetQualityFollowerResponseBodyData) *GetQualityFollowerResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityFollowerResponseBody) SetErrorCode(v string) *GetQualityFollowerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQualityFollowerResponseBody) SetErrorMessage(v string) *GetQualityFollowerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetQualityFollowerResponseBody) SetHttpStatusCode(v int32) *GetQualityFollowerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityFollowerResponseBody) SetRequestId(v string) *GetQualityFollowerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityFollowerResponseBody) SetSuccess(v bool) *GetQualityFollowerResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityFollowerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQualityFollowerResponseBodyData struct {
	// The alert mode. The value is as follows:
	//
	// - 1 (Mail)
	//
	// - 2 (email and SMS)
	//
	// - 4 (DingTalk groups of robots or hook)
	//
	// - 5 (DingTalk groups of robots @ ALL)
	//
	// example:
	//
	// 1
	AlarmMode *int32 `json:"AlarmMode,omitempty" xml:"AlarmMode,omitempty"`
	// The time when the data quality rule subscription configuration was created.
	//
	// example:
	//
	// 1541576644000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the partition expression.
	//
	// example:
	//
	// 1234
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The subscriber to receive alert information.
	//
	// example:
	//
	// 1234
	Follower *string `json:"Follower,omitempty" xml:"Follower,omitempty"`
	// The Alibaba Cloud account name of the subscriber.
	//
	// example:
	//
	// test
	FollowerAccountName *string `json:"FollowerAccountName,omitempty" xml:"FollowerAccountName,omitempty"`
	// The ID of the subscription relationship.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The update time of the data quality rule subscription configuration.
	//
	// example:
	//
	// 1541576644000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the engine or data source.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the partitioned table.
	//
	// example:
	//
	// dual
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetQualityFollowerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityFollowerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityFollowerResponseBodyData) GetAlarmMode() *int32 {
	return s.AlarmMode
}

func (s *GetQualityFollowerResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetQualityFollowerResponseBodyData) GetEntityId() *string {
	return s.EntityId
}

func (s *GetQualityFollowerResponseBodyData) GetFollower() *string {
	return s.Follower
}

func (s *GetQualityFollowerResponseBodyData) GetFollowerAccountName() *string {
	return s.FollowerAccountName
}

func (s *GetQualityFollowerResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetQualityFollowerResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetQualityFollowerResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityFollowerResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetQualityFollowerResponseBodyData) SetAlarmMode(v int32) *GetQualityFollowerResponseBodyData {
	s.AlarmMode = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetCreateTime(v int64) *GetQualityFollowerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetEntityId(v string) *GetQualityFollowerResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetFollower(v string) *GetQualityFollowerResponseBodyData {
	s.Follower = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetFollowerAccountName(v string) *GetQualityFollowerResponseBodyData {
	s.FollowerAccountName = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetId(v int64) *GetQualityFollowerResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetModifyTime(v int64) *GetQualityFollowerResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetProjectName(v string) *GetQualityFollowerResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) SetTableName(v string) *GetQualityFollowerResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetQualityFollowerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
