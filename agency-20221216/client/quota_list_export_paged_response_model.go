// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaListExportPagedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuotaListExportPagedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuotaListExportPagedResponse
	GetStatusCode() *int32
	SetBody(v *QuotaListExportPagedResponseBody) *QuotaListExportPagedResponse
	GetBody() *QuotaListExportPagedResponseBody
}

type QuotaListExportPagedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuotaListExportPagedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuotaListExportPagedResponse) String() string {
	return dara.Prettify(s)
}

func (s QuotaListExportPagedResponse) GoString() string {
	return s.String()
}

func (s *QuotaListExportPagedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuotaListExportPagedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuotaListExportPagedResponse) GetBody() *QuotaListExportPagedResponseBody {
	return s.Body
}

func (s *QuotaListExportPagedResponse) SetHeaders(v map[string]*string) *QuotaListExportPagedResponse {
	s.Headers = v
	return s
}

func (s *QuotaListExportPagedResponse) SetStatusCode(v int32) *QuotaListExportPagedResponse {
	s.StatusCode = &v
	return s
}

func (s *QuotaListExportPagedResponse) SetBody(v *QuotaListExportPagedResponseBody) *QuotaListExportPagedResponse {
	s.Body = v
	return s
}

func (s *QuotaListExportPagedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
