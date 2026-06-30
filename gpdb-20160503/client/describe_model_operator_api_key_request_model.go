// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int32) *DescribeModelOperatorApiKeyRequest
	GetApiKeyId() *int32
}

type DescribeModelOperatorApiKeyRequest struct {
	// example:
	//
	// 1
	ApiKeyId *int32 `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s DescribeModelOperatorApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorApiKeyRequest) GetApiKeyId() *int32 {
	return s.ApiKeyId
}

func (s *DescribeModelOperatorApiKeyRequest) SetApiKeyId(v int32) *DescribeModelOperatorApiKeyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeModelOperatorApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
