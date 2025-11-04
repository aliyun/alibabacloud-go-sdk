// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachVServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachVServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DetachVServerGroupsResponseBody) *DetachVServerGroupsResponse
	GetBody() *DetachVServerGroupsResponseBody
}

type DetachVServerGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachVServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DetachVServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachVServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachVServerGroupsResponse) GetBody() *DetachVServerGroupsResponseBody {
	return s.Body
}

func (s *DetachVServerGroupsResponse) SetHeaders(v map[string]*string) *DetachVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DetachVServerGroupsResponse) SetStatusCode(v int32) *DetachVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVServerGroupsResponse) SetBody(v *DetachVServerGroupsResponseBody) *DetachVServerGroupsResponse {
	s.Body = v
	return s
}

func (s *DetachVServerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
