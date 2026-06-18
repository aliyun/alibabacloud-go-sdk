// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertAiOutboundPhoneNumsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchVersion(v int32) *InsertAiOutboundPhoneNumsShrinkRequest
	GetBatchVersion() *int32
	SetDetailsShrink(v string) *InsertAiOutboundPhoneNumsShrinkRequest
	GetDetailsShrink() *string
	SetInstanceId(v string) *InsertAiOutboundPhoneNumsShrinkRequest
	GetInstanceId() *string
	SetTaskId(v int64) *InsertAiOutboundPhoneNumsShrinkRequest
	GetTaskId() *int64
}

type InsertAiOutboundPhoneNumsShrinkRequest struct {
	// The batch version number of the job.
	//
	// > If this parameter is not specified, numbers are imported into the default batch.
	//
	// example:
	//
	// 2
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
	// Number details.
	//
	// > A maximum of 30 entries are allowed.
	//
	// This parameter is required.
	DetailsShrink *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// You can invoke the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and check the **Data*	- field in the response, or invoke the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and check the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InsertAiOutboundPhoneNumsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) GetBatchVersion() *int32 {
	return s.BatchVersion
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) GetDetailsShrink() *string {
	return s.DetailsShrink
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) SetBatchVersion(v int32) *InsertAiOutboundPhoneNumsShrinkRequest {
	s.BatchVersion = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) SetDetailsShrink(v string) *InsertAiOutboundPhoneNumsShrinkRequest {
	s.DetailsShrink = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) SetInstanceId(v string) *InsertAiOutboundPhoneNumsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) SetTaskId(v int64) *InsertAiOutboundPhoneNumsShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
