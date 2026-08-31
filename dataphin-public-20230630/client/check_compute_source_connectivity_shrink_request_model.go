// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommandShrink(v string) *CheckComputeSourceConnectivityShrinkRequest
	GetCheckCommandShrink() *string
	SetOpTenantId(v int64) *CheckComputeSourceConnectivityShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CheckComputeSourceConnectivityShrinkRequest
	GetOpUserId() *string
}

type CheckComputeSourceConnectivityShrinkRequest struct {
	// The check command.
	//
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
	// The tenant ID.
	//
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

func (s CheckComputeSourceConnectivityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityShrinkRequest) GetCheckCommandShrink() *string {
	return s.CheckCommandShrink
}

func (s *CheckComputeSourceConnectivityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckComputeSourceConnectivityShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckComputeSourceConnectivityShrinkRequest) SetCheckCommandShrink(v string) *CheckComputeSourceConnectivityShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckComputeSourceConnectivityShrinkRequest) SetOpTenantId(v int64) *CheckComputeSourceConnectivityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckComputeSourceConnectivityShrinkRequest) SetOpUserId(v string) *CheckComputeSourceConnectivityShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckComputeSourceConnectivityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
