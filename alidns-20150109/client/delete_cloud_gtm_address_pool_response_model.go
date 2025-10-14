// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudGtmAddressPoolResponseBody) *DeleteCloudGtmAddressPoolResponse
	GetBody() *DeleteCloudGtmAddressPoolResponseBody
}

type DeleteCloudGtmAddressPoolResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudGtmAddressPoolResponse) GetBody() *DeleteCloudGtmAddressPoolResponseBody {
	return s.Body
}

func (s *DeleteCloudGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DeleteCloudGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponse) SetStatusCode(v int32) *DeleteCloudGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponse) SetBody(v *DeleteCloudGtmAddressPoolResponseBody) *DeleteCloudGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
