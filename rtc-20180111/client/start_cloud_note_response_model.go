// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCloudNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCloudNoteResponse
	GetStatusCode() *int32
	SetBody(v *StartCloudNoteResponseBody) *StartCloudNoteResponse
	GetBody() *StartCloudNoteResponseBody
}

type StartCloudNoteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCloudNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCloudNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteResponse) GoString() string {
	return s.String()
}

func (s *StartCloudNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCloudNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCloudNoteResponse) GetBody() *StartCloudNoteResponseBody {
	return s.Body
}

func (s *StartCloudNoteResponse) SetHeaders(v map[string]*string) *StartCloudNoteResponse {
	s.Headers = v
	return s
}

func (s *StartCloudNoteResponse) SetStatusCode(v int32) *StartCloudNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCloudNoteResponse) SetBody(v *StartCloudNoteResponseBody) *StartCloudNoteResponse {
	s.Body = v
	return s
}

func (s *StartCloudNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
