// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyAIDBClusterModelResponseBody
	GetDryRun() *bool
	SetModelType(v string) *ModifyAIDBClusterModelResponseBody
	GetModelType() *string
	SetRequestId(v string) *ModifyAIDBClusterModelResponseBody
	GetRequestId() *string
	SetTargetModelName(v string) *ModifyAIDBClusterModelResponseBody
	GetTargetModelName() *string
	SetTargetOssPath(v string) *ModifyAIDBClusterModelResponseBody
	GetTargetOssPath() *string
	SetTaskId(v int32) *ModifyAIDBClusterModelResponseBody
	GetTaskId() *int32
	SetTotalBatches(v int64) *ModifyAIDBClusterModelResponseBody
	GetTotalBatches() *int64
	SetTotalMsds(v int64) *ModifyAIDBClusterModelResponseBody
	GetTotalMsds() *int64
	SetWarnings(v []*string) *ModifyAIDBClusterModelResponseBody
	GetWarnings() []*string
}

type ModifyAIDBClusterModelResponseBody struct {
	// Indicates whether the request is a dry-run request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The model type of the instance.
	//
	// example:
	//
	// custom
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3AA6E0E4-1234-5678-90AB-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resolved target model name.
	//
	// example:
	//
	// Qwen3-32B
	TargetModelName *string `json:"TargetModelName,omitempty" xml:"TargetModelName,omitempty"`
	// The resolved target OSS path.
	//
	// example:
	//
	// /my-model-bucket/models/qwen3
	TargetOssPath *string `json:"TargetOssPath,omitempty" xml:"TargetOssPath,omitempty"`
	// The ID of the asynchronous task. This parameter is empty when DryRun is set to true.
	//
	// example:
	//
	// 123456
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The number of change batches.
	//
	// example:
	//
	// 1
	TotalBatches *int64 `json:"TotalBatches,omitempty" xml:"TotalBatches,omitempty"`
	// The number of affected model serving instances.
	//
	// example:
	//
	// 2
	TotalMsds *int64 `json:"TotalMsds,omitempty" xml:"TotalMsds,omitempty"`
	// The change warnings returned by the upstream. The caller must display these warnings.
	Warnings []*string `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
}

func (s ModifyAIDBClusterModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterModelResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAIDBClusterModelResponseBody) GetModelType() *string {
	return s.ModelType
}

func (s *ModifyAIDBClusterModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAIDBClusterModelResponseBody) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *ModifyAIDBClusterModelResponseBody) GetTargetOssPath() *string {
	return s.TargetOssPath
}

func (s *ModifyAIDBClusterModelResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyAIDBClusterModelResponseBody) GetTotalBatches() *int64 {
	return s.TotalBatches
}

func (s *ModifyAIDBClusterModelResponseBody) GetTotalMsds() *int64 {
	return s.TotalMsds
}

func (s *ModifyAIDBClusterModelResponseBody) GetWarnings() []*string {
	return s.Warnings
}

func (s *ModifyAIDBClusterModelResponseBody) SetDryRun(v bool) *ModifyAIDBClusterModelResponseBody {
	s.DryRun = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetModelType(v string) *ModifyAIDBClusterModelResponseBody {
	s.ModelType = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetRequestId(v string) *ModifyAIDBClusterModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetTargetModelName(v string) *ModifyAIDBClusterModelResponseBody {
	s.TargetModelName = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetTargetOssPath(v string) *ModifyAIDBClusterModelResponseBody {
	s.TargetOssPath = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetTaskId(v int32) *ModifyAIDBClusterModelResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetTotalBatches(v int64) *ModifyAIDBClusterModelResponseBody {
	s.TotalBatches = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetTotalMsds(v int64) *ModifyAIDBClusterModelResponseBody {
	s.TotalMsds = &v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) SetWarnings(v []*string) *ModifyAIDBClusterModelResponseBody {
	s.Warnings = v
	return s
}

func (s *ModifyAIDBClusterModelResponseBody) Validate() error {
	return dara.Validate(s)
}
