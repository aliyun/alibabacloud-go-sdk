// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOrgByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOrgByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOrgByLayerResponseBody) *DescribeOrgByLayerResponse
	GetBody() *DescribeOrgByLayerResponseBody
}

type DescribeOrgByLayerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrgByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrgByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOrgByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOrgByLayerResponse) GetBody() *DescribeOrgByLayerResponseBody {
	return s.Body
}

func (s *DescribeOrgByLayerResponse) SetHeaders(v map[string]*string) *DescribeOrgByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgByLayerResponse) SetStatusCode(v int32) *DescribeOrgByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrgByLayerResponse) SetBody(v *DescribeOrgByLayerResponseBody) *DescribeOrgByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeOrgByLayerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
