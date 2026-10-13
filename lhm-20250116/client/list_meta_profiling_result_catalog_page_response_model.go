// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultCatalogPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaProfilingResultCatalogPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaProfilingResultCatalogPageResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaProfilingResultCatalogPageResponseBody) *ListMetaProfilingResultCatalogPageResponse
	GetBody() *ListMetaProfilingResultCatalogPageResponseBody
}

type ListMetaProfilingResultCatalogPageResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaProfilingResultCatalogPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaProfilingResultCatalogPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultCatalogPageResponse) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultCatalogPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaProfilingResultCatalogPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaProfilingResultCatalogPageResponse) GetBody() *ListMetaProfilingResultCatalogPageResponseBody {
	return s.Body
}

func (s *ListMetaProfilingResultCatalogPageResponse) SetHeaders(v map[string]*string) *ListMetaProfilingResultCatalogPageResponse {
	s.Headers = v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponse) SetStatusCode(v int32) *ListMetaProfilingResultCatalogPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponse) SetBody(v *ListMetaProfilingResultCatalogPageResponseBody) *ListMetaProfilingResultCatalogPageResponse {
	s.Body = v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
