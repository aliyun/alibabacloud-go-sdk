// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsShared(v bool) *GetRayDashboardRequest
	GetIsShared() *bool
	SetToken(v string) *GetRayDashboardRequest
	GetToken() *string
}

type GetRayDashboardRequest struct {
	// Specifies whether the link is a sharing link. If yes, a token is required.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsShared *bool `json:"isShared,omitempty" xml:"isShared,omitempty"`
	// The token obtained from GetToken
	//
	// example:
	//
	// some_token_value
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetRayDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRayDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetRayDashboardRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *GetRayDashboardRequest) GetToken() *string {
	return s.Token
}

func (s *GetRayDashboardRequest) SetIsShared(v bool) *GetRayDashboardRequest {
	s.IsShared = &v
	return s
}

func (s *GetRayDashboardRequest) SetToken(v string) *GetRayDashboardRequest {
	s.Token = &v
	return s
}

func (s *GetRayDashboardRequest) Validate() error {
	return dara.Validate(s)
}
