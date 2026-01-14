// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexMonitorResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexMonitorResponseBody) *GetIndexMonitorResponse
	GetBody() *GetIndexMonitorResponseBody
}

type GetIndexMonitorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexMonitorResponse) GoString() string {
	return s.String()
}

func (s *GetIndexMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexMonitorResponse) GetBody() *GetIndexMonitorResponseBody {
	return s.Body
}

func (s *GetIndexMonitorResponse) SetHeaders(v map[string]*string) *GetIndexMonitorResponse {
	s.Headers = v
	return s
}

func (s *GetIndexMonitorResponse) SetStatusCode(v int32) *GetIndexMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexMonitorResponse) SetBody(v *GetIndexMonitorResponseBody) *GetIndexMonitorResponse {
	s.Body = v
	return s
}

func (s *GetIndexMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
