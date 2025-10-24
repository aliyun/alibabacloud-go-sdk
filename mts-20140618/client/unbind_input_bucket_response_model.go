// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInputBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindInputBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindInputBucketResponse
	GetStatusCode() *int32
	SetBody(v *UnbindInputBucketResponseBody) *UnbindInputBucketResponse
	GetBody() *UnbindInputBucketResponseBody
}

type UnbindInputBucketResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindInputBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindInputBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindInputBucketResponse) GoString() string {
	return s.String()
}

func (s *UnbindInputBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindInputBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindInputBucketResponse) GetBody() *UnbindInputBucketResponseBody {
	return s.Body
}

func (s *UnbindInputBucketResponse) SetHeaders(v map[string]*string) *UnbindInputBucketResponse {
	s.Headers = v
	return s
}

func (s *UnbindInputBucketResponse) SetStatusCode(v int32) *UnbindInputBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindInputBucketResponse) SetBody(v *UnbindInputBucketResponseBody) *UnbindInputBucketResponse {
	s.Body = v
	return s
}

func (s *UnbindInputBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
