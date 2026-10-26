// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *ApiKey) *CreateApiKeyResponseBody
	GetApiKey() *ApiKey
	SetCode(v string) *CreateApiKeyResponseBody
	GetCode() *string
	SetIpBlacklist(v []*IPConfig) *CreateApiKeyResponseBody
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *CreateApiKeyResponseBody
	GetIpWhitelist() []*IPConfig
	SetMessage(v string) *CreateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApiKeyResponseBody
	GetRequestId() *string
}

type CreateApiKeyResponseBody struct {
	// The API key information.
	//
	// example:
	//
	// [Replace with a masked API key example]
	ApiKey *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of IP blacklist rules.
	IpBlacklist []*IPConfig `json:"ipBlacklist" xml:"ipBlacklist" type:"Repeated"`
	// The list of IP address whitelist rules.
	IpWhitelist []*IPConfig `json:"ipWhitelist" xml:"ipWhitelist" type:"Repeated"`
	// The response message. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBody) GetApiKey() *ApiKey {
	return s.ApiKey
}

func (s *CreateApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApiKeyResponseBody) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *CreateApiKeyResponseBody) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
}

func (s *CreateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiKeyResponseBody) SetApiKey(v *ApiKey) *CreateApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *CreateApiKeyResponseBody) SetCode(v string) *CreateApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetIpBlacklist(v []*IPConfig) *CreateApiKeyResponseBody {
	s.IpBlacklist = v
	return s
}

func (s *CreateApiKeyResponseBody) SetIpWhitelist(v []*IPConfig) *CreateApiKeyResponseBody {
	s.IpWhitelist = v
	return s
}

func (s *CreateApiKeyResponseBody) SetMessage(v string) *CreateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetRequestId(v string) *CreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	if s.IpBlacklist != nil {
		for _, item := range s.IpBlacklist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpWhitelist != nil {
		for _, item := range s.IpWhitelist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
