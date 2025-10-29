// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoomRealTimeStreamAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoomRealTimeStreamAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoomRealTimeStreamAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoomRealTimeStreamAddressResponseBody) *CreateRoomRealTimeStreamAddressResponse
	GetBody() *CreateRoomRealTimeStreamAddressResponseBody
}

type CreateRoomRealTimeStreamAddressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoomRealTimeStreamAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoomRealTimeStreamAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoomRealTimeStreamAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateRoomRealTimeStreamAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoomRealTimeStreamAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoomRealTimeStreamAddressResponse) GetBody() *CreateRoomRealTimeStreamAddressResponseBody {
	return s.Body
}

func (s *CreateRoomRealTimeStreamAddressResponse) SetHeaders(v map[string]*string) *CreateRoomRealTimeStreamAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateRoomRealTimeStreamAddressResponse) SetStatusCode(v int32) *CreateRoomRealTimeStreamAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressResponse) SetBody(v *CreateRoomRealTimeStreamAddressResponseBody) *CreateRoomRealTimeStreamAddressResponse {
	s.Body = v
	return s
}

func (s *CreateRoomRealTimeStreamAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
