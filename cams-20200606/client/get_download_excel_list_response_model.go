// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadExcelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDownloadExcelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDownloadExcelListResponse
	GetStatusCode() *int32
	SetBody(v *GetDownloadExcelListResponseBody) *GetDownloadExcelListResponse
	GetBody() *GetDownloadExcelListResponseBody
}

type GetDownloadExcelListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDownloadExcelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDownloadExcelListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadExcelListResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadExcelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDownloadExcelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDownloadExcelListResponse) GetBody() *GetDownloadExcelListResponseBody {
	return s.Body
}

func (s *GetDownloadExcelListResponse) SetHeaders(v map[string]*string) *GetDownloadExcelListResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadExcelListResponse) SetStatusCode(v int32) *GetDownloadExcelListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDownloadExcelListResponse) SetBody(v *GetDownloadExcelListResponseBody) *GetDownloadExcelListResponse {
	s.Body = v
	return s
}

func (s *GetDownloadExcelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
