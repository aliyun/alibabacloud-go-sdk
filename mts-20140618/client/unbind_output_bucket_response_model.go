// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindOutputBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindOutputBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindOutputBucketResponse
	GetStatusCode() *int32
	SetBody(v *UnbindOutputBucketResponseBody) *UnbindOutputBucketResponse
	GetBody() *UnbindOutputBucketResponseBody
}

type UnbindOutputBucketResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindOutputBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindOutputBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindOutputBucketResponse) GoString() string {
	return s.String()
}

func (s *UnbindOutputBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindOutputBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindOutputBucketResponse) GetBody() *UnbindOutputBucketResponseBody {
	return s.Body
}

func (s *UnbindOutputBucketResponse) SetHeaders(v map[string]*string) *UnbindOutputBucketResponse {
	s.Headers = v
	return s
}

func (s *UnbindOutputBucketResponse) SetStatusCode(v int32) *UnbindOutputBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindOutputBucketResponse) SetBody(v *UnbindOutputBucketResponseBody) *UnbindOutputBucketResponse {
	s.Body = v
	return s
}

func (s *UnbindOutputBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
