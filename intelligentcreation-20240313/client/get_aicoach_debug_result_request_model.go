// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachDebugResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *GetAICoachDebugResultRequest
	GetDataId() *string
	SetDataType(v string) *GetAICoachDebugResultRequest
	GetDataType() *string
	SetScriptDebugId(v string) *GetAICoachDebugResultRequest
	GetScriptDebugId() *string
	SetScriptRecordId(v string) *GetAICoachDebugResultRequest
	GetScriptRecordId() *string
	SetScriptSnapshotId(v string) *GetAICoachDebugResultRequest
	GetScriptSnapshotId() *string
	SetTaskId(v string) *GetAICoachDebugResultRequest
	GetTaskId() *string
}

type GetAICoachDebugResultRequest struct {
	DataId           *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	DataType         *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	ScriptDebugId    *string `json:"scriptDebugId,omitempty" xml:"scriptDebugId,omitempty"`
	ScriptRecordId   *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
	ScriptSnapshotId *string `json:"scriptSnapshotId,omitempty" xml:"scriptSnapshotId,omitempty"`
	TaskId           *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAICoachDebugResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultRequest) GetDataId() *string {
	return s.DataId
}

func (s *GetAICoachDebugResultRequest) GetDataType() *string {
	return s.DataType
}

func (s *GetAICoachDebugResultRequest) GetScriptDebugId() *string {
	return s.ScriptDebugId
}

func (s *GetAICoachDebugResultRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *GetAICoachDebugResultRequest) GetScriptSnapshotId() *string {
	return s.ScriptSnapshotId
}

func (s *GetAICoachDebugResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAICoachDebugResultRequest) SetDataId(v string) *GetAICoachDebugResultRequest {
	s.DataId = &v
	return s
}

func (s *GetAICoachDebugResultRequest) SetDataType(v string) *GetAICoachDebugResultRequest {
	s.DataType = &v
	return s
}

func (s *GetAICoachDebugResultRequest) SetScriptDebugId(v string) *GetAICoachDebugResultRequest {
	s.ScriptDebugId = &v
	return s
}

func (s *GetAICoachDebugResultRequest) SetScriptRecordId(v string) *GetAICoachDebugResultRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *GetAICoachDebugResultRequest) SetScriptSnapshotId(v string) *GetAICoachDebugResultRequest {
	s.ScriptSnapshotId = &v
	return s
}

func (s *GetAICoachDebugResultRequest) SetTaskId(v string) *GetAICoachDebugResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetAICoachDebugResultRequest) Validate() error {
	return dara.Validate(s)
}
