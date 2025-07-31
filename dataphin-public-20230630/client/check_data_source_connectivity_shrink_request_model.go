// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommandShrink(v string) *CheckDataSourceConnectivityShrinkRequest
	GetCheckCommandShrink() *string
	SetOpTenantId(v int64) *CheckDataSourceConnectivityShrinkRequest
	GetOpTenantId() *int64
}

type CheckDataSourceConnectivityShrinkRequest struct {
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckDataSourceConnectivityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityShrinkRequest) GetCheckCommandShrink() *string {
	return s.CheckCommandShrink
}

func (s *CheckDataSourceConnectivityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckDataSourceConnectivityShrinkRequest) SetCheckCommandShrink(v string) *CheckDataSourceConnectivityShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckDataSourceConnectivityShrinkRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckDataSourceConnectivityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
