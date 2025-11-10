// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsShared(v bool) *GetDashboardRequest
	GetIsShared() *bool
	SetToken(v string) *GetDashboardRequest
	GetToken() *string
}

type GetDashboardRequest struct {
	// Specifies whether the link is a sharing link. If yes, a token is required.
	//
	// Enumerated values:
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

func (s GetDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *GetDashboardRequest) GetToken() *string {
	return s.Token
}

func (s *GetDashboardRequest) SetIsShared(v bool) *GetDashboardRequest {
	s.IsShared = &v
	return s
}

func (s *GetDashboardRequest) SetToken(v string) *GetDashboardRequest {
	s.Token = &v
	return s
}

func (s *GetDashboardRequest) Validate() error {
	return dara.Validate(s)
}
