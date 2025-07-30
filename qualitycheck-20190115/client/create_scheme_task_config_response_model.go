// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchemeTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchemeTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchemeTaskConfigResponseBody) *CreateSchemeTaskConfigResponse
	GetBody() *CreateSchemeTaskConfigResponseBody
}

type CreateSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchemeTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemeTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchemeTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchemeTaskConfigResponse) GetBody() *CreateSchemeTaskConfigResponseBody {
	return s.Body
}

func (s *CreateSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *CreateSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemeTaskConfigResponse) SetStatusCode(v int32) *CreateSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemeTaskConfigResponse) SetBody(v *CreateSchemeTaskConfigResponseBody) *CreateSchemeTaskConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSchemeTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
