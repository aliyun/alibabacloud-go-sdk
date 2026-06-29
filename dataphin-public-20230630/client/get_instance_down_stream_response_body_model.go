// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDownStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceDownStreamResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetInstanceDownStreamResponseBody
	GetHttpStatusCode() *int32
	SetInstanceRelationList(v []*GetInstanceDownStreamResponseBodyInstanceRelationList) *GetInstanceDownStreamResponseBody
	GetInstanceRelationList() []*GetInstanceDownStreamResponseBodyInstanceRelationList
	SetMessage(v string) *GetInstanceDownStreamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceDownStreamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceDownStreamResponseBody
	GetSuccess() *bool
}

type GetInstanceDownStreamResponseBody struct {
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
	// Downstream instances and dependency relationships.
	InstanceRelationList []*GetInstanceDownStreamResponseBodyInstanceRelationList `json:"InstanceRelationList,omitempty" xml:"InstanceRelationList,omitempty" type:"Repeated"`
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceDownStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceDownStreamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceDownStreamResponseBody) GetInstanceRelationList() []*GetInstanceDownStreamResponseBodyInstanceRelationList {
	return s.InstanceRelationList
}

func (s *GetInstanceDownStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceDownStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceDownStreamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceDownStreamResponseBody) SetCode(v string) *GetInstanceDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetHttpStatusCode(v int32) *GetInstanceDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetInstanceRelationList(v []*GetInstanceDownStreamResponseBodyInstanceRelationList) *GetInstanceDownStreamResponseBody {
	s.InstanceRelationList = v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetMessage(v string) *GetInstanceDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetRequestId(v string) *GetInstanceDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetSuccess(v bool) *GetInstanceDownStreamResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) Validate() error {
	if s.InstanceRelationList != nil {
		for _, item := range s.InstanceRelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceDownStreamResponseBodyInstanceRelationList struct {
	// Depth level.
	//
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// Extended information of the node instance.
	//
	// example:
	//
	// {"a":"x"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Field instance list.
	FieldInstanceList []*GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList `json:"FieldInstanceList,omitempty" xml:"FieldInstanceList,omitempty" type:"Repeated"`
	// Instance information.
	InstanceInfo *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	// Run status of the instance.
	//
	// example:
	//
	// RUNNING
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	// Selection status of the instance, primarily used for logical fields.
	//
	// - OPTIONAL: Optional.
	//
	// - SELECTED: Required.
	//
	// - DISABLE: Not selectable.
	//
	// - GROUP_SELECTED: Select all or deselect all.
	//
	// example:
	//
	// OPTIONAL
	SelectStatus *string `json:"SelectStatus,omitempty" xml:"SelectStatus,omitempty"`
	// Reason why the instance is selected.
	//
	// - PK_CHANGE
	//
	// - LOGIC_CHANGE
	//
	// - RELATION
	//
	// - BIZ_RELATIO
	//
	// - MV_RELATION
	//
	// - MODEL_REFRESH
	//
	// - FIELD_DELETED
	//
	// - FIELD_ADDED_NO_INSTANCE
	//
	// - PERMISSION_DENY
	//
	// - OUTSIDE_INCOMING
	//
	// - INSTANCE_STATUS
	//
	// - AFFECT_BY_UPSTREAM
	//
	// example:
	//
	// FIELD_DELETED
	SelectStatusCause *string `json:"SelectStatusCause,omitempty" xml:"SelectStatusCause,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationList) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetFieldInstanceList() []*GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	return s.FieldInstanceList
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetInstanceInfo() *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	return s.InstanceInfo
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetRunStatus() *string {
	return s.RunStatus
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetSelectStatus() *string {
	return s.SelectStatus
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) GetSelectStatusCause() *string {
	return s.SelectStatusCause
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetDownStreamDepth(v int32) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetExtendInfo(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.ExtendInfo = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetFieldInstanceList(v []*GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.FieldInstanceList = v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetInstanceInfo(v *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.InstanceInfo = v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetRunStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetSelectStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.SelectStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetSelectStatusCause(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.SelectStatusCause = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) Validate() error {
	if s.FieldInstanceList != nil {
		for _, item := range s.FieldInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstanceInfo != nil {
		if err := s.InstanceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList struct {
	// Field instance ID.
	//
	// example:
	//
	// t_23211
	FieldInstanceId *string `json:"FieldInstanceId,omitempty" xml:"FieldInstanceId,omitempty"`
	// Run status.
	//
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	// Selection status of the instance.
	//
	// example:
	//
	// OPTIONAL
	SelectStatus *string `json:"SelectStatus,omitempty" xml:"SelectStatus,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) GetFieldInstanceId() *string {
	return s.FieldInstanceId
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) GetRunStatus() *string {
	return s.RunStatus
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) GetSelectStatus() *string {
	return s.SelectStatus
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetFieldInstanceId(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.FieldInstanceId = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetRunStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetSelectStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.SelectStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo struct {
	// Instance ID.
	//
	// example:
	//
	// t_232411
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Instance name.
	//
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Node type.
	//
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) GetId() *string {
	return s.Id
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) GetName() *string {
	return s.Name
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) GetType() *string {
	return s.Type
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetId(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Id = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetName(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Name = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetType(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Type = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) Validate() error {
	return dara.Validate(s)
}
