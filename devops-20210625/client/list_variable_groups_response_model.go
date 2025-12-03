// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVariableGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVariableGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListVariableGroupsResponseBody) *ListVariableGroupsResponse
	GetBody() *ListVariableGroupsResponseBody
}

type ListVariableGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVariableGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVariableGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVariableGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVariableGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVariableGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVariableGroupsResponse) GetBody() *ListVariableGroupsResponseBody {
	return s.Body
}

func (s *ListVariableGroupsResponse) SetHeaders(v map[string]*string) *ListVariableGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVariableGroupsResponse) SetStatusCode(v int32) *ListVariableGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariableGroupsResponse) SetBody(v *ListVariableGroupsResponseBody) *ListVariableGroupsResponse {
	s.Body = v
	return s
}

func (s *ListVariableGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
