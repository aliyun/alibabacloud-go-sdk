// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserIntentionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserIntentionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserIntentionsResponseBody) *ListUserIntentionsResponse
	GetBody() *ListUserIntentionsResponseBody
}

type ListUserIntentionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserIntentionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserIntentionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserIntentionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserIntentionsResponse) GetBody() *ListUserIntentionsResponseBody {
	return s.Body
}

func (s *ListUserIntentionsResponse) SetHeaders(v map[string]*string) *ListUserIntentionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserIntentionsResponse) SetStatusCode(v int32) *ListUserIntentionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserIntentionsResponse) SetBody(v *ListUserIntentionsResponseBody) *ListUserIntentionsResponse {
	s.Body = v
	return s
}

func (s *ListUserIntentionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
