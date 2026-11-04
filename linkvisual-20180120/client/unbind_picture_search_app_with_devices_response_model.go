// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPictureSearchAppWithDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindPictureSearchAppWithDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindPictureSearchAppWithDevicesResponse
	GetStatusCode() *int32
	SetBody(v *UnbindPictureSearchAppWithDevicesResponseBody) *UnbindPictureSearchAppWithDevicesResponse
	GetBody() *UnbindPictureSearchAppWithDevicesResponseBody
}

type UnbindPictureSearchAppWithDevicesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPictureSearchAppWithDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindPictureSearchAppWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *UnbindPictureSearchAppWithDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindPictureSearchAppWithDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindPictureSearchAppWithDevicesResponse) GetBody() *UnbindPictureSearchAppWithDevicesResponseBody {
	return s.Body
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetHeaders(v map[string]*string) *UnbindPictureSearchAppWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetStatusCode(v int32) *UnbindPictureSearchAppWithDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponse) SetBody(v *UnbindPictureSearchAppWithDevicesResponseBody) *UnbindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

func (s *UnbindPictureSearchAppWithDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
