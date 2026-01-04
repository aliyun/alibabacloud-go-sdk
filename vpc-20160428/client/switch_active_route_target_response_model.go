// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveRouteTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchActiveRouteTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchActiveRouteTargetResponse
	GetStatusCode() *int32
	SetBody(v *SwitchActiveRouteTargetResponseBody) *SwitchActiveRouteTargetResponse
	GetBody() *SwitchActiveRouteTargetResponseBody
}

type SwitchActiveRouteTargetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchActiveRouteTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchActiveRouteTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveRouteTargetResponse) GoString() string {
	return s.String()
}

func (s *SwitchActiveRouteTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchActiveRouteTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchActiveRouteTargetResponse) GetBody() *SwitchActiveRouteTargetResponseBody {
	return s.Body
}

func (s *SwitchActiveRouteTargetResponse) SetHeaders(v map[string]*string) *SwitchActiveRouteTargetResponse {
	s.Headers = v
	return s
}

func (s *SwitchActiveRouteTargetResponse) SetStatusCode(v int32) *SwitchActiveRouteTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchActiveRouteTargetResponse) SetBody(v *SwitchActiveRouteTargetResponseBody) *SwitchActiveRouteTargetResponse {
	s.Body = v
	return s
}

func (s *SwitchActiveRouteTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
