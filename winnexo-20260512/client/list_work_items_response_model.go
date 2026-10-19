// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkItemsResponseBody) *ListWorkItemsResponse
	GetBody() *ListWorkItemsResponseBody
}

type ListWorkItemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkItemsResponse) GetBody() *ListWorkItemsResponseBody {
	return s.Body
}

func (s *ListWorkItemsResponse) SetHeaders(v map[string]*string) *ListWorkItemsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkItemsResponse) SetStatusCode(v int32) *ListWorkItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkItemsResponse) SetBody(v *ListWorkItemsResponseBody) *ListWorkItemsResponse {
	s.Body = v
	return s
}

func (s *ListWorkItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
