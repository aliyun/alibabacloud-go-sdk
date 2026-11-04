// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceEventResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceEventResponseBody) *QueryDeviceEventResponse
	GetBody() *QueryDeviceEventResponseBody
}

type QueryDeviceEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceEventResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceEventResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceEventResponse) GetBody() *QueryDeviceEventResponseBody {
	return s.Body
}

func (s *QueryDeviceEventResponse) SetHeaders(v map[string]*string) *QueryDeviceEventResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceEventResponse) SetStatusCode(v int32) *QueryDeviceEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceEventResponse) SetBody(v *QueryDeviceEventResponseBody) *QueryDeviceEventResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
