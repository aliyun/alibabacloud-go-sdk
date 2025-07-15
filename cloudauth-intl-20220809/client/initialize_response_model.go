// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeResponse
	GetStatusCode() *int32
	SetBody(v *InitializeResponseBody) *InitializeResponse
	GetBody() *InitializeResponseBody
}

type InitializeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeResponse) GoString() string {
	return s.String()
}

func (s *InitializeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeResponse) GetBody() *InitializeResponseBody {
	return s.Body
}

func (s *InitializeResponse) SetHeaders(v map[string]*string) *InitializeResponse {
	s.Headers = v
	return s
}

func (s *InitializeResponse) SetStatusCode(v int32) *InitializeResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeResponse) SetBody(v *InitializeResponseBody) *InitializeResponse {
	s.Body = v
	return s
}

func (s *InitializeResponse) Validate() error {
	return dara.Validate(s)
}
