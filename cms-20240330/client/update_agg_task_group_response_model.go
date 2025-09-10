// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggTaskGroupResponseBody) *UpdateAggTaskGroupResponse
	GetBody() *UpdateAggTaskGroupResponseBody
}

type UpdateAggTaskGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggTaskGroupResponse) GetBody() *UpdateAggTaskGroupResponseBody {
	return s.Body
}

func (s *UpdateAggTaskGroupResponse) SetHeaders(v map[string]*string) *UpdateAggTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggTaskGroupResponse) SetStatusCode(v int32) *UpdateAggTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggTaskGroupResponse) SetBody(v *UpdateAggTaskGroupResponseBody) *UpdateAggTaskGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateAggTaskGroupResponse) Validate() error {
	return dara.Validate(s)
}
