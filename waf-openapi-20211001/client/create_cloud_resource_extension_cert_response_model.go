// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceExtensionCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudResourceExtensionCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudResourceExtensionCertResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudResourceExtensionCertResponseBody) *CreateCloudResourceExtensionCertResponse
	GetBody() *CreateCloudResourceExtensionCertResponseBody
}

type CreateCloudResourceExtensionCertResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudResourceExtensionCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudResourceExtensionCertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceExtensionCertResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceExtensionCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudResourceExtensionCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudResourceExtensionCertResponse) GetBody() *CreateCloudResourceExtensionCertResponseBody {
	return s.Body
}

func (s *CreateCloudResourceExtensionCertResponse) SetHeaders(v map[string]*string) *CreateCloudResourceExtensionCertResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudResourceExtensionCertResponse) SetStatusCode(v int32) *CreateCloudResourceExtensionCertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudResourceExtensionCertResponse) SetBody(v *CreateCloudResourceExtensionCertResponseBody) *CreateCloudResourceExtensionCertResponse {
	s.Body = v
	return s
}

func (s *CreateCloudResourceExtensionCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
