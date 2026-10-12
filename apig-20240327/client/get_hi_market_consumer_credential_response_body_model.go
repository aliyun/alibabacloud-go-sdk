// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHiMarketConsumerCredentialResponseBody
	GetCode() *string
	SetData(v *GetHiMarketConsumerCredentialResponseBodyData) *GetHiMarketConsumerCredentialResponseBody
	GetData() *GetHiMarketConsumerCredentialResponseBodyData
	SetMessage(v string) *GetHiMarketConsumerCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHiMarketConsumerCredentialResponseBody
	GetRequestId() *string
}

type GetHiMarketConsumerCredentialResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetHiMarketConsumerCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHiMarketConsumerCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHiMarketConsumerCredentialResponseBody) GetData() *GetHiMarketConsumerCredentialResponseBodyData {
	return s.Data
}

func (s *GetHiMarketConsumerCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHiMarketConsumerCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHiMarketConsumerCredentialResponseBody) SetCode(v string) *GetHiMarketConsumerCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBody) SetData(v *GetHiMarketConsumerCredentialResponseBodyData) *GetHiMarketConsumerCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBody) SetMessage(v string) *GetHiMarketConsumerCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBody) SetRequestId(v string) *GetHiMarketConsumerCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHiMarketConsumerCredentialResponseBodyData struct {
	// The API key credential configuration of the consumer. This corresponds to the common data structure HiMarketApiKeyConfig and includes the key, source, and credentials list members. Each item in the credentials list contains an apiKey.
	ApiKeyConfig *HiMarketApiKeyConfig `json:"apiKeyConfig,omitempty" xml:"apiKeyConfig,omitempty"`
	// The HMAC credential configuration of the consumer. This corresponds to the common data structure HiMarketHmacConfig and includes the credentials list member. Each item in the credentials list contains an ak and an sk.
	HmacConfig *HiMarketHmacConfig `json:"hmacConfig,omitempty" xml:"hmacConfig,omitempty"`
}

func (s GetHiMarketConsumerCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerCredentialResponseBodyData) GetApiKeyConfig() *HiMarketApiKeyConfig {
	return s.ApiKeyConfig
}

func (s *GetHiMarketConsumerCredentialResponseBodyData) GetHmacConfig() *HiMarketHmacConfig {
	return s.HmacConfig
}

func (s *GetHiMarketConsumerCredentialResponseBodyData) SetApiKeyConfig(v *HiMarketApiKeyConfig) *GetHiMarketConsumerCredentialResponseBodyData {
	s.ApiKeyConfig = v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBodyData) SetHmacConfig(v *HiMarketHmacConfig) *GetHiMarketConsumerCredentialResponseBodyData {
	s.HmacConfig = v
	return s
}

func (s *GetHiMarketConsumerCredentialResponseBodyData) Validate() error {
	if s.ApiKeyConfig != nil {
		if err := s.ApiKeyConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HmacConfig != nil {
		if err := s.HmacConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
