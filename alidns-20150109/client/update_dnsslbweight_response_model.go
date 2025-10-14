// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNSSLBWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDNSSLBWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDNSSLBWeightResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDNSSLBWeightResponseBody) *UpdateDNSSLBWeightResponse
	GetBody() *UpdateDNSSLBWeightResponseBody
}

type UpdateDNSSLBWeightResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDNSSLBWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDNSSLBWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNSSLBWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDNSSLBWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDNSSLBWeightResponse) GetBody() *UpdateDNSSLBWeightResponseBody {
	return s.Body
}

func (s *UpdateDNSSLBWeightResponse) SetHeaders(v map[string]*string) *UpdateDNSSLBWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateDNSSLBWeightResponse) SetStatusCode(v int32) *UpdateDNSSLBWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDNSSLBWeightResponse) SetBody(v *UpdateDNSSLBWeightResponseBody) *UpdateDNSSLBWeightResponse {
	s.Body = v
	return s
}

func (s *UpdateDNSSLBWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
