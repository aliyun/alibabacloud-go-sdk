// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLiveShiftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseLiveShiftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseLiveShiftResponse
	GetStatusCode() *int32
	SetBody(v *CloseLiveShiftResponseBody) *CloseLiveShiftResponse
	GetBody() *CloseLiveShiftResponseBody
}

type CloseLiveShiftResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseLiveShiftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseLiveShiftResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseLiveShiftResponse) GoString() string {
	return s.String()
}

func (s *CloseLiveShiftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseLiveShiftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseLiveShiftResponse) GetBody() *CloseLiveShiftResponseBody {
	return s.Body
}

func (s *CloseLiveShiftResponse) SetHeaders(v map[string]*string) *CloseLiveShiftResponse {
	s.Headers = v
	return s
}

func (s *CloseLiveShiftResponse) SetStatusCode(v int32) *CloseLiveShiftResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseLiveShiftResponse) SetBody(v *CloseLiveShiftResponseBody) *CloseLiveShiftResponse {
	s.Body = v
	return s
}

func (s *CloseLiveShiftResponse) Validate() error {
	return dara.Validate(s)
}
