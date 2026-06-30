// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeModelOperatorApiKeyResponseBody
	GetApiKey() *string
	SetApiKeyId(v int32) *DescribeModelOperatorApiKeyResponseBody
	GetApiKeyId() *int32
	SetCreateTime(v string) *DescribeModelOperatorApiKeyResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeModelOperatorApiKeyResponseBody
	GetDescription() *string
	SetEndpoint(v string) *DescribeModelOperatorApiKeyResponseBody
	GetEndpoint() *string
	SetRequestId(v string) *DescribeModelOperatorApiKeyResponseBody
	GetRequestId() *string
}

type DescribeModelOperatorApiKeyResponseBody struct {
	// example:
	//
	// QEDGOTAJOG
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 1
	ApiKeyId *int32 `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// 2026-06-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test-apikey
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://xxxx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModelOperatorApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetApiKeyId() *int32 {
	return s.ApiKeyId
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeModelOperatorApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetApiKey(v string) *DescribeModelOperatorApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetApiKeyId(v int32) *DescribeModelOperatorApiKeyResponseBody {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetCreateTime(v string) *DescribeModelOperatorApiKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetDescription(v string) *DescribeModelOperatorApiKeyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetEndpoint(v string) *DescribeModelOperatorApiKeyResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) SetRequestId(v string) *DescribeModelOperatorApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelOperatorApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
