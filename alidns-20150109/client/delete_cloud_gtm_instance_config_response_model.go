// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudGtmInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudGtmInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudGtmInstanceConfigResponseBody) *DeleteCloudGtmInstanceConfigResponse
	GetBody() *DeleteCloudGtmInstanceConfigResponseBody
}

type DeleteCloudGtmInstanceConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudGtmInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudGtmInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudGtmInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudGtmInstanceConfigResponse) GetBody() *DeleteCloudGtmInstanceConfigResponseBody {
	return s.Body
}

func (s *DeleteCloudGtmInstanceConfigResponse) SetHeaders(v map[string]*string) *DeleteCloudGtmInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudGtmInstanceConfigResponse) SetStatusCode(v int32) *DeleteCloudGtmInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigResponse) SetBody(v *DeleteCloudGtmInstanceConfigResponseBody) *DeleteCloudGtmInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudGtmInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
