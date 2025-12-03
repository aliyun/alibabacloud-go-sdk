// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UploadDeviceResponseBody) *UploadDeviceResponse
	GetBody() *UploadDeviceResponseBody
}

type UploadDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDeviceResponse) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDeviceResponse) GetBody() *UploadDeviceResponseBody {
	return s.Body
}

func (s *UploadDeviceResponse) SetHeaders(v map[string]*string) *UploadDeviceResponse {
	s.Headers = v
	return s
}

func (s *UploadDeviceResponse) SetStatusCode(v int32) *UploadDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDeviceResponse) SetBody(v *UploadDeviceResponseBody) *UploadDeviceResponse {
	s.Body = v
	return s
}

func (s *UploadDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
