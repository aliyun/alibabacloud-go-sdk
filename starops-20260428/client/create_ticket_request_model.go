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
	// - The access token expiration time, in seconds. This specifies the expiration time for the user to access page operations. Default value: 86400 (one day). Valid values: 0 to 86400 (one day).
	//
	// - The actual access token expiration time is the minimum value of accessTokenExpirationTime and expirationTime.
	//
	// - If you call this operation by using Security Token Service (STS), the actual access token expiration time is the minimum value of accessTokenExpirationTime, expirationTime, and the STS token expiration time.
	//
	// example:
	//
	// 600
	AccessTokenExpirationTime *int64 `json:"accessTokenExpirationTime,omitempty" xml:"accessTokenExpirationTime,omitempty"`
	// - The expiration time, in seconds. This specifies the expiration time of the embedded page URL. Default value: 86400 (one day). Valid values: 0 to 2592000 (30 days).
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
