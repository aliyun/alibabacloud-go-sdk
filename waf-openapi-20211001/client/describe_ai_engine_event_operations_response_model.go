// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEventOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEventOperationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEventOperationsResponseBody) *DescribeAiEngineEventOperationsResponse
	GetBody() *DescribeAiEngineEventOperationsResponseBody
}

type DescribeAiEngineEventOperationsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEventOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEventOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEventOperationsResponse) GetBody() *DescribeAiEngineEventOperationsResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEventOperationsResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEventOperationsResponse) SetStatusCode(v int32) *DescribeAiEngineEventOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEventOperationsResponse) SetBody(v *DescribeAiEngineEventOperationsResponseBody) *DescribeAiEngineEventOperationsResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEventOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
