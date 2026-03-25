// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *UpdateApiKeyRequest
	GetApiKeyId() *int64
	SetDescription(v string) *UpdateApiKeyRequest
	GetDescription() *string
}

type UpdateApiKeyRequest struct {
	// API Key ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3303332
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *UpdateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiKeyRequest) SetApiKeyId(v int64) *UpdateApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateApiKeyRequest) SetDescription(v string) *UpdateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
