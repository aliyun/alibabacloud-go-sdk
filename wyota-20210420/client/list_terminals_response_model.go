// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTerminalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTerminalsResponse
	GetStatusCode() *int32
	SetBody(v *ListTerminalsResponseBody) *ListTerminalsResponse
	GetBody() *ListTerminalsResponseBody
}

type ListTerminalsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalsResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTerminalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTerminalsResponse) GetBody() *ListTerminalsResponseBody {
	return s.Body
}

func (s *ListTerminalsResponse) SetHeaders(v map[string]*string) *ListTerminalsResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalsResponse) SetStatusCode(v int32) *ListTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalsResponse) SetBody(v *ListTerminalsResponseBody) *ListTerminalsResponse {
	s.Body = v
	return s
}

func (s *ListTerminalsResponse) Validate() error {
	return dara.Validate(s)
}
