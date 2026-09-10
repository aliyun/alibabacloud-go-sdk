// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckTemplateId(v string) *AddDataCheckTaskRequest
	GetCheckTemplateId() *string
	SetCheckType(v int32) *AddDataCheckTaskRequest
	GetCheckType() *int32
	SetDstDsId(v string) *AddDataCheckTaskRequest
	GetDstDsId() *string
	SetDstDsName(v string) *AddDataCheckTaskRequest
	GetDstDsName() *string
	SetDstDsType(v string) *AddDataCheckTaskRequest
	GetDstDsType() *string
	SetSrcDsId(v string) *AddDataCheckTaskRequest
	GetSrcDsId() *string
	SetSrcDsName(v string) *AddDataCheckTaskRequest
	GetSrcDsName() *string
	SetSrcDsType(v string) *AddDataCheckTaskRequest
	GetSrcDsType() *string
	SetTaskMode(v int32) *AddDataCheckTaskRequest
	GetTaskMode() *int32
	SetTaskName(v string) *AddDataCheckTaskRequest
	GetTaskName() *string
}

type AddDataCheckTaskRequest struct {
	// The validation template ID. If not specified, the built-in default template is used.
	//
	// example:
	//
	// 1001
	CheckTemplateId *string `json:"checkTemplateId,omitempty" xml:"checkTemplateId,omitempty"`
	// The validation type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The ID of the destination data source.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Hive
	DstDsType *string `json:"dstDsType,omitempty" xml:"dstDsType,omitempty"`
	// The ID of the source data source.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Hive
	SrcDsType *string `json:"srcDsType,omitempty" xml:"srcDsType,omitempty"`
	// The table detail creation mode. Valid values:
	//
	// - 0: table-by-table fine-grained creation.
	//
	// - 1: batch creation with the same schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TaskMode *int32 `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The task name. Only Chinese characters, English characters, and digits are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s AddDataCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *AddDataCheckTaskRequest) GetCheckTemplateId() *string {
	return s.CheckTemplateId
}

func (s *AddDataCheckTaskRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *AddDataCheckTaskRequest) GetDstDsId() *string {
	return s.DstDsId
}

func (s *AddDataCheckTaskRequest) GetDstDsName() *string {
	return s.DstDsName
}

func (s *AddDataCheckTaskRequest) GetDstDsType() *string {
	return s.DstDsType
}

func (s *AddDataCheckTaskRequest) GetSrcDsId() *string {
	return s.SrcDsId
}

func (s *AddDataCheckTaskRequest) GetSrcDsName() *string {
	return s.SrcDsName
}

func (s *AddDataCheckTaskRequest) GetSrcDsType() *string {
	return s.SrcDsType
}

func (s *AddDataCheckTaskRequest) GetTaskMode() *int32 {
	return s.TaskMode
}

func (s *AddDataCheckTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *AddDataCheckTaskRequest) SetCheckTemplateId(v string) *AddDataCheckTaskRequest {
	s.CheckTemplateId = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetCheckType(v int32) *AddDataCheckTaskRequest {
	s.CheckType = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetDstDsId(v string) *AddDataCheckTaskRequest {
	s.DstDsId = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetDstDsName(v string) *AddDataCheckTaskRequest {
	s.DstDsName = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetDstDsType(v string) *AddDataCheckTaskRequest {
	s.DstDsType = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetSrcDsId(v string) *AddDataCheckTaskRequest {
	s.SrcDsId = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetSrcDsName(v string) *AddDataCheckTaskRequest {
	s.SrcDsName = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetSrcDsType(v string) *AddDataCheckTaskRequest {
	s.SrcDsType = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetTaskMode(v int32) *AddDataCheckTaskRequest {
	s.TaskMode = &v
	return s
}

func (s *AddDataCheckTaskRequest) SetTaskName(v string) *AddDataCheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *AddDataCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
