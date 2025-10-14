// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSL7QpsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSL7QpsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSL7QpsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSL7QpsListResponseBody) *DescribeDDoSL7QpsListResponse
	GetBody() *DescribeDDoSL7QpsListResponseBody
}

type DescribeDDoSL7QpsListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSL7QpsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSL7QpsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSL7QpsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSL7QpsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSL7QpsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSL7QpsListResponse) GetBody() *DescribeDDoSL7QpsListResponseBody {
	return s.Body
}

func (s *DescribeDDoSL7QpsListResponse) SetHeaders(v map[string]*string) *DescribeDDoSL7QpsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSL7QpsListResponse) SetStatusCode(v int32) *DescribeDDoSL7QpsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSL7QpsListResponse) SetBody(v *DescribeDDoSL7QpsListResponseBody) *DescribeDDoSL7QpsListResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSL7QpsListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
