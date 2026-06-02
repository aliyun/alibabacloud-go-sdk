// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *GetUserRequest
	GetChannel() *string
	SetRegion(v string) *GetUserRequest
	GetRegion() *string
}

type GetUserRequest struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	Region  *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetUserRequest) GetRegion() *string {
	return s.Region
}

func (s *GetUserRequest) SetChannel(v string) *GetUserRequest {
	s.Channel = &v
	return s
}

func (s *GetUserRequest) SetRegion(v string) *GetUserRequest {
	s.Region = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
