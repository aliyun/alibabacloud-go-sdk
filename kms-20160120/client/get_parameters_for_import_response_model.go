// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersForImportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParametersForImportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParametersForImportResponse
	GetStatusCode() *int32
	SetBody(v *GetParametersForImportResponseBody) *GetParametersForImportResponse
	GetBody() *GetParametersForImportResponseBody
}

type GetParametersForImportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersForImportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersForImportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParametersForImportResponse) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParametersForImportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParametersForImportResponse) GetBody() *GetParametersForImportResponseBody {
	return s.Body
}

func (s *GetParametersForImportResponse) SetHeaders(v map[string]*string) *GetParametersForImportResponse {
	s.Headers = v
	return s
}

func (s *GetParametersForImportResponse) SetStatusCode(v int32) *GetParametersForImportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersForImportResponse) SetBody(v *GetParametersForImportResponseBody) *GetParametersForImportResponse {
	s.Body = v
	return s
}

func (s *GetParametersForImportResponse) Validate() error {
	return dara.Validate(s)
}
