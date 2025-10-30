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
	Message  *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1717343597000
	CreateTime  *int64                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// n_2321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717343597000
	LastModifiedTime *int64                                                   `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SHELL
	OperatorType *string                                               `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	Owner        *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// MIDDLE
	Priority    *string                                                     `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
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
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// example:
	//
	// 1324211
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
