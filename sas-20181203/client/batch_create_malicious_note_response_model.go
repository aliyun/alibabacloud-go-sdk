// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMaliciousNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateMaliciousNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateMaliciousNoteResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateMaliciousNoteResponseBody) *BatchCreateMaliciousNoteResponse
	GetBody() *BatchCreateMaliciousNoteResponseBody
}

type BatchCreateMaliciousNoteResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateMaliciousNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateMaliciousNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMaliciousNoteResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateMaliciousNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateMaliciousNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateMaliciousNoteResponse) GetBody() *BatchCreateMaliciousNoteResponseBody {
	return s.Body
}

func (s *BatchCreateMaliciousNoteResponse) SetHeaders(v map[string]*string) *BatchCreateMaliciousNoteResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateMaliciousNoteResponse) SetStatusCode(v int32) *BatchCreateMaliciousNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateMaliciousNoteResponse) SetBody(v *BatchCreateMaliciousNoteResponseBody) *BatchCreateMaliciousNoteResponse {
	s.Body = v
	return s
}

func (s *BatchCreateMaliciousNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
