// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportPreCheckDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataExportPreCheckDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataExportPreCheckDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataExportPreCheckDetailResponseBody) *GetDataExportPreCheckDetailResponse
	GetBody() *GetDataExportPreCheckDetailResponseBody
}

type GetDataExportPreCheckDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataExportPreCheckDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataExportPreCheckDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportPreCheckDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataExportPreCheckDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataExportPreCheckDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataExportPreCheckDetailResponse) GetBody() *GetDataExportPreCheckDetailResponseBody {
	return s.Body
}

func (s *GetDataExportPreCheckDetailResponse) SetHeaders(v map[string]*string) *GetDataExportPreCheckDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataExportPreCheckDetailResponse) SetStatusCode(v int32) *GetDataExportPreCheckDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataExportPreCheckDetailResponse) SetBody(v *GetDataExportPreCheckDetailResponseBody) *GetDataExportPreCheckDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataExportPreCheckDetailResponse) Validate() error {
	return dara.Validate(s)
}
