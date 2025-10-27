// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowAliasResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowAliasResponseBody) *UpdateFlowAliasResponse
	GetBody() *UpdateFlowAliasResponseBody
}

type UpdateFlowAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowAliasResponse) GetBody() *UpdateFlowAliasResponseBody {
	return s.Body
}

func (s *UpdateFlowAliasResponse) SetHeaders(v map[string]*string) *UpdateFlowAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowAliasResponse) SetStatusCode(v int32) *UpdateFlowAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowAliasResponse) SetBody(v *UpdateFlowAliasResponseBody) *UpdateFlowAliasResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
