// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSymbolicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSymbolicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSymbolicResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSymbolicResponseBody) *SubmitSymbolicResponse
	GetBody() *SubmitSymbolicResponseBody
}

type SubmitSymbolicResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSymbolicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSymbolicResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSymbolicResponse) GoString() string {
	return s.String()
}

func (s *SubmitSymbolicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSymbolicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSymbolicResponse) GetBody() *SubmitSymbolicResponseBody {
	return s.Body
}

func (s *SubmitSymbolicResponse) SetHeaders(v map[string]*string) *SubmitSymbolicResponse {
	s.Headers = v
	return s
}

func (s *SubmitSymbolicResponse) SetStatusCode(v int32) *SubmitSymbolicResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSymbolicResponse) SetBody(v *SubmitSymbolicResponseBody) *SubmitSymbolicResponse {
	s.Body = v
	return s
}

func (s *SubmitSymbolicResponse) Validate() error {
	return dara.Validate(s)
}
