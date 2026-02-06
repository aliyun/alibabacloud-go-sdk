// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceExtensionCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudResourceExtensionCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudResourceExtensionCertResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudResourceExtensionCertResponseBody) *DeleteCloudResourceExtensionCertResponse
	GetBody() *DeleteCloudResourceExtensionCertResponseBody
}

type DeleteCloudResourceExtensionCertResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudResourceExtensionCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudResourceExtensionCertResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceExtensionCertResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceExtensionCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudResourceExtensionCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudResourceExtensionCertResponse) GetBody() *DeleteCloudResourceExtensionCertResponseBody {
	return s.Body
}

func (s *DeleteCloudResourceExtensionCertResponse) SetHeaders(v map[string]*string) *DeleteCloudResourceExtensionCertResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudResourceExtensionCertResponse) SetStatusCode(v int32) *DeleteCloudResourceExtensionCertResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertResponse) SetBody(v *DeleteCloudResourceExtensionCertResponseBody) *DeleteCloudResourceExtensionCertResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudResourceExtensionCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
