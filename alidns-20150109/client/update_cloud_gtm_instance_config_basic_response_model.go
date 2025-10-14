// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigBasicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigBasicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigBasicResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceConfigBasicResponseBody) *UpdateCloudGtmInstanceConfigBasicResponse
	GetBody() *UpdateCloudGtmInstanceConfigBasicResponseBody
}

type UpdateCloudGtmInstanceConfigBasicResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceConfigBasicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigBasicResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigBasicResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) GetBody() *UpdateCloudGtmInstanceConfigBasicResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigBasicResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigBasicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) SetBody(v *UpdateCloudGtmInstanceConfigBasicResponseBody) *UpdateCloudGtmInstanceConfigBasicResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
