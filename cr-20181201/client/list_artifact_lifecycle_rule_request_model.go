// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableDeleteTag(v bool) *ListArtifactLifecycleRuleRequest
	GetEnableDeleteTag() *bool
	SetInstanceId(v string) *ListArtifactLifecycleRuleRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListArtifactLifecycleRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactLifecycleRuleRequest
	GetPageSize() *int32
}

type ListArtifactLifecycleRuleRequest struct {
	// Specifies whether lifecycle management is enabled.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The ID of the Enterprise instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-eztul9ucz76q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If the specified value exceeds 100, the system returns a parameter error or uses 100 as the actual maximum number of entries returned.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleRequest) GetEnableDeleteTag() *bool {
	return s.EnableDeleteTag
}

func (s *ListArtifactLifecycleRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactLifecycleRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactLifecycleRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *ListArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetInstanceId(v string) *ListArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetPageNo(v int32) *ListArtifactLifecycleRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetPageSize(v int32) *ListArtifactLifecycleRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
