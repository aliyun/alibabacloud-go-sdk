// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataExportOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataExportOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataExportOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataExportOrderDetailResponseBody) *GetDataExportOrderDetailResponse
	GetBody() *GetDataExportOrderDetailResponseBody
}

type GetDataExportOrderDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataExportOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataExportOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataExportOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataExportOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataExportOrderDetailResponse) GetBody() *GetDataExportOrderDetailResponseBody {
	return s.Body
}

func (s *GetDataExportOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataExportOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataExportOrderDetailResponse) SetStatusCode(v int32) *GetDataExportOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataExportOrderDetailResponse) SetBody(v *GetDataExportOrderDetailResponseBody) *GetDataExportOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataExportOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
