// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntentionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntentionsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntentionsResponseBody) *ListIntentionsResponse
	GetBody() *ListIntentionsResponseBody
}

type ListIntentionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntentionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntentionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponse) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntentionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntentionsResponse) GetBody() *ListIntentionsResponseBody {
	return s.Body
}

func (s *ListIntentionsResponse) SetHeaders(v map[string]*string) *ListIntentionsResponse {
	s.Headers = v
	return s
}

func (s *ListIntentionsResponse) SetStatusCode(v int32) *ListIntentionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentionsResponse) SetBody(v *ListIntentionsResponseBody) *ListIntentionsResponse {
	s.Body = v
	return s
}

func (s *ListIntentionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
