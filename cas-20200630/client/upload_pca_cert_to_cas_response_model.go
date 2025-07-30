// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPcaCertToCasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPcaCertToCasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPcaCertToCasResponse
	GetStatusCode() *int32
	SetBody(v *UploadPcaCertToCasResponseBody) *UploadPcaCertToCasResponse
	GetBody() *UploadPcaCertToCasResponseBody
}

type UploadPcaCertToCasResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPcaCertToCasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPcaCertToCasResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPcaCertToCasResponse) GoString() string {
	return s.String()
}

func (s *UploadPcaCertToCasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPcaCertToCasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPcaCertToCasResponse) GetBody() *UploadPcaCertToCasResponseBody {
	return s.Body
}

func (s *UploadPcaCertToCasResponse) SetHeaders(v map[string]*string) *UploadPcaCertToCasResponse {
	s.Headers = v
	return s
}

func (s *UploadPcaCertToCasResponse) SetStatusCode(v int32) *UploadPcaCertToCasResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPcaCertToCasResponse) SetBody(v *UploadPcaCertToCasResponseBody) *UploadPcaCertToCasResponse {
	s.Body = v
	return s
}

func (s *UploadPcaCertToCasResponse) Validate() error {
	return dara.Validate(s)
}
