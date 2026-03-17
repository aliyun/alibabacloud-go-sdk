// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *GetApiKeyRequest
	GetApiKeyId() *string
}

type GetApiKeyRequest struct {
	ApiKeyId *string `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
}

func (s GetApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyRequest) GoString() string {
	return s.String()
}

func (s *GetApiKeyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *GetApiKeyRequest) SetApiKeyId(v string) *GetApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *GetApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
