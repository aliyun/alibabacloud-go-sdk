// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceStatResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceStatResponseBody) *QueryDeviceStatResponse
	GetBody() *QueryDeviceStatResponseBody
}

type QueryDeviceStatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceStatResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceStatResponse) GetBody() *QueryDeviceStatResponseBody {
	return s.Body
}

func (s *QueryDeviceStatResponse) SetHeaders(v map[string]*string) *QueryDeviceStatResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatResponse) SetStatusCode(v int32) *QueryDeviceStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatResponse) SetBody(v *QueryDeviceStatResponseBody) *QueryDeviceStatResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceStatResponse) Validate() error {
	return dara.Validate(s)
}
