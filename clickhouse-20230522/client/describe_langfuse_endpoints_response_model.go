// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseEndpointsResponseBody) *DescribeLangfuseEndpointsResponse
	GetBody() *DescribeLangfuseEndpointsResponseBody
}

type DescribeLangfuseEndpointsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseEndpointsResponse) GetBody() *DescribeLangfuseEndpointsResponseBody {
	return s.Body
}

func (s *DescribeLangfuseEndpointsResponse) SetHeaders(v map[string]*string) *DescribeLangfuseEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseEndpointsResponse) SetStatusCode(v int32) *DescribeLangfuseEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponse) SetBody(v *DescribeLangfuseEndpointsResponseBody) *DescribeLangfuseEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
