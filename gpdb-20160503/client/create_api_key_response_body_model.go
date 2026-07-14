// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateApiKeyResponseBody
	GetApiKey() *string
	SetKeyId(v string) *CreateApiKeyResponseBody
	GetKeyId() *string
	SetRequestId(v string) *CreateApiKeyResponseBody
	GetRequestId() *string
}

type CreateApiKeyResponseBody struct {
	// example:
	//
	// sk-xxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// api-xxxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateApiKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiKeyResponseBody) SetApiKey(v string) *CreateApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetKeyId(v string) *CreateApiKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetRequestId(v string) *CreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
