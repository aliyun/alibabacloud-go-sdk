// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtaskItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubtaskItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubtaskItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubtaskItemsResponseBody) *ListSubtaskItemsResponse
	GetBody() *ListSubtaskItemsResponseBody
}

type ListSubtaskItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubtaskItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubtaskItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubtaskItemsResponse) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubtaskItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubtaskItemsResponse) GetBody() *ListSubtaskItemsResponseBody {
	return s.Body
}

func (s *ListSubtaskItemsResponse) SetHeaders(v map[string]*string) *ListSubtaskItemsResponse {
	s.Headers = v
	return s
}

func (s *ListSubtaskItemsResponse) SetStatusCode(v int32) *ListSubtaskItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubtaskItemsResponse) SetBody(v *ListSubtaskItemsResponseBody) *ListSubtaskItemsResponse {
	s.Body = v
	return s
}

func (s *ListSubtaskItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
