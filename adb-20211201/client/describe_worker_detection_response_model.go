// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkerDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWorkerDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWorkerDetectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWorkerDetectionResponseBody) *DescribeWorkerDetectionResponse
	GetBody() *DescribeWorkerDetectionResponseBody
}

type DescribeWorkerDetectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkerDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkerDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWorkerDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWorkerDetectionResponse) GetBody() *DescribeWorkerDetectionResponseBody {
	return s.Body
}

func (s *DescribeWorkerDetectionResponse) SetHeaders(v map[string]*string) *DescribeWorkerDetectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkerDetectionResponse) SetStatusCode(v int32) *DescribeWorkerDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkerDetectionResponse) SetBody(v *DescribeWorkerDetectionResponseBody) *DescribeWorkerDetectionResponse {
	s.Body = v
	return s
}

func (s *DescribeWorkerDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
