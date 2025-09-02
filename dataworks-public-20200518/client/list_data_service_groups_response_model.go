// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceGroupsResponseBody) *ListDataServiceGroupsResponse
	GetBody() *ListDataServiceGroupsResponseBody
}

type ListDataServiceGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceGroupsResponse) GetBody() *ListDataServiceGroupsResponseBody {
	return s.Body
}

func (s *ListDataServiceGroupsResponse) SetHeaders(v map[string]*string) *ListDataServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceGroupsResponse) SetStatusCode(v int32) *ListDataServiceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceGroupsResponse) SetBody(v *ListDataServiceGroupsResponseBody) *ListDataServiceGroupsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
