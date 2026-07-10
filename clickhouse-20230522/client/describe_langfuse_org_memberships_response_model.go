// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseOrgMembershipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLangfuseOrgMembershipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLangfuseOrgMembershipsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLangfuseOrgMembershipsResponseBody) *DescribeLangfuseOrgMembershipsResponse
	GetBody() *DescribeLangfuseOrgMembershipsResponseBody
}

type DescribeLangfuseOrgMembershipsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLangfuseOrgMembershipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLangfuseOrgMembershipsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseOrgMembershipsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseOrgMembershipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLangfuseOrgMembershipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLangfuseOrgMembershipsResponse) GetBody() *DescribeLangfuseOrgMembershipsResponseBody {
	return s.Body
}

func (s *DescribeLangfuseOrgMembershipsResponse) SetHeaders(v map[string]*string) *DescribeLangfuseOrgMembershipsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponse) SetStatusCode(v int32) *DescribeLangfuseOrgMembershipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponse) SetBody(v *DescribeLangfuseOrgMembershipsResponseBody) *DescribeLangfuseOrgMembershipsResponse {
	s.Body = v
	return s
}

func (s *DescribeLangfuseOrgMembershipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
