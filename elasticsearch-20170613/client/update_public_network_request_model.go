// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePublicNetworkRequest
	GetBody() *string
	SetClientToken(v string) *UpdatePublicNetworkRequest
	GetClientToken() *string
}

type UpdatePublicNetworkRequest struct {
	// The public network access switch configuration for the instance.
	//
	// example:
	//
	// {"enablePublic":true}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must ensure that it is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdatePublicNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePublicNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePublicNetworkRequest) SetBody(v string) *UpdatePublicNetworkRequest {
	s.Body = &v
	return s
}

func (s *UpdatePublicNetworkRequest) SetClientToken(v string) *UpdatePublicNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicNetworkRequest) Validate() error {
	return dara.Validate(s)
}
