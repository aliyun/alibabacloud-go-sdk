// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *GetApiKeyRequest
	GetApiKeyId() *int64
}

type GetApiKeyRequest struct {
	// API Key ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3303332
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
}

func (s GetApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyRequest) GoString() string {
	return s.String()
}

func (s *GetApiKeyRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *GetApiKeyRequest) SetApiKeyId(v int64) *GetApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *GetApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
