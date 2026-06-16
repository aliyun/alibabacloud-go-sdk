// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreAPIKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateContextStoreAPIKeyResponseBody
	GetApiKey() *string
	SetName(v string) *CreateContextStoreAPIKeyResponseBody
	GetName() *string
	SetRequestId(v string) *CreateContextStoreAPIKeyResponseBody
	GetRequestId() *string
}

type CreateContextStoreAPIKeyResponseBody struct {
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// my-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContextStoreAPIKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreAPIKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextStoreAPIKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateContextStoreAPIKeyResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateContextStoreAPIKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextStoreAPIKeyResponseBody) SetApiKey(v string) *CreateContextStoreAPIKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateContextStoreAPIKeyResponseBody) SetName(v string) *CreateContextStoreAPIKeyResponseBody {
	s.Name = &v
	return s
}

func (s *CreateContextStoreAPIKeyResponseBody) SetRequestId(v string) *CreateContextStoreAPIKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextStoreAPIKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
