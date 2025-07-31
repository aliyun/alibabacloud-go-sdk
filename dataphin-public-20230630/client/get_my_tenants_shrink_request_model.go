// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyTenantsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureCodeListShrink(v string) *GetMyTenantsShrinkRequest
	GetFeatureCodeListShrink() *string
	SetOpTenantId(v int64) *GetMyTenantsShrinkRequest
	GetOpTenantId() *int64
}

type GetMyTenantsShrinkRequest struct {
	FeatureCodeListShrink *string `json:"FeatureCodeList,omitempty" xml:"FeatureCodeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetMyTenantsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMyTenantsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMyTenantsShrinkRequest) GetFeatureCodeListShrink() *string {
	return s.FeatureCodeListShrink
}

func (s *GetMyTenantsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetMyTenantsShrinkRequest) SetFeatureCodeListShrink(v string) *GetMyTenantsShrinkRequest {
	s.FeatureCodeListShrink = &v
	return s
}

func (s *GetMyTenantsShrinkRequest) SetOpTenantId(v int64) *GetMyTenantsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetMyTenantsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
