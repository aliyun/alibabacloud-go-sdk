// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesByRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCertificatesByRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCertificatesByRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListCertificatesByRecordResponseBody) *ListCertificatesByRecordResponse
	GetBody() *ListCertificatesByRecordResponseBody
}

type ListCertificatesByRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertificatesByRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertificatesByRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesByRecordResponse) GoString() string {
	return s.String()
}

func (s *ListCertificatesByRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCertificatesByRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCertificatesByRecordResponse) GetBody() *ListCertificatesByRecordResponseBody {
	return s.Body
}

func (s *ListCertificatesByRecordResponse) SetHeaders(v map[string]*string) *ListCertificatesByRecordResponse {
	s.Headers = v
	return s
}

func (s *ListCertificatesByRecordResponse) SetStatusCode(v int32) *ListCertificatesByRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertificatesByRecordResponse) SetBody(v *ListCertificatesByRecordResponseBody) *ListCertificatesByRecordResponse {
	s.Body = v
	return s
}

func (s *ListCertificatesByRecordResponse) Validate() error {
	return dara.Validate(s)
}
