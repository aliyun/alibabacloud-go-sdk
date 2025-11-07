// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasSexStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyPersonasSexStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyPersonasSexStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyPersonasSexStatisticsResponseBody) *DescribeVerifyPersonasSexStatisticsResponse
	GetBody() *DescribeVerifyPersonasSexStatisticsResponseBody
}

type DescribeVerifyPersonasSexStatisticsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyPersonasSexStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyPersonasSexStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasSexStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) GetBody() *DescribeVerifyPersonasSexStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyPersonasSexStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyPersonasSexStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) SetBody(v *DescribeVerifyPersonasSexStatisticsResponseBody) *DescribeVerifyPersonasSexStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
