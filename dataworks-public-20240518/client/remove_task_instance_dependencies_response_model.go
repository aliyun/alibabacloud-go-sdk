// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTaskInstanceDependenciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTaskInstanceDependenciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTaskInstanceDependenciesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTaskInstanceDependenciesResponseBody) *RemoveTaskInstanceDependenciesResponse
	GetBody() *RemoveTaskInstanceDependenciesResponseBody
}

type RemoveTaskInstanceDependenciesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTaskInstanceDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTaskInstanceDependenciesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTaskInstanceDependenciesResponse) GoString() string {
	return s.String()
}

func (s *RemoveTaskInstanceDependenciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTaskInstanceDependenciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTaskInstanceDependenciesResponse) GetBody() *RemoveTaskInstanceDependenciesResponseBody {
	return s.Body
}

func (s *RemoveTaskInstanceDependenciesResponse) SetHeaders(v map[string]*string) *RemoveTaskInstanceDependenciesResponse {
	s.Headers = v
	return s
}

func (s *RemoveTaskInstanceDependenciesResponse) SetStatusCode(v int32) *RemoveTaskInstanceDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesResponse) SetBody(v *RemoveTaskInstanceDependenciesResponseBody) *RemoveTaskInstanceDependenciesResponse {
	s.Body = v
	return s
}

func (s *RemoveTaskInstanceDependenciesResponse) Validate() error {
	return dara.Validate(s)
}
