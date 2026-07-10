// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectMembershipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseProjectMembershipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseProjectMembershipsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseProjectMembershipsResponseBody) *DescribeLangfuseProjectMembershipsResponse
	GetBody() *DescribeLangfuseProjectMembershipsResponseBody
}

type DescribeLangfuseProjectMembershipsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseProjectMembershipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseProjectMembershipsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectMembershipsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectMembershipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseProjectMembershipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseProjectMembershipsResponse) GetBody() *DescribeLangfuseProjectMembershipsResponseBody {
	return s.Body
}

func (s *DescribeLangfuseProjectMembershipsResponse) SetHeaders(v map[string]*string) *DescribeLangfuseProjectMembershipsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponse) SetStatusCode(v int32) *DescribeLangfuseProjectMembershipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponse) SetBody(v *DescribeLangfuseProjectMembershipsResponseBody) *DescribeLangfuseProjectMembershipsResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseProjectMembershipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
