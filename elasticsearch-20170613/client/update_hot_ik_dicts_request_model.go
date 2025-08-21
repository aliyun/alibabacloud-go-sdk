// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotIkDictsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateHotIkDictsRequest
	GetBody() *string
	SetClientToken(v string) *UpdateHotIkDictsRequest
	GetClientToken() *string
}

type UpdateHotIkDictsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateHotIkDictsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotIkDictsRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotIkDictsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateHotIkDictsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateHotIkDictsRequest) SetBody(v string) *UpdateHotIkDictsRequest {
	s.Body = &v
	return s
}

func (s *UpdateHotIkDictsRequest) SetClientToken(v string) *UpdateHotIkDictsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateHotIkDictsRequest) Validate() error {
	return dara.Validate(s)
}
