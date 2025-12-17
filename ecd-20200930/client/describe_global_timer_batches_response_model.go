// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerBatchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalTimerBatchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalTimerBatchesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalTimerBatchesResponseBody) *DescribeGlobalTimerBatchesResponse
	GetBody() *DescribeGlobalTimerBatchesResponseBody
}

type DescribeGlobalTimerBatchesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalTimerBatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalTimerBatchesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerBatchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerBatchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalTimerBatchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalTimerBatchesResponse) GetBody() *DescribeGlobalTimerBatchesResponseBody {
	return s.Body
}

func (s *DescribeGlobalTimerBatchesResponse) SetHeaders(v map[string]*string) *DescribeGlobalTimerBatchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalTimerBatchesResponse) SetStatusCode(v int32) *DescribeGlobalTimerBatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponse) SetBody(v *DescribeGlobalTimerBatchesResponseBody) *DescribeGlobalTimerBatchesResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalTimerBatchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
