// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOtpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUsername(v string) *DeleteOtpConfigRequest
	GetUsername() *string
}

type DeleteOtpConfigRequest struct {
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteOtpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOtpConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteOtpConfigRequest) GetUsername() *string {
	return s.Username
}

func (s *DeleteOtpConfigRequest) SetUsername(v string) *DeleteOtpConfigRequest {
	s.Username = &v
	return s
}

func (s *DeleteOtpConfigRequest) Validate() error {
	return dara.Validate(s)
}
