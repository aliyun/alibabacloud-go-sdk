// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaClipByFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaClipByFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaClipByFaceResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaClipByFaceResponseBody) *SearchMediaClipByFaceResponse
	GetBody() *SearchMediaClipByFaceResponseBody
}

type SearchMediaClipByFaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaClipByFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaClipByFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaClipByFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaClipByFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaClipByFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaClipByFaceResponse) GetBody() *SearchMediaClipByFaceResponseBody {
	return s.Body
}

func (s *SearchMediaClipByFaceResponse) SetHeaders(v map[string]*string) *SearchMediaClipByFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaClipByFaceResponse) SetStatusCode(v int32) *SearchMediaClipByFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaClipByFaceResponse) SetBody(v *SearchMediaClipByFaceResponseBody) *SearchMediaClipByFaceResponse {
	s.Body = v
	return s
}

func (s *SearchMediaClipByFaceResponse) Validate() error {
	return dara.Validate(s)
}
