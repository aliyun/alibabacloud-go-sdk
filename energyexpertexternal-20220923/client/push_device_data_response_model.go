// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushDeviceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushDeviceDataResponse
	GetStatusCode() *int32
	SetBody(v *PushDeviceDataResponseBody) *PushDeviceDataResponse
	GetBody() *PushDeviceDataResponseBody
}

type PushDeviceDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushDeviceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushDeviceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceDataResponse) GoString() string {
	return s.String()
}

func (s *PushDeviceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushDeviceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushDeviceDataResponse) GetBody() *PushDeviceDataResponseBody {
	return s.Body
}

func (s *PushDeviceDataResponse) SetHeaders(v map[string]*string) *PushDeviceDataResponse {
	s.Headers = v
	return s
}

func (s *PushDeviceDataResponse) SetStatusCode(v int32) *PushDeviceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDeviceDataResponse) SetBody(v *PushDeviceDataResponseBody) *PushDeviceDataResponse {
	s.Body = v
	return s
}

func (s *PushDeviceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
