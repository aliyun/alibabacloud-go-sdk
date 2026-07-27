// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMOUsageDetailExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMOUsageDetailExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMOUsageDetailExportResponse
	GetStatusCode() *int32
	SetBody(v *CreateMOUsageDetailExportResponseBody) *CreateMOUsageDetailExportResponse
	GetBody() *CreateMOUsageDetailExportResponseBody
}

type CreateMOUsageDetailExportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMOUsageDetailExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMOUsageDetailExportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMOUsageDetailExportResponse) GoString() string {
	return s.String()
}

func (s *CreateMOUsageDetailExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMOUsageDetailExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMOUsageDetailExportResponse) GetBody() *CreateMOUsageDetailExportResponseBody {
	return s.Body
}

func (s *CreateMOUsageDetailExportResponse) SetHeaders(v map[string]*string) *CreateMOUsageDetailExportResponse {
	s.Headers = v
	return s
}

func (s *CreateMOUsageDetailExportResponse) SetStatusCode(v int32) *CreateMOUsageDetailExportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMOUsageDetailExportResponse) SetBody(v *CreateMOUsageDetailExportResponseBody) *CreateMOUsageDetailExportResponse {
	s.Body = v
	return s
}

func (s *CreateMOUsageDetailExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
