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
	SetIpBlacklist(v []*IPConfig) *ResetApiKeyResponseBody
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *ResetApiKeyResponseBody
	GetIpWhitelist() []*IPConfig
	SetMessage(v string) *ResetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetApiKeyResponseBody
	GetRequestId() *string
}

type ResetApiKeyResponseBody struct {
	// The API key information.
	//
	// example:
	//
	// asdfjoY87-9IUHH
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
	// The request ID.
	//
	// example:
	//
	// 7ADFF8D8-D4BA-5F79-AD49-DDABFEA59B6C
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

func (s *ResetApiKeyResponseBody) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *ResetApiKeyResponseBody) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
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

func (s *ResetApiKeyResponseBody) SetIpBlacklist(v []*IPConfig) *ResetApiKeyResponseBody {
	s.IpBlacklist = v
	return s
}

func (s *ResetApiKeyResponseBody) SetIpWhitelist(v []*IPConfig) *ResetApiKeyResponseBody {
	s.IpWhitelist = v
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
