// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchemeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchemeConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchemeConfigResponseBody) *CreateSchemeConfigResponse
	GetBody() *CreateSchemeConfigResponseBody
}

type CreateSchemeConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchemeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchemeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchemeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchemeConfigResponse) GetBody() *CreateSchemeConfigResponseBody {
	return s.Body
}

func (s *CreateSchemeConfigResponse) SetHeaders(v map[string]*string) *CreateSchemeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemeConfigResponse) SetStatusCode(v int32) *CreateSchemeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemeConfigResponse) SetBody(v *CreateSchemeConfigResponseBody) *CreateSchemeConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSchemeConfigResponse) Validate() error {
	return dara.Validate(s)
}
