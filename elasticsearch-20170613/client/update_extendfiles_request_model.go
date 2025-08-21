// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtendfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateExtendfilesRequest
	GetClientToken() *string
	SetBody(v string) *UpdateExtendfilesRequest
	GetBody() *string
}

type UpdateExtendfilesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtendfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendfilesRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtendfilesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExtendfilesRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateExtendfilesRequest) SetClientToken(v string) *UpdateExtendfilesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExtendfilesRequest) SetBody(v string) *UpdateExtendfilesRequest {
	s.Body = &v
	return s
}

func (s *UpdateExtendfilesRequest) Validate() error {
	return dara.Validate(s)
}
