// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairAIGatewayInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTairAIGatewayInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTairAIGatewayInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateTairAIGatewayInstanceResponseBody) *CreateTairAIGatewayInstanceResponse
	GetBody() *CreateTairAIGatewayInstanceResponseBody
}

type CreateTairAIGatewayInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTairAIGatewayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTairAIGatewayInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTairAIGatewayInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateTairAIGatewayInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTairAIGatewayInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTairAIGatewayInstanceResponse) GetBody() *CreateTairAIGatewayInstanceResponseBody {
	return s.Body
}

func (s *CreateTairAIGatewayInstanceResponse) SetHeaders(v map[string]*string) *CreateTairAIGatewayInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateTairAIGatewayInstanceResponse) SetStatusCode(v int32) *CreateTairAIGatewayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTairAIGatewayInstanceResponse) SetBody(v *CreateTairAIGatewayInstanceResponseBody) *CreateTairAIGatewayInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateTairAIGatewayInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
