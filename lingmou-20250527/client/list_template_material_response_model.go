// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateMaterialResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateMaterialResponseBody) *ListTemplateMaterialResponse
	GetBody() *ListTemplateMaterialResponseBody
}

type ListTemplateMaterialResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateMaterialResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateMaterialResponse) GetBody() *ListTemplateMaterialResponseBody {
	return s.Body
}

func (s *ListTemplateMaterialResponse) SetHeaders(v map[string]*string) *ListTemplateMaterialResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateMaterialResponse) SetStatusCode(v int32) *ListTemplateMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateMaterialResponse) SetBody(v *ListTemplateMaterialResponseBody) *ListTemplateMaterialResponse {
	s.Body = v
	return s
}

func (s *ListTemplateMaterialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
