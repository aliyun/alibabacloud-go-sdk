// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *GetLoginProfileRequest
	GetUserName() *string
}

type GetLoginProfileRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *GetLoginProfileRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetLoginProfileRequest) SetUserName(v string) *GetLoginProfileRequest {
	s.UserName = &v
	return s
}

func (s *GetLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
