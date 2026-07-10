// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *ApiKey) *UpdateApiKeyResponseBody
	GetApiKey() *ApiKey
	SetCode(v string) *UpdateApiKeyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApiKeyResponseBody
	GetRequestId() *string
}

type UpdateApiKeyResponseBody struct {
	ApiKey    *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyResponseBody) GetApiKey() *ApiKey {
	return s.ApiKey
}

func (s *UpdateApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiKeyResponseBody) SetApiKey(v *ApiKey) *UpdateApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *UpdateApiKeyResponseBody) SetCode(v string) *UpdateApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetMessage(v string) *UpdateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApiKeyResponseBody) SetRequestId(v string) *UpdateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}
