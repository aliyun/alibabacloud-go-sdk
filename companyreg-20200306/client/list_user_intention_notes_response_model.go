// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionNotesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserIntentionNotesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserIntentionNotesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserIntentionNotesResponseBody) *ListUserIntentionNotesResponse
	GetBody() *ListUserIntentionNotesResponseBody
}

type ListUserIntentionNotesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserIntentionNotesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserIntentionNotesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionNotesResponse) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserIntentionNotesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserIntentionNotesResponse) GetBody() *ListUserIntentionNotesResponseBody {
	return s.Body
}

func (s *ListUserIntentionNotesResponse) SetHeaders(v map[string]*string) *ListUserIntentionNotesResponse {
	s.Headers = v
	return s
}

func (s *ListUserIntentionNotesResponse) SetStatusCode(v int32) *ListUserIntentionNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserIntentionNotesResponse) SetBody(v *ListUserIntentionNotesResponseBody) *ListUserIntentionNotesResponse {
	s.Body = v
	return s
}

func (s *ListUserIntentionNotesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
