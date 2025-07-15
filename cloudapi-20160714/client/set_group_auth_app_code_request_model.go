// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGroupAuthAppCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthAppCode(v string) *SetGroupAuthAppCodeRequest
	GetAuthAppCode() *string
	SetGroupId(v string) *SetGroupAuthAppCodeRequest
	GetGroupId() *string
	SetSecurityToken(v string) *SetGroupAuthAppCodeRequest
	GetSecurityToken() *string
}

type SetGroupAuthAppCodeRequest struct {
	// This parameter is required.
	AuthAppCode *string `json:"AuthAppCode,omitempty" xml:"AuthAppCode,omitempty"`
	// This parameter is required.
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetGroupAuthAppCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetGroupAuthAppCodeRequest) GoString() string {
	return s.String()
}

func (s *SetGroupAuthAppCodeRequest) GetAuthAppCode() *string {
	return s.AuthAppCode
}

func (s *SetGroupAuthAppCodeRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetGroupAuthAppCodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetGroupAuthAppCodeRequest) SetAuthAppCode(v string) *SetGroupAuthAppCodeRequest {
	s.AuthAppCode = &v
	return s
}

func (s *SetGroupAuthAppCodeRequest) SetGroupId(v string) *SetGroupAuthAppCodeRequest {
	s.GroupId = &v
	return s
}

func (s *SetGroupAuthAppCodeRequest) SetSecurityToken(v string) *SetGroupAuthAppCodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetGroupAuthAppCodeRequest) Validate() error {
	return dara.Validate(s)
}
