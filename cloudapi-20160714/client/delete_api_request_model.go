// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DeleteApiRequest
	GetApiId() *string
	SetGroupId(v string) *DeleteApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DeleteApiRequest
	GetSecurityToken() *string
}

type DeleteApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6f679aeb3be4b91b3688e887ca1fe16
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DeleteApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteApiRequest) SetApiId(v string) *DeleteApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteApiRequest) SetGroupId(v string) *DeleteApiRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiRequest) SetSecurityToken(v string) *DeleteApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteApiRequest) Validate() error {
	return dara.Validate(s)
}
