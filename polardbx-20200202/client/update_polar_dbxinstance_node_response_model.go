// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarDBXInstanceNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarDBXInstanceNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarDBXInstanceNodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarDBXInstanceNodeResponseBody) *UpdatePolarDBXInstanceNodeResponse
	GetBody() *UpdatePolarDBXInstanceNodeResponseBody
}

type UpdatePolarDBXInstanceNodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarDBXInstanceNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarDBXInstanceNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarDBXInstanceNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarDBXInstanceNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarDBXInstanceNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarDBXInstanceNodeResponse) GetBody() *UpdatePolarDBXInstanceNodeResponseBody {
	return s.Body
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetHeaders(v map[string]*string) *UpdatePolarDBXInstanceNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetStatusCode(v int32) *UpdatePolarDBXInstanceNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponse) SetBody(v *UpdatePolarDBXInstanceNodeResponseBody) *UpdatePolarDBXInstanceNodeResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarDBXInstanceNodeResponse) Validate() error {
	return dara.Validate(s)
}
