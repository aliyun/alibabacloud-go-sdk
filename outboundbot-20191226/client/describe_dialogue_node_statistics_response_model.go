// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDialogueNodeStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDialogueNodeStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDialogueNodeStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDialogueNodeStatisticsResponseBody) *DescribeDialogueNodeStatisticsResponse
	GetBody() *DescribeDialogueNodeStatisticsResponseBody
}

type DescribeDialogueNodeStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDialogueNodeStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDialogueNodeStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDialogueNodeStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDialogueNodeStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDialogueNodeStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDialogueNodeStatisticsResponse) GetBody() *DescribeDialogueNodeStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDialogueNodeStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDialogueNodeStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponse) SetStatusCode(v int32) *DescribeDialogueNodeStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponse) SetBody(v *DescribeDialogueNodeStatisticsResponseBody) *DescribeDialogueNodeStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDialogueNodeStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
