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
	// The VPC private network access whitelist or whitelist group configuration of the instance.
	//
	// example:
	//
	// {"privateNetworkIpWhiteList":["0.0.XX.XX","10.2.XX.XX","192.168.XX.XX/25"]}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The modification method. Valid values:
	//
	// - Cover (default): overwrites the original IP address whitelist with the value of the ips parameter.
	//
	// - Append: adds the IP addresses specified by the ips parameter to the original IP address whitelist.
	//
	// - Delete: deletes the IP addresses specified by the ips parameter from the original IP address whitelist. You must retain at least one IP address.
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
