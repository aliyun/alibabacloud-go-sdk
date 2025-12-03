// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemAllFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkItemAllFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkItemAllFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkItemAllFieldsResponseBody) *ListWorkItemAllFieldsResponse
	GetBody() *ListWorkItemAllFieldsResponseBody
}

type ListWorkItemAllFieldsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkItemAllFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkItemAllFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemAllFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkItemAllFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkItemAllFieldsResponse) GetBody() *ListWorkItemAllFieldsResponseBody {
	return s.Body
}

func (s *ListWorkItemAllFieldsResponse) SetHeaders(v map[string]*string) *ListWorkItemAllFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemAllFieldsResponse) SetStatusCode(v int32) *ListWorkItemAllFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemAllFieldsResponse) SetBody(v *ListWorkItemAllFieldsResponseBody) *ListWorkItemAllFieldsResponse {
	s.Body = v
	return s
}

func (s *ListWorkItemAllFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
