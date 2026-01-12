// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSourcePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *CreateSourcePolicyRequest
	GetJwtToken() *string
	SetOssKey(v string) *CreateSourcePolicyRequest
	GetOssKey() *string
}

type CreateSourcePolicyRequest struct {
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s CreateSourcePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSourcePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSourcePolicyRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *CreateSourcePolicyRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateSourcePolicyRequest) SetJwtToken(v string) *CreateSourcePolicyRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateSourcePolicyRequest) SetOssKey(v string) *CreateSourcePolicyRequest {
	s.OssKey = &v
	return s
}

func (s *CreateSourcePolicyRequest) Validate() error {
	return dara.Validate(s)
}
