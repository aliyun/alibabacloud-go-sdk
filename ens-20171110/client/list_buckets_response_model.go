// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketsResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketsResponseBody) *ListBucketsResponse
	GetBody() *ListBucketsResponseBody
}

type ListBucketsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketsResponse) GetBody() *ListBucketsResponseBody {
	return s.Body
}

func (s *ListBucketsResponse) SetHeaders(v map[string]*string) *ListBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListBucketsResponse) SetStatusCode(v int32) *ListBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketsResponse) SetBody(v *ListBucketsResponseBody) *ListBucketsResponse {
	s.Body = v
	return s
}

func (s *ListBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
