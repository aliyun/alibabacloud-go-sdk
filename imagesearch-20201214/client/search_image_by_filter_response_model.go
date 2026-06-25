// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageByFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageByFilterResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageByFilterResponseBody) *SearchImageByFilterResponse
	GetBody() *SearchImageByFilterResponseBody
}

type SearchImageByFilterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByFilterResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageByFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageByFilterResponse) GetBody() *SearchImageByFilterResponseBody {
	return s.Body
}

func (s *SearchImageByFilterResponse) SetHeaders(v map[string]*string) *SearchImageByFilterResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByFilterResponse) SetStatusCode(v int32) *SearchImageByFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByFilterResponse) SetBody(v *SearchImageByFilterResponseBody) *SearchImageByFilterResponse {
	s.Body = v
	return s
}

func (s *SearchImageByFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
