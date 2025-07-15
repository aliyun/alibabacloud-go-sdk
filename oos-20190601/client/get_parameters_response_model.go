// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParametersResponse
	GetStatusCode() *int32
	SetBody(v *GetParametersResponseBody) *GetParametersResponse
	GetBody() *GetParametersResponseBody
}

type GetParametersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParametersResponse) GoString() string {
	return s.String()
}

func (s *GetParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParametersResponse) GetBody() *GetParametersResponseBody {
	return s.Body
}

func (s *GetParametersResponse) SetHeaders(v map[string]*string) *GetParametersResponse {
	s.Headers = v
	return s
}

func (s *GetParametersResponse) SetStatusCode(v int32) *GetParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersResponse) SetBody(v *GetParametersResponseBody) *GetParametersResponse {
	s.Body = v
	return s
}

func (s *GetParametersResponse) Validate() error {
	return dara.Validate(s)
}
