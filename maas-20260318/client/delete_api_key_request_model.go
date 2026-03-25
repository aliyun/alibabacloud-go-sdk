// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *DeleteApiKeyRequest
	GetApiKeyId() *int64
}

type DeleteApiKeyRequest struct {
	// API Key ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3952240
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
}

func (s DeleteApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *DeleteApiKeyRequest) SetApiKeyId(v int64) *DeleteApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *DeleteApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
