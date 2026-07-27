// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOUsageDetailExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMOUsageDetailExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMOUsageDetailExportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMOUsageDetailExportResponseBody) *DescribeMOUsageDetailExportResponse
	GetBody() *DescribeMOUsageDetailExportResponseBody
}

type DescribeMOUsageDetailExportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMOUsageDetailExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMOUsageDetailExportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOUsageDetailExportResponse) GoString() string {
	return s.String()
}

func (s *DescribeMOUsageDetailExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMOUsageDetailExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMOUsageDetailExportResponse) GetBody() *DescribeMOUsageDetailExportResponseBody {
	return s.Body
}

func (s *DescribeMOUsageDetailExportResponse) SetHeaders(v map[string]*string) *DescribeMOUsageDetailExportResponse {
	s.Headers = v
	return s
}

func (s *DescribeMOUsageDetailExportResponse) SetStatusCode(v int32) *DescribeMOUsageDetailExportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMOUsageDetailExportResponse) SetBody(v *DescribeMOUsageDetailExportResponseBody) *DescribeMOUsageDetailExportResponse {
	s.Body = v
	return s
}

func (s *DescribeMOUsageDetailExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
