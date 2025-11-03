// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationTaskRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationTaskRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationTaskRegionResponseBody) *DescribeActiveOperationTaskRegionResponse
	GetBody() *DescribeActiveOperationTaskRegionResponseBody
}

type DescribeActiveOperationTaskRegionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTaskRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationTaskRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationTaskRegionResponse) GetBody() *DescribeActiveOperationTaskRegionResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationTaskRegionResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponse) SetBody(v *DescribeActiveOperationTaskRegionResponseBody) *DescribeActiveOperationTaskRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
