// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggTaskGroupResponseBody) *CreateAggTaskGroupResponse
	GetBody() *CreateAggTaskGroupResponseBody
}

type CreateAggTaskGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAggTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggTaskGroupResponse) GetBody() *CreateAggTaskGroupResponseBody {
	return s.Body
}

func (s *CreateAggTaskGroupResponse) SetHeaders(v map[string]*string) *CreateAggTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAggTaskGroupResponse) SetStatusCode(v int32) *CreateAggTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggTaskGroupResponse) SetBody(v *CreateAggTaskGroupResponseBody) *CreateAggTaskGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAggTaskGroupResponse) Validate() error {
	return dara.Validate(s)
}
