// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnionMediaIndustryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnionMediaIndustryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnionMediaIndustryResponse
	GetStatusCode() *int32
	SetBody(v *ListUnionMediaIndustryResponseBody) *ListUnionMediaIndustryResponse
	GetBody() *ListUnionMediaIndustryResponseBody
}

type ListUnionMediaIndustryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnionMediaIndustryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnionMediaIndustryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnionMediaIndustryResponse) GoString() string {
	return s.String()
}

func (s *ListUnionMediaIndustryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnionMediaIndustryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnionMediaIndustryResponse) GetBody() *ListUnionMediaIndustryResponseBody {
	return s.Body
}

func (s *ListUnionMediaIndustryResponse) SetHeaders(v map[string]*string) *ListUnionMediaIndustryResponse {
	s.Headers = v
	return s
}

func (s *ListUnionMediaIndustryResponse) SetStatusCode(v int32) *ListUnionMediaIndustryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnionMediaIndustryResponse) SetBody(v *ListUnionMediaIndustryResponseBody) *ListUnionMediaIndustryResponse {
	s.Body = v
	return s
}

func (s *ListUnionMediaIndustryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
