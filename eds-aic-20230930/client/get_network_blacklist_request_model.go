// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetNetworkBlacklistRequest
	GetType() *string
}

type GetNetworkBlacklistRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNetworkBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkBlacklistRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkBlacklistRequest) GetType() *string {
	return s.Type
}

func (s *GetNetworkBlacklistRequest) SetType(v string) *GetNetworkBlacklistRequest {
	s.Type = &v
	return s
}

func (s *GetNetworkBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
