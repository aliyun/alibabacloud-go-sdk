// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateTokenPlanKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *RotateTokenPlanKeyRequest
	GetApiKeyId() *string
}

type RotateTokenPlanKeyRequest struct {
	// API Key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ak_123456
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s RotateTokenPlanKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s RotateTokenPlanKeyRequest) GoString() string {
	return s.String()
}

func (s *RotateTokenPlanKeyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *RotateTokenPlanKeyRequest) SetApiKeyId(v string) *RotateTokenPlanKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *RotateTokenPlanKeyRequest) Validate() error {
	return dara.Validate(s)
}
