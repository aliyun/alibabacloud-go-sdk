// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationPreditInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimulationPreditInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimulationPreditInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimulationPreditInfoResponseBody) *DescribeSimulationPreditInfoResponse
	GetBody() *DescribeSimulationPreditInfoResponseBody
}

type DescribeSimulationPreditInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimulationPreditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimulationPreditInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationPreditInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimulationPreditInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimulationPreditInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimulationPreditInfoResponse) GetBody() *DescribeSimulationPreditInfoResponseBody {
	return s.Body
}

func (s *DescribeSimulationPreditInfoResponse) SetHeaders(v map[string]*string) *DescribeSimulationPreditInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimulationPreditInfoResponse) SetStatusCode(v int32) *DescribeSimulationPreditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimulationPreditInfoResponse) SetBody(v *DescribeSimulationPreditInfoResponseBody) *DescribeSimulationPreditInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSimulationPreditInfoResponse) Validate() error {
	return dara.Validate(s)
}
