// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelOperatorApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateModelOperatorApiKeyRequest
	GetDescription() *string
}

type CreateModelOperatorApiKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-apikey
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateModelOperatorApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelOperatorApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateModelOperatorApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelOperatorApiKeyRequest) SetDescription(v string) *CreateModelOperatorApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateModelOperatorApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
