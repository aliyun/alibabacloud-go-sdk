// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportConfigJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExportConfigJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExportConfigJobResponse
	GetStatusCode() *int32
	SetBody(v *GetExportConfigJobResponseBody) *GetExportConfigJobResponse
	GetBody() *GetExportConfigJobResponseBody
}

type GetExportConfigJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExportConfigJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExportConfigJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExportConfigJobResponse) GoString() string {
	return s.String()
}

func (s *GetExportConfigJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExportConfigJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExportConfigJobResponse) GetBody() *GetExportConfigJobResponseBody {
	return s.Body
}

func (s *GetExportConfigJobResponse) SetHeaders(v map[string]*string) *GetExportConfigJobResponse {
	s.Headers = v
	return s
}

func (s *GetExportConfigJobResponse) SetStatusCode(v int32) *GetExportConfigJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExportConfigJobResponse) SetBody(v *GetExportConfigJobResponseBody) *GetExportConfigJobResponse {
	s.Body = v
	return s
}

func (s *GetExportConfigJobResponse) Validate() error {
	return dara.Validate(s)
}
