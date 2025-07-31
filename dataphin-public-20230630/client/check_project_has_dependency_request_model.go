// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProjectHasDependencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CheckProjectHasDependencyRequest
	GetId() *int64
	SetOpTenantId(v int64) *CheckProjectHasDependencyRequest
	GetOpTenantId() *int64
}

type CheckProjectHasDependencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckProjectHasDependencyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckProjectHasDependencyRequest) GoString() string {
	return s.String()
}

func (s *CheckProjectHasDependencyRequest) GetId() *int64 {
	return s.Id
}

func (s *CheckProjectHasDependencyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckProjectHasDependencyRequest) SetId(v int64) *CheckProjectHasDependencyRequest {
	s.Id = &v
	return s
}

func (s *CheckProjectHasDependencyRequest) SetOpTenantId(v int64) *CheckProjectHasDependencyRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckProjectHasDependencyRequest) Validate() error {
	return dara.Validate(s)
}
