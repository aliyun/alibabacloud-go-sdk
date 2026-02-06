// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobGroupsAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobGroupsAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ListJobGroupsAsyncResponseBody) *ListJobGroupsAsyncResponse
	GetBody() *ListJobGroupsAsyncResponseBody
}

type ListJobGroupsAsyncResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobGroupsAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobGroupsAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsAsyncResponse) GoString() string {
	return s.String()
}

func (s *ListJobGroupsAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobGroupsAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobGroupsAsyncResponse) GetBody() *ListJobGroupsAsyncResponseBody {
	return s.Body
}

func (s *ListJobGroupsAsyncResponse) SetHeaders(v map[string]*string) *ListJobGroupsAsyncResponse {
	s.Headers = v
	return s
}

func (s *ListJobGroupsAsyncResponse) SetStatusCode(v int32) *ListJobGroupsAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobGroupsAsyncResponse) SetBody(v *ListJobGroupsAsyncResponseBody) *ListJobGroupsAsyncResponse {
	s.Body = v
	return s
}

func (s *ListJobGroupsAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
