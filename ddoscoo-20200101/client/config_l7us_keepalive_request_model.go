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
	SetDownstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest
	GetDownstreamKeepalive() *string
	SetUpstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest
	GetUpstreamKeepalive() *string
}

type ConfigL7UsKeepaliveRequest struct {
	// The domain name of the website service.
	//
	// > The domain name must have a website service forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The idle timeout period of the persistent connection between the client and Anti-DDoS Proxy.
	//
	// example:
	//
	// 60
	DownstreamKeepalive *string `json:"DownstreamKeepalive,omitempty" xml:"DownstreamKeepalive,omitempty"`
	// The persistent connection settings for back-to-origin requests of Anti-DDoS Proxy. The value is a string that represents a JSON struct. The JSON struct contains the following fields:
	//
	// - **enabled**: Boolean | Required | Specifies whether to enable persistent connections for back-to-origin requests.
	//
	// - **keepalive_requests**: Integer | Required | The number of requests that can be sent over a persistent connection.
	//
	// - **keepalive_timeout**: Integer | Required | The idle timeout period of a persistent connection.
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

func (s *ConfigL7UsKeepaliveRequest) GetDownstreamKeepalive() *string {
	return s.DownstreamKeepalive
}

func (s *ConfigL7UsKeepaliveRequest) GetUpstreamKeepalive() *string {
	return s.UpstreamKeepalive
}

func (s *ConfigL7UsKeepaliveRequest) SetDomain(v string) *ConfigL7UsKeepaliveRequest {
	s.Domain = &v
	return s
}

func (s *ConfigL7UsKeepaliveRequest) SetDownstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest {
	s.DownstreamKeepalive = &v
	return s
}

func (s *ConfigL7UsKeepaliveRequest) SetUpstreamKeepalive(v string) *ConfigL7UsKeepaliveRequest {
	s.UpstreamKeepalive = &v
	return s
}

func (s *ConfigL7UsKeepaliveRequest) Validate() error {
	return dara.Validate(s)
}
