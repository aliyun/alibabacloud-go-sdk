// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketLifecycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketLifecycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketLifecycleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBucketLifecycleResponseBody) *DeleteBucketLifecycleResponse
	GetBody() *DeleteBucketLifecycleResponseBody
}

type DeleteBucketLifecycleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBucketLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBucketLifecycleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketLifecycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketLifecycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketLifecycleResponse) GetBody() *DeleteBucketLifecycleResponseBody {
	return s.Body
}

func (s *DeleteBucketLifecycleResponse) SetHeaders(v map[string]*string) *DeleteBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketLifecycleResponse) SetStatusCode(v int32) *DeleteBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketLifecycleResponse) SetBody(v *DeleteBucketLifecycleResponseBody) *DeleteBucketLifecycleResponse {
	s.Body = v
	return s
}

func (s *DeleteBucketLifecycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
