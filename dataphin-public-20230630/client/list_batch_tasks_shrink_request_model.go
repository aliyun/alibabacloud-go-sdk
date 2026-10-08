// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchTaskQueryShrink(v string) *ListBatchTasksShrinkRequest
	GetBatchTaskQueryShrink() *string
	SetOpTenantId(v int64) *ListBatchTasksShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListBatchTasksShrinkRequest
	GetOpUserId() *string
}

type ListBatchTasksShrinkRequest struct {
	// This parameter is required.
	BatchTaskQueryShrink *string `json:"BatchTaskQuery,omitempty" xml:"BatchTaskQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListBatchTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTasksShrinkRequest) GetBatchTaskQueryShrink() *string {
	return s.BatchTaskQueryShrink
}

func (s *ListBatchTasksShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBatchTasksShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListBatchTasksShrinkRequest) SetBatchTaskQueryShrink(v string) *ListBatchTasksShrinkRequest {
	s.BatchTaskQueryShrink = &v
	return s
}

func (s *ListBatchTasksShrinkRequest) SetOpTenantId(v int64) *ListBatchTasksShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBatchTasksShrinkRequest) SetOpUserId(v string) *ListBatchTasksShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListBatchTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
