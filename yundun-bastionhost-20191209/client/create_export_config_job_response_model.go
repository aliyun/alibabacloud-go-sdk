// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportConfigJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExportConfigJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExportConfigJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateExportConfigJobResponseBody) *CreateExportConfigJobResponse
	GetBody() *CreateExportConfigJobResponseBody
}

type CreateExportConfigJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExportConfigJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExportConfigJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExportConfigJobResponse) GoString() string {
	return s.String()
}

func (s *CreateExportConfigJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExportConfigJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExportConfigJobResponse) GetBody() *CreateExportConfigJobResponseBody {
	return s.Body
}

func (s *CreateExportConfigJobResponse) SetHeaders(v map[string]*string) *CreateExportConfigJobResponse {
	s.Headers = v
	return s
}

func (s *CreateExportConfigJobResponse) SetStatusCode(v int32) *CreateExportConfigJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExportConfigJobResponse) SetBody(v *CreateExportConfigJobResponseBody) *CreateExportConfigJobResponse {
	s.Body = v
	return s
}

func (s *CreateExportConfigJobResponse) Validate() error {
	return dara.Validate(s)
}
