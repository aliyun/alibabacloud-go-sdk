// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindVariableResponse
	GetStatusCode() *int32
	SetBody(v *BindVariableResponseBody) *BindVariableResponse
	GetBody() *BindVariableResponseBody
}

type BindVariableResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s BindVariableResponse) GoString() string {
	return s.String()
}

func (s *BindVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindVariableResponse) GetBody() *BindVariableResponseBody {
	return s.Body
}

func (s *BindVariableResponse) SetHeaders(v map[string]*string) *BindVariableResponse {
	s.Headers = v
	return s
}

func (s *BindVariableResponse) SetStatusCode(v int32) *BindVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *BindVariableResponse) SetBody(v *BindVariableResponseBody) *BindVariableResponse {
	s.Body = v
	return s
}

func (s *BindVariableResponse) Validate() error {
	return dara.Validate(s)
}
