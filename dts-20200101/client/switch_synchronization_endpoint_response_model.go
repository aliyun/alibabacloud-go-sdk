// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSynchronizationEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchSynchronizationEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchSynchronizationEndpointResponse
	GetStatusCode() *int32
	SetBody(v *SwitchSynchronizationEndpointResponseBody) *SwitchSynchronizationEndpointResponse
	GetBody() *SwitchSynchronizationEndpointResponseBody
}

type SwitchSynchronizationEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchSynchronizationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchSynchronizationEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponse) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchSynchronizationEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchSynchronizationEndpointResponse) GetBody() *SwitchSynchronizationEndpointResponseBody {
	return s.Body
}

func (s *SwitchSynchronizationEndpointResponse) SetHeaders(v map[string]*string) *SwitchSynchronizationEndpointResponse {
	s.Headers = v
	return s
}

func (s *SwitchSynchronizationEndpointResponse) SetStatusCode(v int32) *SwitchSynchronizationEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponse) SetBody(v *SwitchSynchronizationEndpointResponseBody) *SwitchSynchronizationEndpointResponse {
	s.Body = v
	return s
}

func (s *SwitchSynchronizationEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
