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
	// The value of the API key. This value is returned only upon creation. Store it in a secure location.
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The display name of the API key.
	//
	// example:
	//
	// Production Service Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3B311FD9-A60B-55E0-A896-A0C73*********
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
