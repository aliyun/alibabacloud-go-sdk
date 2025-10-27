// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaliciousNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaliciousNoteResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaliciousNoteResponseBody) *CreateMaliciousNoteResponse
	GetBody() *CreateMaliciousNoteResponseBody
}

type CreateMaliciousNoteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaliciousNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaliciousNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousNoteResponse) GoString() string {
	return s.String()
}

func (s *CreateMaliciousNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaliciousNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaliciousNoteResponse) GetBody() *CreateMaliciousNoteResponseBody {
	return s.Body
}

func (s *CreateMaliciousNoteResponse) SetHeaders(v map[string]*string) *CreateMaliciousNoteResponse {
	s.Headers = v
	return s
}

func (s *CreateMaliciousNoteResponse) SetStatusCode(v int32) *CreateMaliciousNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaliciousNoteResponse) SetBody(v *CreateMaliciousNoteResponseBody) *CreateMaliciousNoteResponse {
	s.Body = v
	return s
}

func (s *CreateMaliciousNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
