// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceBelongIdentifiersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDeviceBelongIdentifiersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDeviceBelongIdentifiersResponse
	GetStatusCode() *int32
	SetBody(v *AddDeviceBelongIdentifiersResponseBody) *AddDeviceBelongIdentifiersResponse
	GetBody() *AddDeviceBelongIdentifiersResponseBody
}

type AddDeviceBelongIdentifiersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceBelongIdentifiersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceBelongIdentifiersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceBelongIdentifiersResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceBelongIdentifiersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDeviceBelongIdentifiersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDeviceBelongIdentifiersResponse) GetBody() *AddDeviceBelongIdentifiersResponseBody {
	return s.Body
}

func (s *AddDeviceBelongIdentifiersResponse) SetHeaders(v map[string]*string) *AddDeviceBelongIdentifiersResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceBelongIdentifiersResponse) SetStatusCode(v int32) *AddDeviceBelongIdentifiersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceBelongIdentifiersResponse) SetBody(v *AddDeviceBelongIdentifiersResponseBody) *AddDeviceBelongIdentifiersResponse {
	s.Body = v
	return s
}

func (s *AddDeviceBelongIdentifiersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
