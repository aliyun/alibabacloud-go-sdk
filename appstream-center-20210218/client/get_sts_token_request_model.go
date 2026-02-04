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
	// example:
	//
	// Alex
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// 600
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
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
