// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPhysicalNodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPhysicalNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPhysicalNodeResponseBody
	GetMessage() *string
	SetNodeInfo(v *GetPhysicalNodeResponseBodyNodeInfo) *GetPhysicalNodeResponseBody
	GetNodeInfo() *GetPhysicalNodeResponseBodyNodeInfo
	SetRequestId(v string) *GetPhysicalNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPhysicalNodeResponseBody
	GetSuccess() *bool
}

type GetPhysicalNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfo *GetPhysicalNodeResponseBodyNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhysicalNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPhysicalNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhysicalNodeResponseBody) GetNodeInfo() *GetPhysicalNodeResponseBodyNodeInfo {
	return s.NodeInfo
}

func (s *GetPhysicalNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhysicalNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPhysicalNodeResponseBody) SetCode(v string) *GetPhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetMessage(v string) *GetPhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetNodeInfo(v *GetPhysicalNodeResponseBodyNodeInfo) *GetPhysicalNodeResponseBody {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetRequestId(v string) *GetPhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetSuccess(v bool) *GetPhysicalNodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhysicalNodeResponseBodyNodeInfo struct {
	// example:
	//
	// 1717343597000
	CreateTime *int64                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetPhysicalNodeResponseBodyNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// 0 0 10 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 123456789
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// null
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// n_232132
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717343597000
	LastModifiedTime *int64                                       `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetPhysicalNodeResponseBodyNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SHELL
	OperatorType   *string                                   `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	OutputNameList []*string                                 `json:"OutputNameList,omitempty" xml:"OutputNameList,omitempty" type:"Repeated"`
	Owner          *GetPhysicalNodeResponseBodyNodeInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// MIDDLE
	Priority    *string                                         `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectInfo *GetPhysicalNodeResponseBodyNodeInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"expression":"any_success"}
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetCreator() *GetPhysicalNodeResponseBodyNodeInfoCreator {
	return s.Creator
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetFrom() *string {
	return s.From
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetModifier() *GetPhysicalNodeResponseBodyNodeInfoModifier {
	return s.Modifier
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetOperatorType() *string {
	return s.OperatorType
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetOutputNameList() []*string {
	return s.OutputNameList
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetOwner() *GetPhysicalNodeResponseBodyNodeInfoOwner {
	return s.Owner
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetPriority() *string {
	return s.Priority
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetProjectInfo() *GetPhysicalNodeResponseBodyNodeInfoProjectInfo {
	return s.ProjectInfo
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCreateTime(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCreator(v *GetPhysicalNodeResponseBodyNodeInfoCreator) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCronExpression(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.CronExpression = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDataSourceId(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDataSourceSchema(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.DataSourceSchema = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDescription(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetFrom(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetLastModifiedTime(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetModifier(v *GetPhysicalNodeResponseBodyNodeInfoModifier) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOperatorType(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.OperatorType = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOutputNameList(v []*string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.OutputNameList = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOwner(v *GetPhysicalNodeResponseBodyNodeInfoOwner) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Owner = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetPriority(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Priority = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetProjectInfo(v *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) *GetPhysicalNodeResponseBodyNodeInfo {
	s.ProjectInfo = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetScheduleType(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetStatus(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Status = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetTriggerConfig(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.TriggerConfig = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) Validate() error {
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

type GetPhysicalNodeResponseBodyNodeInfoCreator struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoCreator {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeResponseBodyNodeInfoModifier struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoModifier) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoModifier {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeResponseBodyNodeInfoOwner struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoOwner) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoOwner) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoOwner {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoOwner {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) Validate() error {
	return dara.Validate(s)
}

type GetPhysicalNodeResponseBodyNodeInfoProjectInfo struct {
	// example:
	//
	// 102132
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) GetId() *string {
	return s.Id
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) GetName() *string {
	return s.Name
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoProjectInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) Validate() error {
	return dara.Validate(s)
}
