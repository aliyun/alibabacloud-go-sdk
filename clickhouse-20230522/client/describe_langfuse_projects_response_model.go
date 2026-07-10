// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseProjectsResponseBody) *DescribeLangfuseProjectsResponse
	GetBody() *DescribeLangfuseProjectsResponseBody
}

type DescribeLangfuseProjectsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseProjectsResponse) GetBody() *DescribeLangfuseProjectsResponseBody {
	return s.Body
}

func (s *DescribeLangfuseProjectsResponse) SetHeaders(v map[string]*string) *DescribeLangfuseProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseProjectsResponse) SetStatusCode(v int32) *DescribeLangfuseProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseProjectsResponse) SetBody(v *DescribeLangfuseProjectsResponseBody) *DescribeLangfuseProjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
