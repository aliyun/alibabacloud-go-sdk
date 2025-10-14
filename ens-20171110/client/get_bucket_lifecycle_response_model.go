// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketLifecycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketLifecycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketLifecycleResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketLifecycleResponseBody) *GetBucketLifecycleResponse
	GetBody() *GetBucketLifecycleResponseBody
}

type GetBucketLifecycleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketLifecycleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketLifecycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketLifecycleResponse) GetBody() *GetBucketLifecycleResponseBody {
	return s.Body
}

func (s *GetBucketLifecycleResponse) SetHeaders(v map[string]*string) *GetBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLifecycleResponse) SetStatusCode(v int32) *GetBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLifecycleResponse) SetBody(v *GetBucketLifecycleResponseBody) *GetBucketLifecycleResponse {
	s.Body = v
	return s
}

func (s *GetBucketLifecycleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
