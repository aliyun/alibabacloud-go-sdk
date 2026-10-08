// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityOnResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommandShrink(v string) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest
	GetCheckCommandShrink() *string
	SetOpTenantId(v int64) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest
	GetOpUserId() *string
}

type CheckDataSourceConnectivityOnResourceGroupShrinkRequest struct {
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
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

func (s CheckDataSourceConnectivityOnResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityOnResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) GetCheckCommandShrink() *string {
	return s.CheckCommandShrink
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) SetCheckCommandShrink(v string) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) SetOpUserId(v string) *CheckDataSourceConnectivityOnResourceGroupShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckDataSourceConnectivityOnResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
