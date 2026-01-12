// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateSvcMapBindShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *BatchCreateSvcMapBindShrinkRequest
	GetJwtToken() *string
	SetMapIdsShrink(v string) *BatchCreateSvcMapBindShrinkRequest
	GetMapIdsShrink() *string
	SetSvcId(v int64) *BatchCreateSvcMapBindShrinkRequest
	GetSvcId() *int64
}

type BatchCreateSvcMapBindShrinkRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	// This parameter is required.
	MapIdsShrink *string `json:"MapIds,omitempty" xml:"MapIds,omitempty"`
	SvcId        *int64  `json:"SvcId,omitempty" xml:"SvcId,omitempty"`
}

func (s BatchCreateSvcMapBindShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateSvcMapBindShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateSvcMapBindShrinkRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *BatchCreateSvcMapBindShrinkRequest) GetMapIdsShrink() *string {
	return s.MapIdsShrink
}

func (s *BatchCreateSvcMapBindShrinkRequest) GetSvcId() *int64 {
	return s.SvcId
}

func (s *BatchCreateSvcMapBindShrinkRequest) SetJwtToken(v string) *BatchCreateSvcMapBindShrinkRequest {
	s.JwtToken = &v
	return s
}

func (s *BatchCreateSvcMapBindShrinkRequest) SetMapIdsShrink(v string) *BatchCreateSvcMapBindShrinkRequest {
	s.MapIdsShrink = &v
	return s
}

func (s *BatchCreateSvcMapBindShrinkRequest) SetSvcId(v int64) *BatchCreateSvcMapBindShrinkRequest {
	s.SvcId = &v
	return s
}

func (s *BatchCreateSvcMapBindShrinkRequest) Validate() error {
	return dara.Validate(s)
}
