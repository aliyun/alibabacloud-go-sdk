// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParameterResponse
	GetStatusCode() *int32
	SetBody(v *GetParameterResponseBody) *GetParameterResponse
	GetBody() *GetParameterResponseBody
}

type GetParameterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponse) GoString() string {
	return s.String()
}

func (s *GetParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParameterResponse) GetBody() *GetParameterResponseBody {
	return s.Body
}

func (s *GetParameterResponse) SetHeaders(v map[string]*string) *GetParameterResponse {
	s.Headers = v
	return s
}

func (s *GetParameterResponse) SetStatusCode(v int32) *GetParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParameterResponse) SetBody(v *GetParameterResponseBody) *GetParameterResponse {
	s.Body = v
	return s
}

func (s *GetParameterResponse) Validate() error {
	return dara.Validate(s)
}
