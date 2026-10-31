// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeByOutputNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalNodeByOutputNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPhysicalNodeByOutputNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPhysicalNodeByOutputNameResponseBody
	GetMessage() *string
	SetNodeInfo(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) *GetPhysicalNodeByOutputNameResponseBody
	GetNodeInfo() *GetPhysicalNodeByOutputNameResponseBodyNodeInfo
	SetRequestId(v string) *GetPhysicalNodeByOutputNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalNodeByOutputNameResponseBody
	GetSuccess() *bool
}

type GetPhysicalNodeByOutputNameResponseBody struct {
	// The error code. A value of OK indicates that the request is normal.
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
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The node details.
	NodeInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetNodeInfo() *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	return s.NodeInfo
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalNodeByOutputNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetCode(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeByOutputNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetMessage(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetNodeInfo(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) *GetPhysicalNodeByOutputNameResponseBody {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetRequestId(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetSuccess(v bool) *GetPhysicalNodeByOutputNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfo struct {
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1717343597000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the node.
	Creator *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// xx test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The node source. Valid values:
	//
	// - DATA_PROCESS: code development
	//
	// - BLACK_BOX: black box
	//
	// - ONE_ID: extraction
	//
	// - PIPELINE: pipeline
	//
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_2321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1717343597000
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The modifier of the node.
	Modifier *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// The name.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node subtype. Valid values:
	//
	// - MAX_COMPUTE_SQL
	//
	// - HIVE_SQL
	//
	// - SHELL
	//
	// - PYTHON
	//
	// - ONE_SERVICE_SQL
	//
	// - DATABASE_SQL
	//
	// example:
	//
	// SHELL
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The owner of the node.
	Owner *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The priority. Valid values:
	//
	// - HIGHEST
	//
	// - HIGH
	//
	// - MIDDLE
	//
	// - LOW
	//
	// - LOWEST
	//
	// example:
	//
	// MIDDLE
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project to which the node belongs.
	ProjectInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// The scheduling cycle type. Valid values:
	//
	// - MINUTELY
	//
	// - HOURLY
	//
	// - DAILY
	//
	// - WEEKLY
	//
	// - MONTHLY
	//
	// - YEARLY
	//
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The scheduling status of the node.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trigger configuration, which is used to implement field dependencies for logical tables.
	//
	// example:
	//
	// {"expression":"any_success"}
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetCreator() *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator {
	return s.Creator
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetFrom() *string {
	return s.From
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetModifier() *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier {
	return s.Modifier
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetOperatorType() *string {
	return s.OperatorType
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetOwner() *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner {
	return s.Owner
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetPriority() *string {
	return s.Priority
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetProjectInfo() *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo {
	return s.ProjectInfo
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetCreateTime(v int64) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetCreator(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetDescription(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetFrom(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetLastModifiedTime(v int64) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetModifier(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetOperatorType(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.OperatorType = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetOwner(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Owner = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetPriority(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Priority = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetProjectInfo(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.ProjectInfo = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetScheduleType(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetStatus(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Status = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetTriggerConfig(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.TriggerConfig = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.ProjectInfo != nil {
		if err := s.ProjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator struct {
	// The user ID.
	//
	// example:
	//
	// 1311131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier struct {
	// The user ID.
	//
	// example:
	//
	// 1311131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner struct {
	// The user ID.
	//
	// example:
	//
	// 1311131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo struct {
	// The project ID.
	//
	// example:
	//
	// 1324211
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project name.
	//
	// example:
	//
	// xx test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) Validate() error {
	return dara.Validate(s)
}
