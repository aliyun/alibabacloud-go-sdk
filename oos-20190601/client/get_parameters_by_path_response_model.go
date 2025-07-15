// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersByPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParametersByPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParametersByPathResponse
	GetStatusCode() *int32
	SetBody(v *GetParametersByPathResponseBody) *GetParametersByPathResponse
	GetBody() *GetParametersByPathResponseBody
}

type GetParametersByPathResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersByPathResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParametersByPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParametersByPathResponse) GetBody() *GetParametersByPathResponseBody {
	return s.Body
}

func (s *GetParametersByPathResponse) SetHeaders(v map[string]*string) *GetParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetParametersByPathResponse) SetStatusCode(v int32) *GetParametersByPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersByPathResponse) SetBody(v *GetParametersByPathResponseBody) *GetParametersByPathResponse {
	s.Body = v
	return s
}

func (s *GetParametersByPathResponse) Validate() error {
	return dara.Validate(s)
}
