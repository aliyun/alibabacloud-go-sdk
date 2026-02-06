// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyJobGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyJobGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyJobGroupResponseBody) *ModifyJobGroupResponse
	GetBody() *ModifyJobGroupResponseBody
}

type ModifyJobGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyJobGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyJobGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyJobGroupResponse) GetBody() *ModifyJobGroupResponseBody {
	return s.Body
}

func (s *ModifyJobGroupResponse) SetHeaders(v map[string]*string) *ModifyJobGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyJobGroupResponse) SetStatusCode(v int32) *ModifyJobGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyJobGroupResponse) SetBody(v *ModifyJobGroupResponseBody) *ModifyJobGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyJobGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
