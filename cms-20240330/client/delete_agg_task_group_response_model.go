// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggTaskGroupResponseBody) *DeleteAggTaskGroupResponse
	GetBody() *DeleteAggTaskGroupResponseBody
}

type DeleteAggTaskGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggTaskGroupResponse) GetBody() *DeleteAggTaskGroupResponseBody {
	return s.Body
}

func (s *DeleteAggTaskGroupResponse) SetHeaders(v map[string]*string) *DeleteAggTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggTaskGroupResponse) SetStatusCode(v int32) *DeleteAggTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggTaskGroupResponse) SetBody(v *DeleteAggTaskGroupResponseBody) *DeleteAggTaskGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAggTaskGroupResponse) Validate() error {
	return dara.Validate(s)
}
