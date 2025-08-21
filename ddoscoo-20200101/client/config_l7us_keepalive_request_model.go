// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7UsKeepaliveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigL7UsKeepaliveRequest
	GetDomain() *string
	SetUpstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest
	GetUpstreamKeepalive() *string
}

type ConfigL7UsKeepaliveRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The settings for back-to-origin persistent connections. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **enabled**: the switch for back-to-origin persistent connections. This field is required, and the value is of the Boolean type.
	//
	// 	- **keepalive_requests**: the number of requests that reuse a persistent connection. This field is required, and the value is of the integer type.
	//
	// 	- **keepalive_timeout**: the timeout period for an idle persistent connection. This field is required, and the value is of the integer type.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"enabled": true, "keepalive_requests": 1000,"keepalive_timeout": 30}
	UpstreamKeepalive *string `json:"UpstreamKeepalive,omitempty" xml:"UpstreamKeepalive,omitempty"`
}

func (s ConfigL7UsKeepaliveRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7UsKeepaliveRequest) GoString() string {
	return s.String()
}

func (s *ConfigL7UsKeepaliveRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigL7UsKeepaliveRequest) GetUpstreamKeepalive() *string {
	return s.UpstreamKeepalive
}

func (s *ConfigL7UsKeepaliveRequest) SetDomain(v string) *ConfigL7UsKeepaliveRequest {
	s.Domain = &v
	return s
}

func (s *ConfigL7UsKeepaliveRequest) SetUpstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest {
	s.UpstreamKeepalive = &v
	return s
}

func (s *ConfigL7UsKeepaliveRequest) Validate() error {
	return dara.Validate(s)
}
