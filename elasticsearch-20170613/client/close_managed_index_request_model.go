// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseManagedIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CloseManagedIndexRequest
	GetClientToken() *string
}

type CloseManagedIndexRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CloseManagedIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseManagedIndexRequest) GoString() string {
	return s.String()
}

func (s *CloseManagedIndexRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloseManagedIndexRequest) SetClientToken(v string) *CloseManagedIndexRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseManagedIndexRequest) Validate() error {
	return dara.Validate(s)
}
