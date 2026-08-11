// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *ApiKey) *DescribeApiKeyResponseBody
	GetApiKey() *ApiKey
	SetCode(v string) *DescribeApiKeyResponseBody
	GetCode() *string
	SetIpBlacklist(v []*IPConfig) *DescribeApiKeyResponseBody
	GetIpBlacklist() []*IPConfig
	SetIpWhitelist(v []*IPConfig) *DescribeApiKeyResponseBody
	GetIpWhitelist() []*IPConfig
	SetMessage(v string) *DescribeApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApiKeyResponseBody
	GetRequestId() *string
}

type DescribeApiKeyResponseBody struct {
	// The Bailian API key.
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
	// 2BCFAE0A-9FA9-5F72-8E8B-724632BC19A9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiKeyResponseBody) GetApiKey() *ApiKey {
	return s.ApiKey
}

func (s *DescribeApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApiKeyResponseBody) GetIpBlacklist() []*IPConfig {
	return s.IpBlacklist
}

func (s *DescribeApiKeyResponseBody) GetIpWhitelist() []*IPConfig {
	return s.IpWhitelist
}

func (s *DescribeApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiKeyResponseBody) SetApiKey(v *ApiKey) *DescribeApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *DescribeApiKeyResponseBody) SetCode(v string) *DescribeApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApiKeyResponseBody) SetIpBlacklist(v []*IPConfig) *DescribeApiKeyResponseBody {
	s.IpBlacklist = v
	return s
}

func (s *DescribeApiKeyResponseBody) SetIpWhitelist(v []*IPConfig) *DescribeApiKeyResponseBody {
	s.IpWhitelist = v
	return s
}

func (s *DescribeApiKeyResponseBody) SetMessage(v string) *DescribeApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApiKeyResponseBody) SetRequestId(v string) *DescribeApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiKeyResponseBody) Validate() error {
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
