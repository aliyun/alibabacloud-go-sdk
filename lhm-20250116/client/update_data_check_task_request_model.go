// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckTemplateId(v string) *UpdateDataCheckTaskRequest
	GetCheckTemplateId() *string
	SetDstDsId(v string) *UpdateDataCheckTaskRequest
	GetDstDsId() *string
	SetDstDsName(v string) *UpdateDataCheckTaskRequest
	GetDstDsName() *string
	SetDstDsType(v string) *UpdateDataCheckTaskRequest
	GetDstDsType() *string
	SetDstEngineId(v string) *UpdateDataCheckTaskRequest
	GetDstEngineId() *string
	SetDstEngineName(v string) *UpdateDataCheckTaskRequest
	GetDstEngineName() *string
	SetDstEngineType(v string) *UpdateDataCheckTaskRequest
	GetDstEngineType() *string
	SetId(v int64) *UpdateDataCheckTaskRequest
	GetId() *int64
	SetSrcDsId(v string) *UpdateDataCheckTaskRequest
	GetSrcDsId() *string
	SetSrcDsName(v string) *UpdateDataCheckTaskRequest
	GetSrcDsName() *string
	SetSrcDsType(v string) *UpdateDataCheckTaskRequest
	GetSrcDsType() *string
	SetSrcEngineId(v string) *UpdateDataCheckTaskRequest
	GetSrcEngineId() *string
	SetSrcEngineName(v string) *UpdateDataCheckTaskRequest
	GetSrcEngineName() *string
	SetSrcEngineType(v string) *UpdateDataCheckTaskRequest
	GetSrcEngineType() *string
	SetTaskDescription(v string) *UpdateDataCheckTaskRequest
	GetTaskDescription() *string
	SetTaskName(v string) *UpdateDataCheckTaskRequest
	GetTaskName() *string
}

type UpdateDataCheckTaskRequest struct {
	// The ID of the validation template. If this field is not specified, the original value is retained.
	//
	// example:
	//
	// 1001
	CheckTemplateId *string `json:"checkTemplateId,omitempty" xml:"checkTemplateId,omitempty"`
	// The ID of the destination data source.
	//
	// example:
	//
	// 2001
	DstDsId *string `json:"dstDsId,omitempty" xml:"dstDsId,omitempty"`
	// The name of the destination data source.
	//
	// example:
	//
	// ds_demo
	DstDsName *string `json:"dstDsName,omitempty" xml:"dstDsName,omitempty"`
	// The type of the destination data source.
	//
	// example:
	//
	// Hive
	DstDsType *string `json:"dstDsType,omitempty" xml:"dstDsType,omitempty"`
	// The ID of the destination validation engine.
	//
	// example:
	//
	// 2001
	DstEngineId *string `json:"dstEngineId,omitempty" xml:"dstEngineId,omitempty"`
	// The name of the destination validation engine.
	//
	// example:
	//
	// engine_demo
	DstEngineName *string `json:"dstEngineName,omitempty" xml:"dstEngineName,omitempty"`
	// The type of the destination validation engine.
	//
	// example:
	//
	// Tez
	DstEngineType *string `json:"dstEngineType,omitempty" xml:"dstEngineType,omitempty"`
	// The ID of the task to modify. This field is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the source data source.
	//
	// example:
	//
	// 1001
	SrcDsId *string `json:"srcDsId,omitempty" xml:"srcDsId,omitempty"`
	// The name of the source data source.
	//
	// example:
	//
	// ds_demo
	SrcDsName *string `json:"srcDsName,omitempty" xml:"srcDsName,omitempty"`
	// The type of the source data source.
	//
	// example:
	//
	// Hive
	SrcDsType *string `json:"srcDsType,omitempty" xml:"srcDsType,omitempty"`
	// The ID of the source validation engine.
	//
	// example:
	//
	// 1001
	SrcEngineId *string `json:"srcEngineId,omitempty" xml:"srcEngineId,omitempty"`
	// The name of the source validation engine.
	//
	// example:
	//
	// engine_demo
	SrcEngineName *string `json:"srcEngineName,omitempty" xml:"srcEngineName,omitempty"`
	// The type of the source validation engine.
	//
	// example:
	//
	// Tez
	SrcEngineType *string `json:"srcEngineType,omitempty" xml:"srcEngineType,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// Data validation task description
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// The name of the task. Only Chinese characters, English letters, and digits are supported.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s UpdateDataCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTaskRequest) GetCheckTemplateId() *string {
	return s.CheckTemplateId
}

func (s *UpdateDataCheckTaskRequest) GetDstDsId() *string {
	return s.DstDsId
}

func (s *UpdateDataCheckTaskRequest) GetDstDsName() *string {
	return s.DstDsName
}

func (s *UpdateDataCheckTaskRequest) GetDstDsType() *string {
	return s.DstDsType
}

func (s *UpdateDataCheckTaskRequest) GetDstEngineId() *string {
	return s.DstEngineId
}

func (s *UpdateDataCheckTaskRequest) GetDstEngineName() *string {
	return s.DstEngineName
}

func (s *UpdateDataCheckTaskRequest) GetDstEngineType() *string {
	return s.DstEngineType
}

func (s *UpdateDataCheckTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataCheckTaskRequest) GetSrcDsId() *string {
	return s.SrcDsId
}

func (s *UpdateDataCheckTaskRequest) GetSrcDsName() *string {
	return s.SrcDsName
}

func (s *UpdateDataCheckTaskRequest) GetSrcDsType() *string {
	return s.SrcDsType
}

func (s *UpdateDataCheckTaskRequest) GetSrcEngineId() *string {
	return s.SrcEngineId
}

func (s *UpdateDataCheckTaskRequest) GetSrcEngineName() *string {
	return s.SrcEngineName
}

func (s *UpdateDataCheckTaskRequest) GetSrcEngineType() *string {
	return s.SrcEngineType
}

func (s *UpdateDataCheckTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *UpdateDataCheckTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateDataCheckTaskRequest) SetCheckTemplateId(v string) *UpdateDataCheckTaskRequest {
	s.CheckTemplateId = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstDsId(v string) *UpdateDataCheckTaskRequest {
	s.DstDsId = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstDsName(v string) *UpdateDataCheckTaskRequest {
	s.DstDsName = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstDsType(v string) *UpdateDataCheckTaskRequest {
	s.DstDsType = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstEngineId(v string) *UpdateDataCheckTaskRequest {
	s.DstEngineId = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstEngineName(v string) *UpdateDataCheckTaskRequest {
	s.DstEngineName = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetDstEngineType(v string) *UpdateDataCheckTaskRequest {
	s.DstEngineType = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetId(v int64) *UpdateDataCheckTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcDsId(v string) *UpdateDataCheckTaskRequest {
	s.SrcDsId = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcDsName(v string) *UpdateDataCheckTaskRequest {
	s.SrcDsName = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcDsType(v string) *UpdateDataCheckTaskRequest {
	s.SrcDsType = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcEngineId(v string) *UpdateDataCheckTaskRequest {
	s.SrcEngineId = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcEngineName(v string) *UpdateDataCheckTaskRequest {
	s.SrcEngineName = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetSrcEngineType(v string) *UpdateDataCheckTaskRequest {
	s.SrcEngineType = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetTaskDescription(v string) *UpdateDataCheckTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) SetTaskName(v string) *UpdateDataCheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateDataCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
