// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHighDefinitionMonitorLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHighDefinitionMonitorLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHighDefinitionMonitorLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetHighDefinitionMonitorLogStatusResponseBody) *SetHighDefinitionMonitorLogStatusResponse
	GetBody() *SetHighDefinitionMonitorLogStatusResponseBody
}

type SetHighDefinitionMonitorLogStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHighDefinitionMonitorLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHighDefinitionMonitorLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHighDefinitionMonitorLogStatusResponse) GoString() string {
	return s.String()
}

func (s *SetHighDefinitionMonitorLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHighDefinitionMonitorLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHighDefinitionMonitorLogStatusResponse) GetBody() *SetHighDefinitionMonitorLogStatusResponseBody {
	return s.Body
}

func (s *SetHighDefinitionMonitorLogStatusResponse) SetHeaders(v map[string]*string) *SetHighDefinitionMonitorLogStatusResponse {
	s.Headers = v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusResponse) SetStatusCode(v int32) *SetHighDefinitionMonitorLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusResponse) SetBody(v *SetHighDefinitionMonitorLogStatusResponseBody) *SetHighDefinitionMonitorLogStatusResponse {
	s.Body = v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
