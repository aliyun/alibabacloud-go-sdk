// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLogstashDeletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelLogstashDeletionRequest
	GetClientToken() *string
}

type CancelLogstashDeletionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CancelLogstashDeletionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelLogstashDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelLogstashDeletionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelLogstashDeletionRequest) SetClientToken(v string) *CancelLogstashDeletionRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelLogstashDeletionRequest) Validate() error {
	return dara.Validate(s)
}
