// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSvcMapBindShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *BatchDeleteSvcMapBindShrinkRequest
	GetIdsShrink() *string
	SetJwtToken(v string) *BatchDeleteSvcMapBindShrinkRequest
	GetJwtToken() *string
}

type BatchDeleteSvcMapBindShrinkRequest struct {
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s BatchDeleteSvcMapBindShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSvcMapBindShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteSvcMapBindShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *BatchDeleteSvcMapBindShrinkRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *BatchDeleteSvcMapBindShrinkRequest) SetIdsShrink(v string) *BatchDeleteSvcMapBindShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *BatchDeleteSvcMapBindShrinkRequest) SetJwtToken(v string) *BatchDeleteSvcMapBindShrinkRequest {
	s.JwtToken = &v
	return s
}

func (s *BatchDeleteSvcMapBindShrinkRequest) Validate() error {
	return dara.Validate(s)
}
