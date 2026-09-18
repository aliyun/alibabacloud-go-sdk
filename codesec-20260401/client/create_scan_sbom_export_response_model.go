// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanSbomExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScanSbomExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScanSbomExportResponse
	GetStatusCode() *int32
	SetBody(v *CreateScanSbomExportResponseBody) *CreateScanSbomExportResponse
	GetBody() *CreateScanSbomExportResponseBody
}

type CreateScanSbomExportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScanSbomExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScanSbomExportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScanSbomExportResponse) GoString() string {
	return s.String()
}

func (s *CreateScanSbomExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScanSbomExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScanSbomExportResponse) GetBody() *CreateScanSbomExportResponseBody {
	return s.Body
}

func (s *CreateScanSbomExportResponse) SetHeaders(v map[string]*string) *CreateScanSbomExportResponse {
	s.Headers = v
	return s
}

func (s *CreateScanSbomExportResponse) SetStatusCode(v int32) *CreateScanSbomExportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScanSbomExportResponse) SetBody(v *CreateScanSbomExportResponseBody) *CreateScanSbomExportResponse {
	s.Body = v
	return s
}

func (s *CreateScanSbomExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
