// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectHistoryRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBProjectHistoryRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBProjectHistoryRetentionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBProjectHistoryRetentionResponseBody) *DescribeAgenticDBProjectHistoryRetentionResponse
	GetBody() *DescribeAgenticDBProjectHistoryRetentionResponseBody
}

type DescribeAgenticDBProjectHistoryRetentionResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBProjectHistoryRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBProjectHistoryRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectHistoryRetentionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) GetBody() *DescribeAgenticDBProjectHistoryRetentionResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBProjectHistoryRetentionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) SetStatusCode(v int32) *DescribeAgenticDBProjectHistoryRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) SetBody(v *DescribeAgenticDBProjectHistoryRetentionResponseBody) *DescribeAgenticDBProjectHistoryRetentionResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBProjectHistoryRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
