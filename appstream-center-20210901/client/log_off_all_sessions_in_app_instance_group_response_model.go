// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogOffAllSessionsInAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LogOffAllSessionsInAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LogOffAllSessionsInAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *LogOffAllSessionsInAppInstanceGroupResponseBody) *LogOffAllSessionsInAppInstanceGroupResponse
	GetBody() *LogOffAllSessionsInAppInstanceGroupResponseBody
}

type LogOffAllSessionsInAppInstanceGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogOffAllSessionsInAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) GetBody() *LogOffAllSessionsInAppInstanceGroupResponseBody {
	return s.Body
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetHeaders(v map[string]*string) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetStatusCode(v int32) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetBody(v *LogOffAllSessionsInAppInstanceGroupResponseBody) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
