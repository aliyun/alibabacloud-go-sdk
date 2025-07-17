// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DetachServerGroupsResponseBody) *DetachServerGroupsResponse
	GetBody() *DetachServerGroupsResponseBody
}

type DetachServerGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DetachServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachServerGroupsResponse) GetBody() *DetachServerGroupsResponseBody {
	return s.Body
}

func (s *DetachServerGroupsResponse) SetHeaders(v map[string]*string) *DetachServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DetachServerGroupsResponse) SetStatusCode(v int32) *DetachServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachServerGroupsResponse) SetBody(v *DetachServerGroupsResponseBody) *DetachServerGroupsResponse {
	s.Body = v
	return s
}

func (s *DetachServerGroupsResponse) Validate() error {
	return dara.Validate(s)
}
