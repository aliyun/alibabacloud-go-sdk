// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateModelRequestBody) *UpdateModelRequest
	GetBody() *UpdateModelRequestBody
	SetClientToken(v string) *UpdateModelRequest
	GetClientToken() *string
}

type UpdateModelRequest struct {
	Body *UpdateModelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) GetBody() *UpdateModelRequestBody {
	return s.Body
}

func (s *UpdateModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelRequest) SetBody(v *UpdateModelRequestBody) *UpdateModelRequest {
	s.Body = v
	return s
}

func (s *UpdateModelRequest) SetClientToken(v string) *UpdateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateModelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateModelRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelRequestBody) SetDescription(v string) *UpdateModelRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateModelRequestBody) Validate() error {
	return dara.Validate(s)
}
