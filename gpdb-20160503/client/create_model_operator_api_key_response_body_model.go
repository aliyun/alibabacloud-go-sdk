// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelOperatorApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int32) *CreateModelOperatorApiKeyResponseBody
	GetApiKeyId() *int32
	SetRequestId(v string) *CreateModelOperatorApiKeyResponseBody
	GetRequestId() *string
}

type CreateModelOperatorApiKeyResponseBody struct {
	// example:
	//
	// 1
	ApiKeyId *int32 `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelOperatorApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelOperatorApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelOperatorApiKeyResponseBody) GetApiKeyId() *int32 {
	return s.ApiKeyId
}

func (s *CreateModelOperatorApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelOperatorApiKeyResponseBody) SetApiKeyId(v int32) *CreateModelOperatorApiKeyResponseBody {
	s.ApiKeyId = &v
	return s
}

func (s *CreateModelOperatorApiKeyResponseBody) SetRequestId(v string) *CreateModelOperatorApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelOperatorApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
