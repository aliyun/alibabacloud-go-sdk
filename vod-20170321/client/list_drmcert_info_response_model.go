// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDRMCertInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDRMCertInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDRMCertInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListDRMCertInfoResponseBody) *ListDRMCertInfoResponse
	GetBody() *ListDRMCertInfoResponseBody
}

type ListDRMCertInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDRMCertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDRMCertInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDRMCertInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDRMCertInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDRMCertInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDRMCertInfoResponse) GetBody() *ListDRMCertInfoResponseBody {
	return s.Body
}

func (s *ListDRMCertInfoResponse) SetHeaders(v map[string]*string) *ListDRMCertInfoResponse {
	s.Headers = v
	return s
}

func (s *ListDRMCertInfoResponse) SetStatusCode(v int32) *ListDRMCertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDRMCertInfoResponse) SetBody(v *ListDRMCertInfoResponseBody) *ListDRMCertInfoResponse {
	s.Body = v
	return s
}

func (s *ListDRMCertInfoResponse) Validate() error {
	return dara.Validate(s)
}
