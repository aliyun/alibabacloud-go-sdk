// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineScanReportUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineScanReportUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineScanReportUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineScanReportUrlResponseBody) *GetPipelineScanReportUrlResponse
	GetBody() *GetPipelineScanReportUrlResponseBody
}

type GetPipelineScanReportUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineScanReportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineScanReportUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineScanReportUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineScanReportUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineScanReportUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineScanReportUrlResponse) GetBody() *GetPipelineScanReportUrlResponseBody {
	return s.Body
}

func (s *GetPipelineScanReportUrlResponse) SetHeaders(v map[string]*string) *GetPipelineScanReportUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineScanReportUrlResponse) SetStatusCode(v int32) *GetPipelineScanReportUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineScanReportUrlResponse) SetBody(v *GetPipelineScanReportUrlResponseBody) *GetPipelineScanReportUrlResponse {
	s.Body = v
	return s
}

func (s *GetPipelineScanReportUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
