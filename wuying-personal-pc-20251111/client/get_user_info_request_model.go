// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *GetUserInfoRequest
	GetApiKey() *string
}

type GetUserInfoRequest struct {
	// This parameter is required.
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetUserInfoRequest) SetApiKey(v string) *GetUserInfoRequest {
	s.ApiKey = &v
	return s
}

func (s *GetUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
