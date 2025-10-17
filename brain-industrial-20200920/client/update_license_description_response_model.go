// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLicenseDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLicenseDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLicenseDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLicenseDescriptionResponseBody) *UpdateLicenseDescriptionResponse
	GetBody() *UpdateLicenseDescriptionResponseBody
}

type UpdateLicenseDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLicenseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLicenseDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLicenseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLicenseDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLicenseDescriptionResponse) GetBody() *UpdateLicenseDescriptionResponseBody {
	return s.Body
}

func (s *UpdateLicenseDescriptionResponse) SetHeaders(v map[string]*string) *UpdateLicenseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLicenseDescriptionResponse) SetStatusCode(v int32) *UpdateLicenseDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLicenseDescriptionResponse) SetBody(v *UpdateLicenseDescriptionResponseBody) *UpdateLicenseDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateLicenseDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
