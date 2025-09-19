// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaliciousNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaliciousNoteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaliciousNoteResponseBody) *DeleteMaliciousNoteResponse
	GetBody() *DeleteMaliciousNoteResponseBody
}

type DeleteMaliciousNoteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaliciousNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaliciousNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousNoteResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaliciousNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaliciousNoteResponse) GetBody() *DeleteMaliciousNoteResponseBody {
	return s.Body
}

func (s *DeleteMaliciousNoteResponse) SetHeaders(v map[string]*string) *DeleteMaliciousNoteResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaliciousNoteResponse) SetStatusCode(v int32) *DeleteMaliciousNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaliciousNoteResponse) SetBody(v *DeleteMaliciousNoteResponseBody) *DeleteMaliciousNoteResponse {
	s.Body = v
	return s
}

func (s *DeleteMaliciousNoteResponse) Validate() error {
	return dara.Validate(s)
}
