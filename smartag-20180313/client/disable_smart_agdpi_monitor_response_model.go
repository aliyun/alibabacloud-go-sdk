// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAGDpiMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSmartAGDpiMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSmartAGDpiMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DisableSmartAGDpiMonitorResponseBody) *DisableSmartAGDpiMonitorResponse
	GetBody() *DisableSmartAGDpiMonitorResponseBody
}

type DisableSmartAGDpiMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSmartAGDpiMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSmartAGDpiMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAGDpiMonitorResponse) GoString() string {
	return s.String()
}

func (s *DisableSmartAGDpiMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSmartAGDpiMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSmartAGDpiMonitorResponse) GetBody() *DisableSmartAGDpiMonitorResponseBody {
	return s.Body
}

func (s *DisableSmartAGDpiMonitorResponse) SetHeaders(v map[string]*string) *DisableSmartAGDpiMonitorResponse {
	s.Headers = v
	return s
}

func (s *DisableSmartAGDpiMonitorResponse) SetStatusCode(v int32) *DisableSmartAGDpiMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSmartAGDpiMonitorResponse) SetBody(v *DisableSmartAGDpiMonitorResponseBody) *DisableSmartAGDpiMonitorResponse {
	s.Body = v
	return s
}

func (s *DisableSmartAGDpiMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
