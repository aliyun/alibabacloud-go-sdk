// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMCertInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDRMCertInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDRMCertInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDRMCertInfoResponseBody) *GetDRMCertInfoResponse
	GetBody() *GetDRMCertInfoResponseBody
}

type GetDRMCertInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDRMCertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDRMCertInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDRMCertInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDRMCertInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDRMCertInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDRMCertInfoResponse) GetBody() *GetDRMCertInfoResponseBody {
	return s.Body
}

func (s *GetDRMCertInfoResponse) SetHeaders(v map[string]*string) *GetDRMCertInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDRMCertInfoResponse) SetStatusCode(v int32) *GetDRMCertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDRMCertInfoResponse) SetBody(v *GetDRMCertInfoResponseBody) *GetDRMCertInfoResponse {
	s.Body = v
	return s
}

func (s *GetDRMCertInfoResponse) Validate() error {
	return dara.Validate(s)
}
