// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListArtifactSubscriptionTaskRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListArtifactSubscriptionTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactSubscriptionTaskRequest
	GetPageSize() *int32
}

type ListArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-m9ob8792vm****
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
}

func (s ListArtifactSubscriptionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactSubscriptionTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactSubscriptionTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactSubscriptionTaskRequest) SetInstanceId(v string) *ListArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskRequest) SetPageNo(v int32) *ListArtifactSubscriptionTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionTaskRequest) SetPageSize(v int32) *ListArtifactSubscriptionTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionTaskRequest) Validate() error {
	return dara.Validate(s)
}
