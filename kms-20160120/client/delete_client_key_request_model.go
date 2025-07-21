// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientKeyId(v string) *DeleteClientKeyRequest
	GetClientKeyId() *string
}

type DeleteClientKeyRequest struct {
	// The ID of the client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
}

func (s DeleteClientKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyRequest) GetClientKeyId() *string {
	return s.ClientKeyId
}

func (s *DeleteClientKeyRequest) SetClientKeyId(v string) *DeleteClientKeyRequest {
	s.ClientKeyId = &v
	return s
}

func (s *DeleteClientKeyRequest) Validate() error {
	return dara.Validate(s)
}
