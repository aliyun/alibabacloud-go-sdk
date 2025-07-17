// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseExportOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseExportOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseExportOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabaseExportOrderDetailResponseBody) *GetDatabaseExportOrderDetailResponse
	GetBody() *GetDatabaseExportOrderDetailResponseBody
}

type GetDatabaseExportOrderDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseExportOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseExportOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseExportOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseExportOrderDetailResponse) GetBody() *GetDatabaseExportOrderDetailResponseBody {
	return s.Body
}

func (s *GetDatabaseExportOrderDetailResponse) SetHeaders(v map[string]*string) *GetDatabaseExportOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponse) SetStatusCode(v int32) *GetDatabaseExportOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponse) SetBody(v *GetDatabaseExportOrderDetailResponseBody) *GetDatabaseExportOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
