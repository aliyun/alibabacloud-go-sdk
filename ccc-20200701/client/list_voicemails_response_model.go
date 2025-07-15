// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicemailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoicemailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoicemailsResponse
	GetStatusCode() *int32
	SetBody(v *ListVoicemailsResponseBody) *ListVoicemailsResponse
	GetBody() *ListVoicemailsResponseBody
}

type ListVoicemailsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoicemailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoicemailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoicemailsResponse) GoString() string {
	return s.String()
}

func (s *ListVoicemailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoicemailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoicemailsResponse) GetBody() *ListVoicemailsResponseBody {
	return s.Body
}

func (s *ListVoicemailsResponse) SetHeaders(v map[string]*string) *ListVoicemailsResponse {
	s.Headers = v
	return s
}

func (s *ListVoicemailsResponse) SetStatusCode(v int32) *ListVoicemailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoicemailsResponse) SetBody(v *ListVoicemailsResponseBody) *ListVoicemailsResponse {
	s.Body = v
	return s
}

func (s *ListVoicemailsResponse) Validate() error {
	return dara.Validate(s)
}
