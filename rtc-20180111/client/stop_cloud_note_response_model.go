// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCloudNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCloudNoteResponse
	GetStatusCode() *int32
	SetBody(v *StopCloudNoteResponseBody) *StopCloudNoteResponse
	GetBody() *StopCloudNoteResponseBody
}

type StopCloudNoteResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCloudNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCloudNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCloudNoteResponse) GoString() string {
	return s.String()
}

func (s *StopCloudNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCloudNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCloudNoteResponse) GetBody() *StopCloudNoteResponseBody {
	return s.Body
}

func (s *StopCloudNoteResponse) SetHeaders(v map[string]*string) *StopCloudNoteResponse {
	s.Headers = v
	return s
}

func (s *StopCloudNoteResponse) SetStatusCode(v int32) *StopCloudNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCloudNoteResponse) SetBody(v *StopCloudNoteResponseBody) *StopCloudNoteResponse {
	s.Body = v
	return s
}

func (s *StopCloudNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
