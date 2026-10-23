// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *RenewLogstashRequest
	GetBody() *string
	SetClientToken(v string) *RenewLogstashRequest
	GetClientToken() *string
}

type RenewLogstashRequest struct {
	// The renewal duration and cycle for the Logstash instance.
	//
	// example:
	//
	// {"duration":1,"pricingCycle":"Year"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must ensure that it is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RenewLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewLogstashRequest) GoString() string {
	return s.String()
}

func (s *RenewLogstashRequest) GetBody() *string {
	return s.Body
}

func (s *RenewLogstashRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewLogstashRequest) SetBody(v string) *RenewLogstashRequest {
	s.Body = &v
	return s
}

func (s *RenewLogstashRequest) SetClientToken(v string) *RenewLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewLogstashRequest) Validate() error {
	return dara.Validate(s)
}
