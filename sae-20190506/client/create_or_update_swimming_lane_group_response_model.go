// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateSwimmingLaneGroupResponseBody) *CreateOrUpdateSwimmingLaneGroupResponse
	GetBody() *CreateOrUpdateSwimmingLaneGroupResponseBody
}

type CreateOrUpdateSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) GetBody() *CreateOrUpdateSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *CreateOrUpdateSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) SetStatusCode(v int32) *CreateOrUpdateSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) SetBody(v *CreateOrUpdateSwimmingLaneGroupResponseBody) *CreateOrUpdateSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponse) Validate() error {
	return dara.Validate(s)
}
