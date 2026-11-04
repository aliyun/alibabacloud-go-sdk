// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceVodUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceVodUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceVodUrlResponseBody) *QueryDeviceVodUrlResponse
	GetBody() *QueryDeviceVodUrlResponseBody
}

type QueryDeviceVodUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceVodUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceVodUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceVodUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceVodUrlResponse) GetBody() *QueryDeviceVodUrlResponseBody {
	return s.Body
}

func (s *QueryDeviceVodUrlResponse) SetHeaders(v map[string]*string) *QueryDeviceVodUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceVodUrlResponse) SetStatusCode(v int32) *QueryDeviceVodUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceVodUrlResponse) SetBody(v *QueryDeviceVodUrlResponseBody) *QueryDeviceVodUrlResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceVodUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
