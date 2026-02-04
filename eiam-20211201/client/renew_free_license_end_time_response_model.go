// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseEndTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewFreeLicenseEndTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewFreeLicenseEndTimeResponse
	GetStatusCode() *int32
	SetBody(v *RenewFreeLicenseEndTimeResponseBody) *RenewFreeLicenseEndTimeResponse
	GetBody() *RenewFreeLicenseEndTimeResponseBody
}

type RenewFreeLicenseEndTimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewFreeLicenseEndTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewFreeLicenseEndTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseEndTimeResponse) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseEndTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewFreeLicenseEndTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewFreeLicenseEndTimeResponse) GetBody() *RenewFreeLicenseEndTimeResponseBody {
	return s.Body
}

func (s *RenewFreeLicenseEndTimeResponse) SetHeaders(v map[string]*string) *RenewFreeLicenseEndTimeResponse {
	s.Headers = v
	return s
}

func (s *RenewFreeLicenseEndTimeResponse) SetStatusCode(v int32) *RenewFreeLicenseEndTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewFreeLicenseEndTimeResponse) SetBody(v *RenewFreeLicenseEndTimeResponseBody) *RenewFreeLicenseEndTimeResponse {
	s.Body = v
	return s
}

func (s *RenewFreeLicenseEndTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
