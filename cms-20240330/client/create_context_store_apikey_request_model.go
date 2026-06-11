// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreAPIKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateContextStoreAPIKeyRequest
	GetName() *string
}

type CreateContextStoreAPIKeyRequest struct {
	// The display name of the API key. Use this name to identify its purpose.
	//
	// This parameter is required.
	//
	// example:
	//
	// Production Service Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateContextStoreAPIKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreAPIKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateContextStoreAPIKeyRequest) GetName() *string {
	return s.Name
}

func (s *CreateContextStoreAPIKeyRequest) SetName(v string) *CreateContextStoreAPIKeyRequest {
	s.Name = &v
	return s
}

func (s *CreateContextStoreAPIKeyRequest) Validate() error {
	return dara.Validate(s)
}
