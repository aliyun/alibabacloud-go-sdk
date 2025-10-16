// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTimerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTimerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTimerGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTimerGroupResponseBody) *ModifyTimerGroupResponse
	GetBody() *ModifyTimerGroupResponseBody
}

type ModifyTimerGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTimerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTimerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTimerGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyTimerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTimerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTimerGroupResponse) GetBody() *ModifyTimerGroupResponseBody {
	return s.Body
}

func (s *ModifyTimerGroupResponse) SetHeaders(v map[string]*string) *ModifyTimerGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyTimerGroupResponse) SetStatusCode(v int32) *ModifyTimerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTimerGroupResponse) SetBody(v *ModifyTimerGroupResponseBody) *ModifyTimerGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyTimerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
