// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImportFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImportFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImportFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetImportFileUrlResponseBody) *GetImportFileUrlResponse
	GetBody() *GetImportFileUrlResponseBody
}

type GetImportFileUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImportFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImportFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImportFileUrlResponse) GetBody() *GetImportFileUrlResponseBody {
	return s.Body
}

func (s *GetImportFileUrlResponse) SetHeaders(v map[string]*string) *GetImportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetImportFileUrlResponse) SetStatusCode(v int32) *GetImportFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImportFileUrlResponse) SetBody(v *GetImportFileUrlResponseBody) *GetImportFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetImportFileUrlResponse) Validate() error {
	return dara.Validate(s)
}
