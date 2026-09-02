// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchChangeTableOwnerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCrossTenant(v bool) *SubmitBatchChangeTableOwnerShrinkRequest
	GetEnableCrossTenant() *bool
	SetOwner(v string) *SubmitBatchChangeTableOwnerShrinkRequest
	GetOwner() *string
	SetTableMetaEntityIdsShrink(v string) *SubmitBatchChangeTableOwnerShrinkRequest
	GetTableMetaEntityIdsShrink() *string
}

type SubmitBatchChangeTableOwnerShrinkRequest struct {
	// example:
	//
	// false
	EnableCrossTenant *bool `json:"EnableCrossTenant,omitempty" xml:"EnableCrossTenant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2xxxxx8146415628
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	TableMetaEntityIdsShrink *string `json:"TableMetaEntityIds,omitempty" xml:"TableMetaEntityIds,omitempty"`
}

func (s SubmitBatchChangeTableOwnerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchChangeTableOwnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) GetEnableCrossTenant() *bool {
	return s.EnableCrossTenant
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) GetTableMetaEntityIdsShrink() *string {
	return s.TableMetaEntityIdsShrink
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) SetEnableCrossTenant(v bool) *SubmitBatchChangeTableOwnerShrinkRequest {
	s.EnableCrossTenant = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) SetOwner(v string) *SubmitBatchChangeTableOwnerShrinkRequest {
	s.Owner = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) SetTableMetaEntityIdsShrink(v string) *SubmitBatchChangeTableOwnerShrinkRequest {
	s.TableMetaEntityIdsShrink = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
