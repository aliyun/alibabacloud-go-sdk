// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchChangeTableOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCrossTenant(v bool) *SubmitBatchChangeTableOwnerRequest
	GetEnableCrossTenant() *bool
	SetOwner(v string) *SubmitBatchChangeTableOwnerRequest
	GetOwner() *string
	SetTableMetaEntityIds(v []*string) *SubmitBatchChangeTableOwnerRequest
	GetTableMetaEntityIds() []*string
}

type SubmitBatchChangeTableOwnerRequest struct {
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
	TableMetaEntityIds []*string `json:"TableMetaEntityIds,omitempty" xml:"TableMetaEntityIds,omitempty" type:"Repeated"`
}

func (s SubmitBatchChangeTableOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchChangeTableOwnerRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchChangeTableOwnerRequest) GetEnableCrossTenant() *bool {
	return s.EnableCrossTenant
}

func (s *SubmitBatchChangeTableOwnerRequest) GetOwner() *string {
	return s.Owner
}

func (s *SubmitBatchChangeTableOwnerRequest) GetTableMetaEntityIds() []*string {
	return s.TableMetaEntityIds
}

func (s *SubmitBatchChangeTableOwnerRequest) SetEnableCrossTenant(v bool) *SubmitBatchChangeTableOwnerRequest {
	s.EnableCrossTenant = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerRequest) SetOwner(v string) *SubmitBatchChangeTableOwnerRequest {
	s.Owner = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerRequest) SetTableMetaEntityIds(v []*string) *SubmitBatchChangeTableOwnerRequest {
	s.TableMetaEntityIds = v
	return s
}

func (s *SubmitBatchChangeTableOwnerRequest) Validate() error {
	return dara.Validate(s)
}
