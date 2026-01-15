// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSimilarByImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompareSimilarByImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompareSimilarByImageResponse
	GetStatusCode() *int32
	SetBody(v *CompareSimilarByImageResponseBody) *CompareSimilarByImageResponse
	GetBody() *CompareSimilarByImageResponseBody
}

type CompareSimilarByImageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareSimilarByImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareSimilarByImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageResponse) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompareSimilarByImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompareSimilarByImageResponse) GetBody() *CompareSimilarByImageResponseBody {
	return s.Body
}

func (s *CompareSimilarByImageResponse) SetHeaders(v map[string]*string) *CompareSimilarByImageResponse {
	s.Headers = v
	return s
}

func (s *CompareSimilarByImageResponse) SetStatusCode(v int32) *CompareSimilarByImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareSimilarByImageResponse) SetBody(v *CompareSimilarByImageResponseBody) *CompareSimilarByImageResponse {
	s.Body = v
	return s
}

func (s *CompareSimilarByImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
