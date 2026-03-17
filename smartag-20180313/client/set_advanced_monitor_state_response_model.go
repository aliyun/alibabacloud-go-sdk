// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdvancedMonitorStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAdvancedMonitorStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAdvancedMonitorStateResponse
	GetStatusCode() *int32
	SetBody(v *SetAdvancedMonitorStateResponseBody) *SetAdvancedMonitorStateResponse
	GetBody() *SetAdvancedMonitorStateResponseBody
}

type SetAdvancedMonitorStateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAdvancedMonitorStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAdvancedMonitorStateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAdvancedMonitorStateResponse) GoString() string {
	return s.String()
}

func (s *SetAdvancedMonitorStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAdvancedMonitorStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAdvancedMonitorStateResponse) GetBody() *SetAdvancedMonitorStateResponseBody {
	return s.Body
}

func (s *SetAdvancedMonitorStateResponse) SetHeaders(v map[string]*string) *SetAdvancedMonitorStateResponse {
	s.Headers = v
	return s
}

func (s *SetAdvancedMonitorStateResponse) SetStatusCode(v int32) *SetAdvancedMonitorStateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAdvancedMonitorStateResponse) SetBody(v *SetAdvancedMonitorStateResponseBody) *SetAdvancedMonitorStateResponse {
	s.Body = v
	return s
}

func (s *SetAdvancedMonitorStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
