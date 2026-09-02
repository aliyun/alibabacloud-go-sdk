// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidVulNumClassifyStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUuidVulNumClassifyStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUuidVulNumClassifyStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUuidVulNumClassifyStatisticResponseBody) *DescribeUuidVulNumClassifyStatisticResponse
	GetBody() *DescribeUuidVulNumClassifyStatisticResponseBody
}

type DescribeUuidVulNumClassifyStatisticResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUuidVulNumClassifyStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUuidVulNumClassifyStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidVulNumClassifyStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) GetBody() *DescribeUuidVulNumClassifyStatisticResponseBody {
	return s.Body
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) SetHeaders(v map[string]*string) *DescribeUuidVulNumClassifyStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) SetStatusCode(v int32) *DescribeUuidVulNumClassifyStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) SetBody(v *DescribeUuidVulNumClassifyStatisticResponseBody) *DescribeUuidVulNumClassifyStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
