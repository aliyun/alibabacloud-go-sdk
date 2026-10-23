// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePublicWhiteIpsRequest
	GetBody() *string
	SetClientToken(v string) *UpdatePublicWhiteIpsRequest
	GetClientToken() *string
	SetModifyMode(v string) *UpdatePublicWhiteIpsRequest
	GetModifyMode() *string
}

type UpdatePublicWhiteIpsRequest struct {
	// The public network access whitelist or whitelist group configuration of the instance.
	//
	// example:
	//
	// {"publicIpWhitelist":["0.0.0.0/0","0.0.0.0/1"]}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The modification method. Valid values:
	//
	// - Cover: overwrites the original IP address whitelist with the value of the ips parameter. This is the default value.
	//
	// - Append: adds the IP addresses specified by the ips parameter to the original IP address whitelist.
	//
	// - Delete: deletes the IP addresses specified by the ips parameter from the original IP address whitelist. At least one IP address must be retained.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePublicWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePublicWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePublicWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdatePublicWhiteIpsRequest) SetBody(v string) *UpdatePublicWhiteIpsRequest {
	s.Body = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) SetClientToken(v string) *UpdatePublicWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) SetModifyMode(v string) *UpdatePublicWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) Validate() error {
	return dara.Validate(s)
}
