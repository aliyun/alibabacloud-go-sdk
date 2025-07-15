// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageGroupResponseBody) *ModifyLiveMessageGroupResponse
	GetBody() *ModifyLiveMessageGroupResponseBody
}

type ModifyLiveMessageGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageGroupResponse) GetBody() *ModifyLiveMessageGroupResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageGroupResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageGroupResponse) SetStatusCode(v int32) *ModifyLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageGroupResponse) SetBody(v *ModifyLiveMessageGroupResponseBody) *ModifyLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
