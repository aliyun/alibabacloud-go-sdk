// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogicInstanceTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogicInstanceTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogicInstanceTopologyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogicInstanceTopologyResponseBody) *DescribeLogicInstanceTopologyResponse
	GetBody() *DescribeLogicInstanceTopologyResponseBody
}

type DescribeLogicInstanceTopologyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogicInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogicInstanceTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogicInstanceTopologyResponse) GetBody() *DescribeLogicInstanceTopologyResponseBody {
	return s.Body
}

func (s *DescribeLogicInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeLogicInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponse) SetStatusCode(v int32) *DescribeLogicInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponse) SetBody(v *DescribeLogicInstanceTopologyResponseBody) *DescribeLogicInstanceTopologyResponse {
	s.Body = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
