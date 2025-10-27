// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventOnStageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventOnStageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventOnStageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventOnStageResponseBody) *DescribeEventOnStageResponse
	GetBody() *DescribeEventOnStageResponseBody
}

type DescribeEventOnStageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventOnStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventOnStageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventOnStageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventOnStageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventOnStageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventOnStageResponse) GetBody() *DescribeEventOnStageResponseBody {
	return s.Body
}

func (s *DescribeEventOnStageResponse) SetHeaders(v map[string]*string) *DescribeEventOnStageResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventOnStageResponse) SetStatusCode(v int32) *DescribeEventOnStageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventOnStageResponse) SetBody(v *DescribeEventOnStageResponseBody) *DescribeEventOnStageResponse {
	s.Body = v
	return s
}

func (s *DescribeEventOnStageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
