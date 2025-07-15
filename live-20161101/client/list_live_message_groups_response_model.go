// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveMessageGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveMessageGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveMessageGroupsResponseBody) *ListLiveMessageGroupsResponse
	GetBody() *ListLiveMessageGroupsResponseBody
}

type ListLiveMessageGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveMessageGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveMessageGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveMessageGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveMessageGroupsResponse) GetBody() *ListLiveMessageGroupsResponseBody {
	return s.Body
}

func (s *ListLiveMessageGroupsResponse) SetHeaders(v map[string]*string) *ListLiveMessageGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveMessageGroupsResponse) SetStatusCode(v int32) *ListLiveMessageGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveMessageGroupsResponse) SetBody(v *ListLiveMessageGroupsResponseBody) *ListLiveMessageGroupsResponse {
	s.Body = v
	return s
}

func (s *ListLiveMessageGroupsResponse) Validate() error {
	return dara.Validate(s)
}
