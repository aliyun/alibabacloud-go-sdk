// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigCounterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigCounterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutomateResponseConfigCounterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutomateResponseConfigCounterResponseBody) *DescribeAutomateResponseConfigCounterResponse
	GetBody() *DescribeAutomateResponseConfigCounterResponseBody
}

type DescribeAutomateResponseConfigCounterResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigCounterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutomateResponseConfigCounterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutomateResponseConfigCounterResponse) GetBody() *DescribeAutomateResponseConfigCounterResponseBody {
	return s.Body
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetBody(v *DescribeAutomateResponseConfigCounterResponseBody) *DescribeAutomateResponseConfigCounterResponse {
	s.Body = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
