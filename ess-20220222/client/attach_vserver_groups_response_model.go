// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachVServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachVServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *AttachVServerGroupsResponseBody) *AttachVServerGroupsResponse
	GetBody() *AttachVServerGroupsResponseBody
}

type AttachVServerGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachVServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachVServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachVServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachVServerGroupsResponse) GetBody() *AttachVServerGroupsResponseBody {
	return s.Body
}

func (s *AttachVServerGroupsResponse) SetHeaders(v map[string]*string) *AttachVServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *AttachVServerGroupsResponse) SetStatusCode(v int32) *AttachVServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVServerGroupsResponse) SetBody(v *AttachVServerGroupsResponseBody) *AttachVServerGroupsResponse {
	s.Body = v
	return s
}

func (s *AttachVServerGroupsResponse) Validate() error {
	return dara.Validate(s)
}
