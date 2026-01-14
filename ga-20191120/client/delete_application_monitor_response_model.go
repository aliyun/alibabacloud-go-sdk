// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationMonitorResponseBody) *DeleteApplicationMonitorResponse
	GetBody() *DeleteApplicationMonitorResponseBody
}

type DeleteApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationMonitorResponse) GetBody() *DeleteApplicationMonitorResponseBody {
	return s.Body
}

func (s *DeleteApplicationMonitorResponse) SetHeaders(v map[string]*string) *DeleteApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationMonitorResponse) SetStatusCode(v int32) *DeleteApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationMonitorResponse) SetBody(v *DeleteApplicationMonitorResponseBody) *DeleteApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
