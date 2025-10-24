// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindOutputBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindOutputBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindOutputBucketResponse
	GetStatusCode() *int32
	SetBody(v *BindOutputBucketResponseBody) *BindOutputBucketResponse
	GetBody() *BindOutputBucketResponseBody
}

type BindOutputBucketResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindOutputBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindOutputBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s BindOutputBucketResponse) GoString() string {
	return s.String()
}

func (s *BindOutputBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindOutputBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindOutputBucketResponse) GetBody() *BindOutputBucketResponseBody {
	return s.Body
}

func (s *BindOutputBucketResponse) SetHeaders(v map[string]*string) *BindOutputBucketResponse {
	s.Headers = v
	return s
}

func (s *BindOutputBucketResponse) SetStatusCode(v int32) *BindOutputBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *BindOutputBucketResponse) SetBody(v *BindOutputBucketResponseBody) *BindOutputBucketResponse {
	s.Body = v
	return s
}

func (s *BindOutputBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
