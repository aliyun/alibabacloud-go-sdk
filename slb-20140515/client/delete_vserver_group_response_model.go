// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVServerGroupResponseBody) *DeleteVServerGroupResponse
	GetBody() *DeleteVServerGroupResponseBody
}

type DeleteVServerGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVServerGroupResponse) GetBody() *DeleteVServerGroupResponseBody {
	return s.Body
}

func (s *DeleteVServerGroupResponse) SetHeaders(v map[string]*string) *DeleteVServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVServerGroupResponse) SetStatusCode(v int32) *DeleteVServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVServerGroupResponse) SetBody(v *DeleteVServerGroupResponseBody) *DeleteVServerGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteVServerGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
