// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCoachCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CoachCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CoachCallResponse
	GetStatusCode() *int32
	SetBody(v *CoachCallResponseBody) *CoachCallResponse
	GetBody() *CoachCallResponseBody
}

type CoachCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CoachCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CoachCallResponse) String() string {
	return dara.Prettify(s)
}

func (s CoachCallResponse) GoString() string {
	return s.String()
}

func (s *CoachCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CoachCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CoachCallResponse) GetBody() *CoachCallResponseBody {
	return s.Body
}

func (s *CoachCallResponse) SetHeaders(v map[string]*string) *CoachCallResponse {
	s.Headers = v
	return s
}

func (s *CoachCallResponse) SetStatusCode(v int32) *CoachCallResponse {
	s.StatusCode = &v
	return s
}

func (s *CoachCallResponse) SetBody(v *CoachCallResponseBody) *CoachCallResponse {
	s.Body = v
	return s
}

func (s *CoachCallResponse) Validate() error {
	return dara.Validate(s)
}
