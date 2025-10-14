// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLifecycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketLifecycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketLifecycleResponse
	GetStatusCode() *int32
	SetBody(v *PutBucketLifecycleResponseBody) *PutBucketLifecycleResponse
	GetBody() *PutBucketLifecycleResponseBody
}

type PutBucketLifecycleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutBucketLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketLifecycleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketLifecycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketLifecycleResponse) GetBody() *PutBucketLifecycleResponseBody {
	return s.Body
}

func (s *PutBucketLifecycleResponse) SetHeaders(v map[string]*string) *PutBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *PutBucketLifecycleResponse) SetStatusCode(v int32) *PutBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketLifecycleResponse) SetBody(v *PutBucketLifecycleResponseBody) *PutBucketLifecycleResponse {
	s.Body = v
	return s
}

func (s *PutBucketLifecycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
