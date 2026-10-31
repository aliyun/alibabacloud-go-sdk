// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8SResourceGroupIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListK8SResourceGroupIdsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListK8SResourceGroupIdsRequest
	GetOpUserId() *string
}

type ListK8SResourceGroupIdsRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListK8SResourceGroupIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8SResourceGroupIdsRequest) GoString() string {
	return s.String()
}

func (s *ListK8SResourceGroupIdsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListK8SResourceGroupIdsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListK8SResourceGroupIdsRequest) SetOpTenantId(v int64) *ListK8SResourceGroupIdsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListK8SResourceGroupIdsRequest) SetOpUserId(v string) *ListK8SResourceGroupIdsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListK8SResourceGroupIdsRequest) Validate() error {
	return dara.Validate(s)
}
