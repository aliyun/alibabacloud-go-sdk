// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDRMCertInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDRMCertInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDRMCertInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDRMCertInfoResponseBody) *DeleteDRMCertInfoResponse
	GetBody() *DeleteDRMCertInfoResponseBody
}

type DeleteDRMCertInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDRMCertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDRMCertInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDRMCertInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteDRMCertInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDRMCertInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDRMCertInfoResponse) GetBody() *DeleteDRMCertInfoResponseBody {
	return s.Body
}

func (s *DeleteDRMCertInfoResponse) SetHeaders(v map[string]*string) *DeleteDRMCertInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteDRMCertInfoResponse) SetStatusCode(v int32) *DeleteDRMCertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDRMCertInfoResponse) SetBody(v *DeleteDRMCertInfoResponseBody) *DeleteDRMCertInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteDRMCertInfoResponse) Validate() error {
	return dara.Validate(s)
}
