// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIntentionNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIntentionNoteResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIntentionNoteResponseBody) *SubmitIntentionNoteResponse
	GetBody() *SubmitIntentionNoteResponseBody
}

type SubmitIntentionNoteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIntentionNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIntentionNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionNoteResponse) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIntentionNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIntentionNoteResponse) GetBody() *SubmitIntentionNoteResponseBody {
	return s.Body
}

func (s *SubmitIntentionNoteResponse) SetHeaders(v map[string]*string) *SubmitIntentionNoteResponse {
	s.Headers = v
	return s
}

func (s *SubmitIntentionNoteResponse) SetStatusCode(v int32) *SubmitIntentionNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIntentionNoteResponse) SetBody(v *SubmitIntentionNoteResponseBody) *SubmitIntentionNoteResponse {
	s.Body = v
	return s
}

func (s *SubmitIntentionNoteResponse) Validate() error {
	return dara.Validate(s)
}
