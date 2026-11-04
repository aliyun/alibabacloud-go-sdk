// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPictureSearchAppWithDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPictureSearchAppWithDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPictureSearchAppWithDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BindPictureSearchAppWithDevicesResponseBody) *BindPictureSearchAppWithDevicesResponse
	GetBody() *BindPictureSearchAppWithDevicesResponseBody
}

type BindPictureSearchAppWithDevicesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPictureSearchAppWithDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPictureSearchAppWithDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPictureSearchAppWithDevicesResponse) GoString() string {
	return s.String()
}

func (s *BindPictureSearchAppWithDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPictureSearchAppWithDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPictureSearchAppWithDevicesResponse) GetBody() *BindPictureSearchAppWithDevicesResponseBody {
	return s.Body
}

func (s *BindPictureSearchAppWithDevicesResponse) SetHeaders(v map[string]*string) *BindPictureSearchAppWithDevicesResponse {
	s.Headers = v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponse) SetStatusCode(v int32) *BindPictureSearchAppWithDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponse) SetBody(v *BindPictureSearchAppWithDevicesResponseBody) *BindPictureSearchAppWithDevicesResponse {
	s.Body = v
	return s
}

func (s *BindPictureSearchAppWithDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
