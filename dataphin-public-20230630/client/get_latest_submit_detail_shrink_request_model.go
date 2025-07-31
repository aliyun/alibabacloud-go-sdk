// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestSubmitDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetLatestSubmitDetailShrinkRequest
	GetOpTenantId() *int64
	SetSubmitDetailQueryShrink(v string) *GetLatestSubmitDetailShrinkRequest
	GetSubmitDetailQueryShrink() *string
}

type GetLatestSubmitDetailShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitDetailQueryShrink *string `json:"SubmitDetailQuery,omitempty" xml:"SubmitDetailQuery,omitempty"`
}

func (s GetLatestSubmitDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetLatestSubmitDetailShrinkRequest) GetSubmitDetailQueryShrink() *string {
	return s.SubmitDetailQueryShrink
}

func (s *GetLatestSubmitDetailShrinkRequest) SetOpTenantId(v int64) *GetLatestSubmitDetailShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetLatestSubmitDetailShrinkRequest) SetSubmitDetailQueryShrink(v string) *GetLatestSubmitDetailShrinkRequest {
	s.SubmitDetailQueryShrink = &v
	return s
}

func (s *GetLatestSubmitDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
