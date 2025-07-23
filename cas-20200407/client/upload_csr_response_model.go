// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadCsrResponse
	GetStatusCode() *int32
	SetBody(v *UploadCsrResponseBody) *UploadCsrResponse
	GetBody() *UploadCsrResponseBody
}

type UploadCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadCsrResponse) GoString() string {
	return s.String()
}

func (s *UploadCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadCsrResponse) GetBody() *UploadCsrResponseBody {
	return s.Body
}

func (s *UploadCsrResponse) SetHeaders(v map[string]*string) *UploadCsrResponse {
	s.Headers = v
	return s
}

func (s *UploadCsrResponse) SetStatusCode(v int32) *UploadCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCsrResponse) SetBody(v *UploadCsrResponseBody) *UploadCsrResponse {
	s.Body = v
	return s
}

func (s *UploadCsrResponse) Validate() error {
	return dara.Validate(s)
}
