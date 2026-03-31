// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *DeleteLoginProfileRequest
	GetUserName() *string
}

type DeleteLoginProfileRequest struct {
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteLoginProfileRequest) SetUserName(v string) *DeleteLoginProfileRequest {
	s.UserName = &v
	return s
}

func (s *DeleteLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
