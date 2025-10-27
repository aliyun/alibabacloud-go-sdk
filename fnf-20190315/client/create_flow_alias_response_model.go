// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowAliasResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowAliasResponseBody) *CreateFlowAliasResponse
	GetBody() *CreateFlowAliasResponseBody
}

type CreateFlowAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowAliasResponse) GetBody() *CreateFlowAliasResponseBody {
	return s.Body
}

func (s *CreateFlowAliasResponse) SetHeaders(v map[string]*string) *CreateFlowAliasResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowAliasResponse) SetStatusCode(v int32) *CreateFlowAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowAliasResponse) SetBody(v *CreateFlowAliasResponseBody) *CreateFlowAliasResponse {
	s.Body = v
	return s
}

func (s *CreateFlowAliasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
