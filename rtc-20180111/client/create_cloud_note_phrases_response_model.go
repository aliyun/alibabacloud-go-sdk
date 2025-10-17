// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudNotePhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudNotePhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudNotePhrasesResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudNotePhrasesResponseBody) *CreateCloudNotePhrasesResponse
	GetBody() *CreateCloudNotePhrasesResponseBody
}

type CreateCloudNotePhrasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudNotePhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudNotePhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudNotePhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudNotePhrasesResponse) GetBody() *CreateCloudNotePhrasesResponseBody {
	return s.Body
}

func (s *CreateCloudNotePhrasesResponse) SetHeaders(v map[string]*string) *CreateCloudNotePhrasesResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudNotePhrasesResponse) SetStatusCode(v int32) *CreateCloudNotePhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudNotePhrasesResponse) SetBody(v *CreateCloudNotePhrasesResponseBody) *CreateCloudNotePhrasesResponse {
	s.Body = v
	return s
}

func (s *CreateCloudNotePhrasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
