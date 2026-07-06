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
	// Indicates whether automatic tag deletion is enabled.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The ID of the Enterprise Edition instance.
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
	// The number of entries to return per page. The maximum value is 100. If you specify a value greater than 100, the system may return a parameter error or use a page size of 100.
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
