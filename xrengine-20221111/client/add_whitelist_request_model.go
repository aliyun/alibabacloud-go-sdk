// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *AddWhitelistRequest
	GetAliyunUid() *string
	SetJwtToken(v string) *AddWhitelistRequest
	GetJwtToken() *string
	SetRemark(v string) *AddWhitelistRequest
	GetRemark() *string
}

type AddWhitelistRequest struct {
	// This parameter is required.
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWhitelistRequest) GoString() string {
	return s.String()
}

func (s *AddWhitelistRequest) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *AddWhitelistRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *AddWhitelistRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddWhitelistRequest) SetAliyunUid(v string) *AddWhitelistRequest {
	s.AliyunUid = &v
	return s
}

func (s *AddWhitelistRequest) SetJwtToken(v string) *AddWhitelistRequest {
	s.JwtToken = &v
	return s
}

func (s *AddWhitelistRequest) SetRemark(v string) *AddWhitelistRequest {
	s.Remark = &v
	return s
}

func (s *AddWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
