// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *GetYikeUserRequest
	GetUserName() *string
}

type GetYikeUserRequest struct {
	// example:
	//
	// test.xxx@xxx.yikeai
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetYikeUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserRequest) GoString() string {
	return s.String()
}

func (s *GetYikeUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetYikeUserRequest) SetUserName(v string) *GetYikeUserRequest {
	s.UserName = &v
	return s
}

func (s *GetYikeUserRequest) Validate() error {
	return dara.Validate(s)
}
