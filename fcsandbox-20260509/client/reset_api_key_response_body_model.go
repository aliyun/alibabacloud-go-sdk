// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *ApiKey) *ResetApiKeyResponseBody
	GetApiKey() *ApiKey
	SetCode(v string) *ResetApiKeyResponseBody
	GetCode() *string
	SetMessage(v string) *ResetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetApiKeyResponseBody
	GetRequestId() *string
}

type ResetApiKeyResponseBody struct {
	ApiKey    *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ResetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBody) GetApiKey() *ApiKey {
	return s.ApiKey
}

func (s *ResetApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetApiKeyResponseBody) SetApiKey(v *ApiKey) *ResetApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *ResetApiKeyResponseBody) SetCode(v string) *ResetApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetMessage(v string) *ResetApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetRequestId(v string) *ResetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}
