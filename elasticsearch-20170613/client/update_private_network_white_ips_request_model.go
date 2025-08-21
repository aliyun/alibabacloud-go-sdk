// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetworkWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePrivateNetworkWhiteIpsRequest
	GetBody() *string
	SetClientToken(v string) *UpdatePrivateNetworkWhiteIpsRequest
	GetClientToken() *string
	SetModifyMode(v string) *UpdatePrivateNetworkWhiteIpsRequest
	GetModifyMode() *string
}

type UpdatePrivateNetworkWhiteIpsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The results that are returned.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePrivateNetworkWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) SetBody(v string) *UpdatePrivateNetworkWhiteIpsRequest {
	s.Body = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) SetClientToken(v string) *UpdatePrivateNetworkWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) SetModifyMode(v string) *UpdatePrivateNetworkWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsRequest) Validate() error {
	return dara.Validate(s)
}
