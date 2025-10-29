// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbanLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbanLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbanLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *UnbanLiveMessageGroupResponseBody) *UnbanLiveMessageGroupResponse
	GetBody() *UnbanLiveMessageGroupResponseBody
}

type UnbanLiveMessageGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbanLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbanLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbanLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbanLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbanLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbanLiveMessageGroupResponse) GetBody() *UnbanLiveMessageGroupResponseBody {
	return s.Body
}

func (s *UnbanLiveMessageGroupResponse) SetHeaders(v map[string]*string) *UnbanLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *UnbanLiveMessageGroupResponse) SetStatusCode(v int32) *UnbanLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbanLiveMessageGroupResponse) SetBody(v *UnbanLiveMessageGroupResponseBody) *UnbanLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *UnbanLiveMessageGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
