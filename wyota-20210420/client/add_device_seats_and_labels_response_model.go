// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceSeatsAndLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDeviceSeatsAndLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDeviceSeatsAndLabelsResponse
	GetStatusCode() *int32
	SetBody(v *AddDeviceSeatsAndLabelsResponseBody) *AddDeviceSeatsAndLabelsResponse
	GetBody() *AddDeviceSeatsAndLabelsResponseBody
}

type AddDeviceSeatsAndLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceSeatsAndLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceSeatsAndLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceSeatsAndLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceSeatsAndLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDeviceSeatsAndLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDeviceSeatsAndLabelsResponse) GetBody() *AddDeviceSeatsAndLabelsResponseBody {
	return s.Body
}

func (s *AddDeviceSeatsAndLabelsResponse) SetHeaders(v map[string]*string) *AddDeviceSeatsAndLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponse) SetStatusCode(v int32) *AddDeviceSeatsAndLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponse) SetBody(v *AddDeviceSeatsAndLabelsResponseBody) *AddDeviceSeatsAndLabelsResponse {
	s.Body = v
	return s
}

func (s *AddDeviceSeatsAndLabelsResponse) Validate() error {
	return dara.Validate(s)
}
