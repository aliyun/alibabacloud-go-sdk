// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByDeviceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserByDeviceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserByDeviceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserByDeviceIdResponseBody) *GetUserByDeviceIdResponse
	GetBody() *GetUserByDeviceIdResponseBody
}

type GetUserByDeviceIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserByDeviceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserByDeviceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserByDeviceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserByDeviceIdResponse) GetBody() *GetUserByDeviceIdResponseBody {
	return s.Body
}

func (s *GetUserByDeviceIdResponse) SetHeaders(v map[string]*string) *GetUserByDeviceIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserByDeviceIdResponse) SetStatusCode(v int32) *GetUserByDeviceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserByDeviceIdResponse) SetBody(v *GetUserByDeviceIdResponseBody) *GetUserByDeviceIdResponse {
	s.Body = v
	return s
}

func (s *GetUserByDeviceIdResponse) Validate() error {
	return dara.Validate(s)
}
