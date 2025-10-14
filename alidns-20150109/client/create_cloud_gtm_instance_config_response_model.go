// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudGtmInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudGtmInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudGtmInstanceConfigResponseBody) *CreateCloudGtmInstanceConfigResponse
	GetBody() *CreateCloudGtmInstanceConfigResponseBody
}

type CreateCloudGtmInstanceConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudGtmInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudGtmInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudGtmInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudGtmInstanceConfigResponse) GetBody() *CreateCloudGtmInstanceConfigResponseBody {
	return s.Body
}

func (s *CreateCloudGtmInstanceConfigResponse) SetHeaders(v map[string]*string) *CreateCloudGtmInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponse) SetStatusCode(v int32) *CreateCloudGtmInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponse) SetBody(v *CreateCloudGtmInstanceConfigResponseBody) *CreateCloudGtmInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
