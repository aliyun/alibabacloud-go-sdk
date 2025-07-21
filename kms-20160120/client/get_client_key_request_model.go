// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientKeyId(v string) *GetClientKeyRequest
	GetClientKeyId() *string
}

type GetClientKeyRequest struct {
	// The ID of the client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAAP.66abf237-63f6-4625-b8cf-47e1086e****
	ClientKeyId *string `json:"ClientKeyId,omitempty" xml:"ClientKeyId,omitempty"`
}

func (s GetClientKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientKeyRequest) GoString() string {
	return s.String()
}

func (s *GetClientKeyRequest) GetClientKeyId() *string {
	return s.ClientKeyId
}

func (s *GetClientKeyRequest) SetClientKeyId(v string) *GetClientKeyRequest {
	s.ClientKeyId = &v
	return s
}

func (s *GetClientKeyRequest) Validate() error {
	return dara.Validate(s)
}
