// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFaceResponse
	GetStatusCode() *int32
	SetBody(v *SearchFaceResponseBody) *SearchFaceResponse
	GetBody() *SearchFaceResponseBody
}

type SearchFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFaceResponse) GetBody() *SearchFaceResponseBody {
	return s.Body
}

func (s *SearchFaceResponse) SetHeaders(v map[string]*string) *SearchFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchFaceResponse) SetStatusCode(v int32) *SearchFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFaceResponse) SetBody(v *SearchFaceResponseBody) *SearchFaceResponse {
	s.Body = v
	return s
}

func (s *SearchFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
