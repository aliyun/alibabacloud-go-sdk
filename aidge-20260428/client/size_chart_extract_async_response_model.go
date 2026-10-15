// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartExtractAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SizeChartExtractAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SizeChartExtractAsyncResponse
	GetStatusCode() *int32
	SetBody(v *SizeChartExtractAsyncResponseBody) *SizeChartExtractAsyncResponse
	GetBody() *SizeChartExtractAsyncResponseBody
}

type SizeChartExtractAsyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SizeChartExtractAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SizeChartExtractAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s SizeChartExtractAsyncResponse) GoString() string {
	return s.String()
}

func (s *SizeChartExtractAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SizeChartExtractAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SizeChartExtractAsyncResponse) GetBody() *SizeChartExtractAsyncResponseBody {
	return s.Body
}

func (s *SizeChartExtractAsyncResponse) SetHeaders(v map[string]*string) *SizeChartExtractAsyncResponse {
	s.Headers = v
	return s
}

func (s *SizeChartExtractAsyncResponse) SetStatusCode(v int32) *SizeChartExtractAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *SizeChartExtractAsyncResponse) SetBody(v *SizeChartExtractAsyncResponseBody) *SizeChartExtractAsyncResponse {
	s.Body = v
	return s
}

func (s *SizeChartExtractAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
