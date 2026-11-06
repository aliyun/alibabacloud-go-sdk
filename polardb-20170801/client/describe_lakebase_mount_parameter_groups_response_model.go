// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakebaseMountParameterGroupsResponseBody) *DescribeLakebaseMountParameterGroupsResponse
	GetBody() *DescribeLakebaseMountParameterGroupsResponseBody
}

type DescribeLakebaseMountParameterGroupsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakebaseMountParameterGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakebaseMountParameterGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakebaseMountParameterGroupsResponse) GetBody() *DescribeLakebaseMountParameterGroupsResponseBody {
	return s.Body
}

func (s *DescribeLakebaseMountParameterGroupsResponse) SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponse) SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponse) SetBody(v *DescribeLakebaseMountParameterGroupsResponseBody) *DescribeLakebaseMountParameterGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
