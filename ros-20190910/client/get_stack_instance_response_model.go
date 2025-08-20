// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetStackInstanceResponseBody) *GetStackInstanceResponse
	GetBody() *GetStackInstanceResponseBody
}

type GetStackInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackInstanceResponse) GetBody() *GetStackInstanceResponseBody {
	return s.Body
}

func (s *GetStackInstanceResponse) SetHeaders(v map[string]*string) *GetStackInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetStackInstanceResponse) SetStatusCode(v int32) *GetStackInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackInstanceResponse) SetBody(v *GetStackInstanceResponseBody) *GetStackInstanceResponse {
	s.Body = v
	return s
}

func (s *GetStackInstanceResponse) Validate() error {
	return dara.Validate(s)
}
