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
	// - Access token expiration time (in seconds), which is the expiration time for the user to access the page interface. The default value is 86400 seconds (one day), and the range of values is from 0 to 86400 seconds (one day).
	//
	// - The access token expiration time is the minimum value between `accessTokenExpirationTime` and `expirationTime`.
	//
	// - If called through STS, the access token expiration time (i.e., the time during which the user can access the page interface) is the minimum value among `accessTokenExpirationTime`, `expirationTime`, and the STS expiration time.
	//
	// example:
	//
	// 600
	AccessTokenExpirationTime *int64 `json:"accessTokenExpirationTime,omitempty" xml:"accessTokenExpirationTime,omitempty"`
	// - Expiration time (in seconds), which is the expiration time for the embedded page URL. The default value is 86400 seconds (one day), and the range of values is from 0 to 2592000 seconds (30 days).
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
