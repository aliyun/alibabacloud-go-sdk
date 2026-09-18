// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *VerifyConnectorRequestBody) *VerifyConnectorRequest
	GetBody() *VerifyConnectorRequestBody
}

type VerifyConnectorRequest struct {
	// The validation request body.
	//
	// This parameter is required.
	Body *VerifyConnectorRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s VerifyConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorRequest) GoString() string {
	return s.String()
}

func (s *VerifyConnectorRequest) GetBody() *VerifyConnectorRequestBody {
	return s.Body
}

func (s *VerifyConnectorRequest) SetBody(v *VerifyConnectorRequestBody) *VerifyConnectorRequest {
	s.Body = v
	return s
}

func (s *VerifyConnectorRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyConnectorRequestBody struct {
	// The Connector configuration JSON string to validate. Set site to global or cn. The serviceAccountKeys field must contain at least one item with a serviceAccountKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"site":"global","organizationId":"org-xxxx","apiKey":"ak-xxxx","serviceAccountKeys":[{"name":"default","serviceAccountKey":"sk-xxxx"}]}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s VerifyConnectorRequestBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorRequestBody) GoString() string {
	return s.String()
}

func (s *VerifyConnectorRequestBody) GetMetadata() *string {
	return s.Metadata
}

func (s *VerifyConnectorRequestBody) SetMetadata(v string) *VerifyConnectorRequestBody {
	s.Metadata = &v
	return s
}

func (s *VerifyConnectorRequestBody) Validate() error {
	return dara.Validate(s)
}
