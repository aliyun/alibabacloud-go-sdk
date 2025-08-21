// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceStatusDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceStatusDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceStatusDetailResponseBody) *GetDeviceStatusDetailResponse
	GetBody() *GetDeviceStatusDetailResponseBody
}

type GetDeviceStatusDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceStatusDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceStatusDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceStatusDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceStatusDetailResponse) GetBody() *GetDeviceStatusDetailResponseBody {
	return s.Body
}

func (s *GetDeviceStatusDetailResponse) SetHeaders(v map[string]*string) *GetDeviceStatusDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceStatusDetailResponse) SetStatusCode(v int32) *GetDeviceStatusDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceStatusDetailResponse) SetBody(v *GetDeviceStatusDetailResponseBody) *GetDeviceStatusDetailResponse {
	s.Body = v
	return s
}

func (s *GetDeviceStatusDetailResponse) Validate() error {
	return dara.Validate(s)
}
