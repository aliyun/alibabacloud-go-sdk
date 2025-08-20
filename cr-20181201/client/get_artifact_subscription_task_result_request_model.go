// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetArtifactSubscriptionTaskResultRequest
	GetInstanceId() *string
	SetPageNo(v int32) *GetArtifactSubscriptionTaskResultRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetArtifactSubscriptionTaskResultRequest
	GetPageSize() *int32
	SetTaskId(v string) *GetArtifactSubscriptionTaskResultRequest
	GetTaskId() *string
}

type GetArtifactSubscriptionTaskResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-90fxryf9pwf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crast-y64sq01bgad****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionTaskResultRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetArtifactSubscriptionTaskResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetArtifactSubscriptionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetInstanceId(v string) *GetArtifactSubscriptionTaskResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetPageNo(v int32) *GetArtifactSubscriptionTaskResultRequest {
	s.PageNo = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetPageSize(v int32) *GetArtifactSubscriptionTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetTaskId(v string) *GetArtifactSubscriptionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
