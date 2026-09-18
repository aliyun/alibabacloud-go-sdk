// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateConnectorRequestBody) *UpdateConnectorRequest
	GetBody() *UpdateConnectorRequestBody
}

type UpdateConnectorRequest struct {
	// The update request body.
	//
	// This parameter is required.
	Body *UpdateConnectorRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequest) GetBody() *UpdateConnectorRequestBody {
	return s.Body
}

func (s *UpdateConnectorRequest) SetBody(v *UpdateConnectorRequestBody) *UpdateConnectorRequest {
	s.Body = v
	return s
}

func (s *UpdateConnectorRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConnectorRequestBody struct {
	// The Connector configuration JSON string. The site value must match the value specified when the Connector was enabled. The organizationId value, if provided, must match the value specified when the Connector was enabled. If apiKey is omitted, the original value is retained. If serviceAccountKeys is provided, it represents the complete updated key collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"site":"global","organizationId":"org-xxxx"}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s UpdateConnectorRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectorRequestBody) GetMetadata() *string {
	return s.Metadata
}

func (s *UpdateConnectorRequestBody) SetMetadata(v string) *UpdateConnectorRequestBody {
	s.Metadata = &v
	return s
}

func (s *UpdateConnectorRequestBody) Validate() error {
	return dara.Validate(s)
}
