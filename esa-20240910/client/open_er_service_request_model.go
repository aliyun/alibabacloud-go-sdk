// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenErServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenErServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *OpenErServiceRequest
	GetSecurityToken() *string
}

type OpenErServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenErServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenErServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenErServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenErServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenErServiceRequest) SetOwnerId(v int64) *OpenErServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenErServiceRequest) SetSecurityToken(v string) *OpenErServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenErServiceRequest) Validate() error {
	return dara.Validate(s)
}
