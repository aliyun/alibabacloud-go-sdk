// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ModelRouterCreateApiKeyRequest
	GetClientId() *int64
}

type ModelRouterCreateApiKeyRequest struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s ModelRouterCreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateApiKeyRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterCreateApiKeyRequest) SetClientId(v int64) *ModelRouterCreateApiKeyRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterCreateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
