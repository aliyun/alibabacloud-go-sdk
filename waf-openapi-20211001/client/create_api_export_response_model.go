// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiExportResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiExportResponseBody) *CreateApiExportResponse
	GetBody() *CreateApiExportResponseBody
}

type CreateApiExportResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiExportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiExportResponse) GoString() string {
	return s.String()
}

func (s *CreateApiExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiExportResponse) GetBody() *CreateApiExportResponseBody {
	return s.Body
}

func (s *CreateApiExportResponse) SetHeaders(v map[string]*string) *CreateApiExportResponse {
	s.Headers = v
	return s
}

func (s *CreateApiExportResponse) SetStatusCode(v int32) *CreateApiExportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiExportResponse) SetBody(v *CreateApiExportResponseBody) *CreateApiExportResponse {
	s.Body = v
	return s
}

func (s *CreateApiExportResponse) Validate() error {
	return dara.Validate(s)
}
