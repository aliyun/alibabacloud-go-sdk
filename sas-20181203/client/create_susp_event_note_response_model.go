// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuspEventNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSuspEventNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSuspEventNoteResponse
	GetStatusCode() *int32
	SetBody(v *CreateSuspEventNoteResponseBody) *CreateSuspEventNoteResponse
	GetBody() *CreateSuspEventNoteResponseBody
}

type CreateSuspEventNoteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSuspEventNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSuspEventNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSuspEventNoteResponse) GoString() string {
	return s.String()
}

func (s *CreateSuspEventNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSuspEventNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSuspEventNoteResponse) GetBody() *CreateSuspEventNoteResponseBody {
	return s.Body
}

func (s *CreateSuspEventNoteResponse) SetHeaders(v map[string]*string) *CreateSuspEventNoteResponse {
	s.Headers = v
	return s
}

func (s *CreateSuspEventNoteResponse) SetStatusCode(v int32) *CreateSuspEventNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSuspEventNoteResponse) SetBody(v *CreateSuspEventNoteResponseBody) *CreateSuspEventNoteResponse {
	s.Body = v
	return s
}

func (s *CreateSuspEventNoteResponse) Validate() error {
	return dara.Validate(s)
}
