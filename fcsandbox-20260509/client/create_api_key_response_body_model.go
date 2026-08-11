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
	// API Key。
	//
	// example:
	//
	// {\\"gmtCreate\\": 1776997128000, \\"workspaceId\\": \\"ws-3w77kird5sblwwfk\\", \\"apiKeyId\\": 4808780, \\"createdBy\\": \\"AssumedRoleUser300873166069492100\\", \\"apiKeyValue\\": \\"sk-110ff0028ad64a24b9453c8955c4c191\\"}
	ApiKey *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code        *string     `json:"code,omitempty" xml:"code,omitempty"`
	IpBlacklist []*IPConfig `json:"ipBlacklist,omitempty" xml:"ipBlacklist,omitempty" type:"Repeated"`
	IpWhitelist []*IPConfig `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
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
