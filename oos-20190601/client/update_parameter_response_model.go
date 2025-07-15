// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateParameterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateParameterResponseBody) *UpdateParameterResponse
	GetBody() *UpdateParameterResponseBody
}

type UpdateParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateParameterResponse) GetBody() *UpdateParameterResponseBody {
	return s.Body
}

func (s *UpdateParameterResponse) SetHeaders(v map[string]*string) *UpdateParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateParameterResponse) SetStatusCode(v int32) *UpdateParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParameterResponse) SetBody(v *UpdateParameterResponseBody) *UpdateParameterResponse {
	s.Body = v
	return s
}

func (s *UpdateParameterResponse) Validate() error {
	return dara.Validate(s)
}
