// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudNotePhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudNotePhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudNotePhrasesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudNotePhrasesResponseBody) *DeleteCloudNotePhrasesResponse
	GetBody() *DeleteCloudNotePhrasesResponseBody
}

type DeleteCloudNotePhrasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudNotePhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudNotePhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudNotePhrasesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudNotePhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudNotePhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudNotePhrasesResponse) GetBody() *DeleteCloudNotePhrasesResponseBody {
	return s.Body
}

func (s *DeleteCloudNotePhrasesResponse) SetHeaders(v map[string]*string) *DeleteCloudNotePhrasesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudNotePhrasesResponse) SetStatusCode(v int32) *DeleteCloudNotePhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudNotePhrasesResponse) SetBody(v *DeleteCloudNotePhrasesResponseBody) *DeleteCloudNotePhrasesResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudNotePhrasesResponse) Validate() error {
	return dara.Validate(s)
}
