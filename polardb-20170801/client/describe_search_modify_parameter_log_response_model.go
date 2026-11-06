// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchModifyParameterLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSearchModifyParameterLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSearchModifyParameterLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSearchModifyParameterLogResponseBody) *DescribeSearchModifyParameterLogResponse
	GetBody() *DescribeSearchModifyParameterLogResponseBody
}

type DescribeSearchModifyParameterLogResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSearchModifyParameterLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSearchModifyParameterLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchModifyParameterLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSearchModifyParameterLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSearchModifyParameterLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSearchModifyParameterLogResponse) GetBody() *DescribeSearchModifyParameterLogResponseBody {
	return s.Body
}

func (s *DescribeSearchModifyParameterLogResponse) SetHeaders(v map[string]*string) *DescribeSearchModifyParameterLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSearchModifyParameterLogResponse) SetStatusCode(v int32) *DescribeSearchModifyParameterLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSearchModifyParameterLogResponse) SetBody(v *DescribeSearchModifyParameterLogResponseBody) *DescribeSearchModifyParameterLogResponse {
	s.Body = v
	return s
}

func (s *DescribeSearchModifyParameterLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
