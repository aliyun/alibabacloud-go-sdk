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
	SetIpBlacklist(v []*IPConfig) *UpdateApiKeyResponseBody
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *UpdateApiKeyResponseBody
	GetIpWhitelist() []*IPConfig
	SetMessage(v string) *UpdateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApiKeyResponseBody
	GetRequestId() *string
}

type UpdateApiKeyResponseBody struct {
	// API Key。
	//
	// example:
	//
	// asdfjoY87-9IUHH
	ApiKey *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The response status code.
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
	// 2BCFAE0A-9FA9-5F72-8E8B-724632BC19A9
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

func (s *UpdateApiKeyResponseBody) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *UpdateApiKeyResponseBody) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
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

func (s *UpdateApiKeyResponseBody) SetIpBlacklist(v []*IPConfig) *UpdateApiKeyResponseBody {
	s.IpBlacklist = v
	return s
}

func (s *UpdateApiKeyResponseBody) SetIpWhitelist(v []*IPConfig) *UpdateApiKeyResponseBody {
	s.IpWhitelist = v
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
