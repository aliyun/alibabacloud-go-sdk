// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelDeletionRequest
	GetClientToken() *string
}

type CancelDeletionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CancelDeletionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelDeletionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelDeletionRequest) SetClientToken(v string) *CancelDeletionRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDeletionRequest) Validate() error {
	return dara.Validate(s)
}
