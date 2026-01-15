// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageByTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageByTextResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageByTextResponseBody) *SearchImageByTextResponse
	GetBody() *SearchImageByTextResponseBody
}

type SearchImageByTextResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByTextResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageByTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageByTextResponse) GetBody() *SearchImageByTextResponseBody {
	return s.Body
}

func (s *SearchImageByTextResponse) SetHeaders(v map[string]*string) *SearchImageByTextResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByTextResponse) SetStatusCode(v int32) *SearchImageByTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByTextResponse) SetBody(v *SearchImageByTextResponseBody) *SearchImageByTextResponse {
	s.Body = v
	return s
}

func (s *SearchImageByTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
