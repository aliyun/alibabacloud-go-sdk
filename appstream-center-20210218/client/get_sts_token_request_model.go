// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStsTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *GetStsTokenRequest
	GetEndUserId() *string
	SetExpiration(v int64) *GetStsTokenRequest
	GetExpiration() *int64
	SetExternalId(v string) *GetStsTokenRequest
	GetExternalId() *string
}

type GetStsTokenRequest struct {
	// The Elastic Desktop Service (EDS) username.
	//
	// > Either EndUserId or ExternalId is required.
	//
	// example:
	//
	// Alex
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The validity period of the token, in seconds. The maximum period is two days.
	//
	// example:
	//
	// 600
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The ID of the external user.
	//
	// example:
	//
	// Alex
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
}

func (s GetStsTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStsTokenRequest) GoString() string {
	return s.String()
}

func (s *GetStsTokenRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetStsTokenRequest) GetExpiration() *int64 {
	return s.Expiration
}

func (s *GetStsTokenRequest) GetExternalId() *string {
	return s.ExternalId
}

func (s *GetStsTokenRequest) SetEndUserId(v string) *GetStsTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *GetStsTokenRequest) SetExpiration(v int64) *GetStsTokenRequest {
	s.Expiration = &v
	return s
}

func (s *GetStsTokenRequest) SetExternalId(v string) *GetStsTokenRequest {
	s.ExternalId = &v
	return s
}

func (s *GetStsTokenRequest) Validate() error {
	return dara.Validate(s)
}
