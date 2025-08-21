// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntentResponse
	GetStatusCode() *int32
	SetBody(v *ListIntentResponseBody) *ListIntentResponse
	GetBody() *ListIntentResponseBody
}

type ListIntentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntentResponse) GoString() string {
	return s.String()
}

func (s *ListIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntentResponse) GetBody() *ListIntentResponseBody {
	return s.Body
}

func (s *ListIntentResponse) SetHeaders(v map[string]*string) *ListIntentResponse {
	s.Headers = v
	return s
}

func (s *ListIntentResponse) SetStatusCode(v int32) *ListIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentResponse) SetBody(v *ListIntentResponseBody) *ListIntentResponse {
	s.Body = v
	return s
}

func (s *ListIntentResponse) Validate() error {
	return dara.Validate(s)
}
