// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilePutUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FilePutUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FilePutUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *FilePutUserTagsResponseBody) *FilePutUserTagsResponse
	GetBody() *FilePutUserTagsResponseBody
}

type FilePutUserTagsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FilePutUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FilePutUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s FilePutUserTagsResponse) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FilePutUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FilePutUserTagsResponse) GetBody() *FilePutUserTagsResponseBody {
	return s.Body
}

func (s *FilePutUserTagsResponse) SetHeaders(v map[string]*string) *FilePutUserTagsResponse {
	s.Headers = v
	return s
}

func (s *FilePutUserTagsResponse) SetStatusCode(v int32) *FilePutUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *FilePutUserTagsResponse) SetBody(v *FilePutUserTagsResponseBody) *FilePutUserTagsResponse {
	s.Body = v
	return s
}

func (s *FilePutUserTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
