// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SizeChartExtractResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SizeChartExtractResponse
	GetStatusCode() *int32
	SetBody(v *SizeChartExtractResponseBody) *SizeChartExtractResponse
	GetBody() *SizeChartExtractResponseBody
}

type SizeChartExtractResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SizeChartExtractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SizeChartExtractResponse) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractResponse) GoString() string {
	return s.String()
}

func (s *SizeChartExtractResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SizeChartExtractResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SizeChartExtractResponse) GetBody() *SizeChartExtractResponseBody {
	return s.Body
}

func (s *SizeChartExtractResponse) SetHeaders(v map[string]*string) *SizeChartExtractResponse {
	s.Headers = v
	return s
}

func (s *SizeChartExtractResponse) SetStatusCode(v int32) *SizeChartExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *SizeChartExtractResponse) SetBody(v *SizeChartExtractResponseBody) *SizeChartExtractResponse {
	s.Body = v
	return s
}

func (s *SizeChartExtractResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
