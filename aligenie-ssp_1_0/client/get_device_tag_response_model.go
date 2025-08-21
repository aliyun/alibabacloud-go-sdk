// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceTagResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceTagResponseBody) *GetDeviceTagResponse
	GetBody() *GetDeviceTagResponseBody
}

type GetDeviceTagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceTagResponse) GetBody() *GetDeviceTagResponseBody {
	return s.Body
}

func (s *GetDeviceTagResponse) SetHeaders(v map[string]*string) *GetDeviceTagResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceTagResponse) SetStatusCode(v int32) *GetDeviceTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceTagResponse) SetBody(v *GetDeviceTagResponseBody) *GetDeviceTagResponse {
	s.Body = v
	return s
}

func (s *GetDeviceTagResponse) Validate() error {
	return dara.Validate(s)
}
