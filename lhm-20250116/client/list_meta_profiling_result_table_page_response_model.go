// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultTablePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaProfilingResultTablePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaProfilingResultTablePageResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaProfilingResultTablePageResponseBody) *ListMetaProfilingResultTablePageResponse
	GetBody() *ListMetaProfilingResultTablePageResponseBody
}

type ListMetaProfilingResultTablePageResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaProfilingResultTablePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaProfilingResultTablePageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultTablePageResponse) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultTablePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaProfilingResultTablePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaProfilingResultTablePageResponse) GetBody() *ListMetaProfilingResultTablePageResponseBody {
	return s.Body
}

func (s *ListMetaProfilingResultTablePageResponse) SetHeaders(v map[string]*string) *ListMetaProfilingResultTablePageResponse {
	s.Headers = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponse) SetStatusCode(v int32) *ListMetaProfilingResultTablePageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponse) SetBody(v *ListMetaProfilingResultTablePageResponseBody) *ListMetaProfilingResultTablePageResponse {
	s.Body = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
