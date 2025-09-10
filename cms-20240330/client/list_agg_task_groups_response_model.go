// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggTaskGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggTaskGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggTaskGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggTaskGroupsResponseBody) *ListAggTaskGroupsResponse
	GetBody() *ListAggTaskGroupsResponseBody
}

type ListAggTaskGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggTaskGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggTaskGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggTaskGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAggTaskGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggTaskGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggTaskGroupsResponse) GetBody() *ListAggTaskGroupsResponseBody {
	return s.Body
}

func (s *ListAggTaskGroupsResponse) SetHeaders(v map[string]*string) *ListAggTaskGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAggTaskGroupsResponse) SetStatusCode(v int32) *ListAggTaskGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggTaskGroupsResponse) SetBody(v *ListAggTaskGroupsResponseBody) *ListAggTaskGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAggTaskGroupsResponse) Validate() error {
	return dara.Validate(s)
}
