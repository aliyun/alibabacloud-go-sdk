// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByPicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageByPicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageByPicResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageByPicResponseBody) *SearchImageByPicResponse
	GetBody() *SearchImageByPicResponseBody
}

type SearchImageByPicResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByPicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByPicResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageByPicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageByPicResponse) GetBody() *SearchImageByPicResponseBody {
	return s.Body
}

func (s *SearchImageByPicResponse) SetHeaders(v map[string]*string) *SearchImageByPicResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByPicResponse) SetStatusCode(v int32) *SearchImageByPicResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByPicResponse) SetBody(v *SearchImageByPicResponseBody) *SearchImageByPicResponse {
	s.Body = v
	return s
}

func (s *SearchImageByPicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
