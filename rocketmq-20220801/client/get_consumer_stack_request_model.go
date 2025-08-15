// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *GetConsumerStackRequest
	GetClientId() *string
}

type GetConsumerStackRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.26.76.48@Lqd7dImlp9KJ5V84
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s GetConsumerStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerStackRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerStackRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetConsumerStackRequest) SetClientId(v string) *GetConsumerStackRequest {
	s.ClientId = &v
	return s
}

func (s *GetConsumerStackRequest) Validate() error {
	return dara.Validate(s)
}
