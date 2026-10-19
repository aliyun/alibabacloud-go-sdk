// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkItemResponseBody) *UpdateWorkItemResponse
	GetBody() *UpdateWorkItemResponseBody
}

type UpdateWorkItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkItemResponse) GetBody() *UpdateWorkItemResponseBody {
	return s.Body
}

func (s *UpdateWorkItemResponse) SetHeaders(v map[string]*string) *UpdateWorkItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkItemResponse) SetStatusCode(v int32) *UpdateWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkItemResponse) SetBody(v *UpdateWorkItemResponseBody) *UpdateWorkItemResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
