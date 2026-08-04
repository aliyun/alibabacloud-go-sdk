// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadIdentityRegistrationDocConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadIdentityRegistrationDocConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadIdentityRegistrationDocConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadIdentityRegistrationDocConfigResponseBody) *GetUploadIdentityRegistrationDocConfigResponse
	GetBody() *GetUploadIdentityRegistrationDocConfigResponseBody
}

type GetUploadIdentityRegistrationDocConfigResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadIdentityRegistrationDocConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadIdentityRegistrationDocConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadIdentityRegistrationDocConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) GetBody() *GetUploadIdentityRegistrationDocConfigResponseBody {
	return s.Body
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) SetHeaders(v map[string]*string) *GetUploadIdentityRegistrationDocConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) SetStatusCode(v int32) *GetUploadIdentityRegistrationDocConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) SetBody(v *GetUploadIdentityRegistrationDocConfigResponseBody) *GetUploadIdentityRegistrationDocConfigResponse {
	s.Body = v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
