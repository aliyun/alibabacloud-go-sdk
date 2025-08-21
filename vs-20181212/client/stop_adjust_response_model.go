// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdjustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAdjustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAdjustResponse
	GetStatusCode() *int32
	SetBody(v *StopAdjustResponseBody) *StopAdjustResponse
	GetBody() *StopAdjustResponseBody
}

type StopAdjustResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAdjustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAdjustResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAdjustResponse) GoString() string {
	return s.String()
}

func (s *StopAdjustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAdjustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAdjustResponse) GetBody() *StopAdjustResponseBody {
	return s.Body
}

func (s *StopAdjustResponse) SetHeaders(v map[string]*string) *StopAdjustResponse {
	s.Headers = v
	return s
}

func (s *StopAdjustResponse) SetStatusCode(v int32) *StopAdjustResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAdjustResponse) SetBody(v *StopAdjustResponseBody) *StopAdjustResponse {
	s.Body = v
	return s
}

func (s *StopAdjustResponse) Validate() error {
	return dara.Validate(s)
}
