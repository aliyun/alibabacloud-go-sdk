// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaByFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaByFaceResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaByFaceResponseBody) *SearchMediaByFaceResponse
	GetBody() *SearchMediaByFaceResponseBody
}

type SearchMediaByFaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaByFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaByFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaByFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaByFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaByFaceResponse) GetBody() *SearchMediaByFaceResponseBody {
	return s.Body
}

func (s *SearchMediaByFaceResponse) SetHeaders(v map[string]*string) *SearchMediaByFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaByFaceResponse) SetStatusCode(v int32) *SearchMediaByFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaByFaceResponse) SetBody(v *SearchMediaByFaceResponseBody) *SearchMediaByFaceResponse {
	s.Body = v
	return s
}

func (s *SearchMediaByFaceResponse) Validate() error {
	return dara.Validate(s)
}
