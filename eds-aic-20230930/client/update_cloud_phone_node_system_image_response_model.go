// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPhoneNodeSystemImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudPhoneNodeSystemImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudPhoneNodeSystemImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudPhoneNodeSystemImageResponseBody) *UpdateCloudPhoneNodeSystemImageResponse
	GetBody() *UpdateCloudPhoneNodeSystemImageResponseBody
}

type UpdateCloudPhoneNodeSystemImageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudPhoneNodeSystemImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudPhoneNodeSystemImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPhoneNodeSystemImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) GetBody() *UpdateCloudPhoneNodeSystemImageResponseBody {
	return s.Body
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) SetHeaders(v map[string]*string) *UpdateCloudPhoneNodeSystemImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) SetStatusCode(v int32) *UpdateCloudPhoneNodeSystemImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) SetBody(v *UpdateCloudPhoneNodeSystemImageResponseBody) *UpdateCloudPhoneNodeSystemImageResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
