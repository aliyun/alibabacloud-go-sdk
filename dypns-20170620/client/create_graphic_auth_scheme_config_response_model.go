// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphicAuthSchemeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGraphicAuthSchemeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGraphicAuthSchemeConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateGraphicAuthSchemeConfigResponseBody) *CreateGraphicAuthSchemeConfigResponse
	GetBody() *CreateGraphicAuthSchemeConfigResponseBody
}

type CreateGraphicAuthSchemeConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGraphicAuthSchemeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGraphicAuthSchemeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphicAuthSchemeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateGraphicAuthSchemeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGraphicAuthSchemeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGraphicAuthSchemeConfigResponse) GetBody() *CreateGraphicAuthSchemeConfigResponseBody {
	return s.Body
}

func (s *CreateGraphicAuthSchemeConfigResponse) SetHeaders(v map[string]*string) *CreateGraphicAuthSchemeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponse) SetStatusCode(v int32) *CreateGraphicAuthSchemeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponse) SetBody(v *CreateGraphicAuthSchemeConfigResponseBody) *CreateGraphicAuthSchemeConfigResponse {
	s.Body = v
	return s
}

func (s *CreateGraphicAuthSchemeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
