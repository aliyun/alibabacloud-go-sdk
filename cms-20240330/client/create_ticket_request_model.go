// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenExpirationTime(v int64) *CreateTicketRequest
	GetAccessTokenExpirationTime() *int64
	SetExpirationTime(v int64) *CreateTicketRequest
	GetExpirationTime() *int64
}

type CreateTicketRequest struct {
	// - The expiration time of the access token, in seconds. This is the period during which a user can access the page APIs. The value can range from 0 to 86,400 seconds (one day). The default value is 86,400 seconds (one day).
	//
	// - The effective expiration time of the access token is the minimum value of accessTokenExpirationTime and expirationTime.
	//
	// - If you call the operation using a Security Token Service (STS) token, the effective expiration time of the access token is the minimum value of accessTokenExpirationTime, expirationTime, and the expiration time of the STS token.
	//
	// example:
	//
	// 600
	AccessTokenExpirationTime *int64 `json:"accessTokenExpirationTime,omitempty" xml:"accessTokenExpirationTime,omitempty"`
	// - The expiration time of the URL for the embedded page, in seconds. The value can range from 0 to 2,592,000 seconds (30 days). The default value is 86,400 seconds (one day).
	//
	// example:
	//
	// 86400
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetAccessTokenExpirationTime() *int64 {
	return s.AccessTokenExpirationTime
}

func (s *CreateTicketRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *CreateTicketRequest) SetAccessTokenExpirationTime(v int64) *CreateTicketRequest {
	s.AccessTokenExpirationTime = &v
	return s
}

func (s *CreateTicketRequest) SetExpirationTime(v int64) *CreateTicketRequest {
	s.ExpirationTime = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	return dara.Validate(s)
}
