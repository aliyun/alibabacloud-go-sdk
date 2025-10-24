// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInputBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindInputBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindInputBucketResponse
	GetStatusCode() *int32
	SetBody(v *BindInputBucketResponseBody) *BindInputBucketResponse
	GetBody() *BindInputBucketResponseBody
}

type BindInputBucketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindInputBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindInputBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s BindInputBucketResponse) GoString() string {
	return s.String()
}

func (s *BindInputBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindInputBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindInputBucketResponse) GetBody() *BindInputBucketResponseBody {
	return s.Body
}

func (s *BindInputBucketResponse) SetHeaders(v map[string]*string) *BindInputBucketResponse {
	s.Headers = v
	return s
}

func (s *BindInputBucketResponse) SetStatusCode(v int32) *BindInputBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *BindInputBucketResponse) SetBody(v *BindInputBucketResponseBody) *BindInputBucketResponse {
	s.Body = v
	return s
}

func (s *BindInputBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
