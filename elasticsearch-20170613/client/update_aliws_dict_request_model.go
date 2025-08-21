// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliwsDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateAliwsDictRequest
	GetBody() *string
	SetClientToken(v string) *UpdateAliwsDictRequest
	GetClientToken() *string
}

type UpdateAliwsDictRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAliwsDictRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliwsDictRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateAliwsDictRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAliwsDictRequest) SetBody(v string) *UpdateAliwsDictRequest {
	s.Body = &v
	return s
}

func (s *UpdateAliwsDictRequest) SetClientToken(v string) *UpdateAliwsDictRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAliwsDictRequest) Validate() error {
	return dara.Validate(s)
}
