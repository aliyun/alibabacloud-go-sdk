// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultUdfPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaProfilingResultUdfPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaProfilingResultUdfPageResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaProfilingResultUdfPageResponseBody) *ListMetaProfilingResultUdfPageResponse
	GetBody() *ListMetaProfilingResultUdfPageResponseBody
}

type ListMetaProfilingResultUdfPageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaProfilingResultUdfPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaProfilingResultUdfPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultUdfPageResponse) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultUdfPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaProfilingResultUdfPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaProfilingResultUdfPageResponse) GetBody() *ListMetaProfilingResultUdfPageResponseBody {
	return s.Body
}

func (s *ListMetaProfilingResultUdfPageResponse) SetHeaders(v map[string]*string) *ListMetaProfilingResultUdfPageResponse {
	s.Headers = v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponse) SetStatusCode(v int32) *ListMetaProfilingResultUdfPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponse) SetBody(v *ListMetaProfilingResultUdfPageResponseBody) *ListMetaProfilingResultUdfPageResponse {
	s.Body = v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
