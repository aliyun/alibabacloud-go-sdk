// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipalName(v string) *DeleteLoginProfileRequest
	GetUserPrincipalName() *string
}

type DeleteLoginProfileRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DeleteLoginProfileRequest) SetUserPrincipalName(v string) *DeleteLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeleteLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
