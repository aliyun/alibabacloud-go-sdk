// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageByNameResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageByNameResponseBody) *SearchImageByNameResponse
	GetBody() *SearchImageByNameResponseBody
}

type SearchImageByNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageByNameResponse) GetBody() *SearchImageByNameResponseBody {
	return s.Body
}

func (s *SearchImageByNameResponse) SetHeaders(v map[string]*string) *SearchImageByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByNameResponse) SetStatusCode(v int32) *SearchImageByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByNameResponse) SetBody(v *SearchImageByNameResponseBody) *SearchImageByNameResponse {
	s.Body = v
	return s
}

func (s *SearchImageByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
