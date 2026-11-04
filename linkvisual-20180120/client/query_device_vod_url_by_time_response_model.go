// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlByTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceVodUrlByTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceVodUrlByTimeResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceVodUrlByTimeResponseBody) *QueryDeviceVodUrlByTimeResponse
	GetBody() *QueryDeviceVodUrlByTimeResponseBody
}

type QueryDeviceVodUrlByTimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceVodUrlByTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceVodUrlByTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceVodUrlByTimeResponse) GetBody() *QueryDeviceVodUrlByTimeResponseBody {
	return s.Body
}

func (s *QueryDeviceVodUrlByTimeResponse) SetHeaders(v map[string]*string) *QueryDeviceVodUrlByTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponse) SetStatusCode(v int32) *QueryDeviceVodUrlByTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponse) SetBody(v *QueryDeviceVodUrlByTimeResponseBody) *QueryDeviceVodUrlByTimeResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
