// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *GetUserGroupInfoRequest
	GetKeyword() *string
}

type GetUserGroupInfoRequest struct {
	// Keyword of the user group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s GetUserGroupInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetUserGroupInfoRequest) SetKeyword(v string) *GetUserGroupInfoRequest {
	s.Keyword = &v
	return s
}

func (s *GetUserGroupInfoRequest) Validate() error {
	return dara.Validate(s)
}
