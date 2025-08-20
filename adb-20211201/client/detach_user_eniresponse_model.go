// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachUserENIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachUserENIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachUserENIResponse
	GetStatusCode() *int32
	SetBody(v *DetachUserENIResponseBody) *DetachUserENIResponse
	GetBody() *DetachUserENIResponseBody
}

type DetachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachUserENIResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachUserENIResponse) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachUserENIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachUserENIResponse) GetBody() *DetachUserENIResponseBody {
	return s.Body
}

func (s *DetachUserENIResponse) SetHeaders(v map[string]*string) *DetachUserENIResponse {
	s.Headers = v
	return s
}

func (s *DetachUserENIResponse) SetStatusCode(v int32) *DetachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachUserENIResponse) SetBody(v *DetachUserENIResponseBody) *DetachUserENIResponse {
	s.Body = v
	return s
}

func (s *DetachUserENIResponse) Validate() error {
	return dara.Validate(s)
}
