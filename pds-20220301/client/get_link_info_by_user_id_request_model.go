// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkInfoByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *GetLinkInfoByUserIdRequest
	GetUserId() *string
}

type GetLinkInfoByUserIdRequest struct {
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetLinkInfoByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLinkInfoByUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetLinkInfoByUserIdRequest) SetUserId(v string) *GetLinkInfoByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *GetLinkInfoByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
