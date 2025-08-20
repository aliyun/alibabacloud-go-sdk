// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WakeUpAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WakeUpAppResponse
	GetStatusCode() *int32
}

type WakeUpAppResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s WakeUpAppResponse) String() string {
	return dara.Prettify(s)
}

func (s WakeUpAppResponse) GoString() string {
	return s.String()
}

func (s *WakeUpAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WakeUpAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WakeUpAppResponse) SetHeaders(v map[string]*string) *WakeUpAppResponse {
	s.Headers = v
	return s
}

func (s *WakeUpAppResponse) SetStatusCode(v int32) *WakeUpAppResponse {
	s.StatusCode = &v
	return s
}

func (s *WakeUpAppResponse) Validate() error {
	return dara.Validate(s)
}
