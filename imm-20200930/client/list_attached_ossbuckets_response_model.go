// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttachedOSSBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttachedOSSBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttachedOSSBucketsResponse
	GetStatusCode() *int32
	SetBody(v *ListAttachedOSSBucketsResponseBody) *ListAttachedOSSBucketsResponse
	GetBody() *ListAttachedOSSBucketsResponseBody
}

type ListAttachedOSSBucketsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttachedOSSBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttachedOSSBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttachedOSSBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListAttachedOSSBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttachedOSSBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttachedOSSBucketsResponse) GetBody() *ListAttachedOSSBucketsResponseBody {
	return s.Body
}

func (s *ListAttachedOSSBucketsResponse) SetHeaders(v map[string]*string) *ListAttachedOSSBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListAttachedOSSBucketsResponse) SetStatusCode(v int32) *ListAttachedOSSBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttachedOSSBucketsResponse) SetBody(v *ListAttachedOSSBucketsResponseBody) *ListAttachedOSSBucketsResponse {
	s.Body = v
	return s
}

func (s *ListAttachedOSSBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
