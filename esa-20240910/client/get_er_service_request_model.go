// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetErServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *GetErServiceRequest
	GetSecurityToken() *string
}

type GetErServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetErServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErServiceRequest) GoString() string {
	return s.String()
}

func (s *GetErServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetErServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetErServiceRequest) SetOwnerId(v int64) *GetErServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetErServiceRequest) SetSecurityToken(v string) *GetErServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetErServiceRequest) Validate() error {
	return dara.Validate(s)
}
