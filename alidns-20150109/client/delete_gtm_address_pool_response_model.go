// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGtmAddressPoolResponseBody) *DeleteGtmAddressPoolResponse
	GetBody() *DeleteGtmAddressPoolResponseBody
}

type DeleteGtmAddressPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGtmAddressPoolResponse) GetBody() *DeleteGtmAddressPoolResponseBody {
	return s.Body
}

func (s *DeleteGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DeleteGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteGtmAddressPoolResponse) SetStatusCode(v int32) *DeleteGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGtmAddressPoolResponse) SetBody(v *DeleteGtmAddressPoolResponseBody) *DeleteGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
