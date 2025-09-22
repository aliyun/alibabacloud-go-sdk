// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermEditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TermEditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TermEditResponse
	GetStatusCode() *int32
	SetBody(v *TermEditResponseBody) *TermEditResponse
	GetBody() *TermEditResponseBody
}

type TermEditResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TermEditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TermEditResponse) String() string {
	return dara.Prettify(s)
}

func (s TermEditResponse) GoString() string {
	return s.String()
}

func (s *TermEditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TermEditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TermEditResponse) GetBody() *TermEditResponseBody {
	return s.Body
}

func (s *TermEditResponse) SetHeaders(v map[string]*string) *TermEditResponse {
	s.Headers = v
	return s
}

func (s *TermEditResponse) SetStatusCode(v int32) *TermEditResponse {
	s.StatusCode = &v
	return s
}

func (s *TermEditResponse) SetBody(v *TermEditResponseBody) *TermEditResponse {
	s.Body = v
	return s
}

func (s *TermEditResponse) Validate() error {
	return dara.Validate(s)
}
