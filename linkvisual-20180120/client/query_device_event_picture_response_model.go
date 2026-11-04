// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventPictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceEventPictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceEventPictureResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceEventPictureResponseBody) *QueryDeviceEventPictureResponse
	GetBody() *QueryDeviceEventPictureResponseBody
}

type QueryDeviceEventPictureResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventPictureResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventPictureResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventPictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceEventPictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceEventPictureResponse) GetBody() *QueryDeviceEventPictureResponseBody {
	return s.Body
}

func (s *QueryDeviceEventPictureResponse) SetHeaders(v map[string]*string) *QueryDeviceEventPictureResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventPictureResponse) SetStatusCode(v int32) *QueryDeviceEventPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventPictureResponse) SetBody(v *QueryDeviceEventPictureResponseBody) *QueryDeviceEventPictureResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceEventPictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
