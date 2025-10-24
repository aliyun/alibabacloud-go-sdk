// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMediaBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllMediaBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllMediaBucketResponse
	GetStatusCode() *int32
	SetBody(v *ListAllMediaBucketResponseBody) *ListAllMediaBucketResponse
	GetBody() *ListAllMediaBucketResponseBody
}

type ListAllMediaBucketResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllMediaBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllMediaBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllMediaBucketResponse) GoString() string {
	return s.String()
}

func (s *ListAllMediaBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllMediaBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllMediaBucketResponse) GetBody() *ListAllMediaBucketResponseBody {
	return s.Body
}

func (s *ListAllMediaBucketResponse) SetHeaders(v map[string]*string) *ListAllMediaBucketResponse {
	s.Headers = v
	return s
}

func (s *ListAllMediaBucketResponse) SetStatusCode(v int32) *ListAllMediaBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllMediaBucketResponse) SetBody(v *ListAllMediaBucketResponseBody) *ListAllMediaBucketResponse {
	s.Body = v
	return s
}

func (s *ListAllMediaBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
