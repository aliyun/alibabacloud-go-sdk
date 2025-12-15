// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyParameterGroupResponseBody) *ModifyParameterGroupResponse
	GetBody() *ModifyParameterGroupResponseBody
}

type ModifyParameterGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyParameterGroupResponse) GetBody() *ModifyParameterGroupResponseBody {
	return s.Body
}

func (s *ModifyParameterGroupResponse) SetHeaders(v map[string]*string) *ModifyParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyParameterGroupResponse) SetStatusCode(v int32) *ModifyParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParameterGroupResponse) SetBody(v *ModifyParameterGroupResponseBody) *ModifyParameterGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
