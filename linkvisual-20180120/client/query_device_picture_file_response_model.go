// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicePictureFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDevicePictureFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDevicePictureFileResponse
	GetStatusCode() *int32
	SetBody(v *QueryDevicePictureFileResponseBody) *QueryDevicePictureFileResponse
	GetBody() *QueryDevicePictureFileResponseBody
}

type QueryDevicePictureFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePictureFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicePictureFileResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicePictureFileResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePictureFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDevicePictureFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDevicePictureFileResponse) GetBody() *QueryDevicePictureFileResponseBody {
	return s.Body
}

func (s *QueryDevicePictureFileResponse) SetHeaders(v map[string]*string) *QueryDevicePictureFileResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePictureFileResponse) SetStatusCode(v int32) *QueryDevicePictureFileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePictureFileResponse) SetBody(v *QueryDevicePictureFileResponseBody) *QueryDevicePictureFileResponse {
	s.Body = v
	return s
}

func (s *QueryDevicePictureFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
