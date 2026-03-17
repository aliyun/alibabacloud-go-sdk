// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedMonitorStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdvancedMonitorStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdvancedMonitorStateResponse
	GetStatusCode() *int32
	SetBody(v *GetAdvancedMonitorStateResponseBody) *GetAdvancedMonitorStateResponse
	GetBody() *GetAdvancedMonitorStateResponseBody
}

type GetAdvancedMonitorStateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvancedMonitorStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvancedMonitorStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedMonitorStateResponse) GoString() string {
	return s.String()
}

func (s *GetAdvancedMonitorStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdvancedMonitorStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdvancedMonitorStateResponse) GetBody() *GetAdvancedMonitorStateResponseBody {
	return s.Body
}

func (s *GetAdvancedMonitorStateResponse) SetHeaders(v map[string]*string) *GetAdvancedMonitorStateResponse {
	s.Headers = v
	return s
}

func (s *GetAdvancedMonitorStateResponse) SetStatusCode(v int32) *GetAdvancedMonitorStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvancedMonitorStateResponse) SetBody(v *GetAdvancedMonitorStateResponseBody) *GetAdvancedMonitorStateResponse {
	s.Body = v
	return s
}

func (s *GetAdvancedMonitorStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
