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
	SetMessage(v string) *DescribeApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApiKeyResponseBody
	GetRequestId() *string
}

type DescribeApiKeyResponseBody struct {
	ApiKey    *ApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
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
	return nil
}
