// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntentionNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntentionNoteResponse
	GetStatusCode() *int32
	SetBody(v *ListIntentionNoteResponseBody) *ListIntentionNoteResponse
	GetBody() *ListIntentionNoteResponseBody
}

type ListIntentionNoteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntentionNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntentionNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionNoteResponse) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntentionNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntentionNoteResponse) GetBody() *ListIntentionNoteResponseBody {
	return s.Body
}

func (s *ListIntentionNoteResponse) SetHeaders(v map[string]*string) *ListIntentionNoteResponse {
	s.Headers = v
	return s
}

func (s *ListIntentionNoteResponse) SetStatusCode(v int32) *ListIntentionNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentionNoteResponse) SetBody(v *ListIntentionNoteResponseBody) *ListIntentionNoteResponse {
	s.Body = v
	return s
}

func (s *ListIntentionNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
