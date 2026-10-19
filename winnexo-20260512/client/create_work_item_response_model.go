// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkItemResponseBody) *CreateWorkItemResponse
	GetBody() *CreateWorkItemResponseBody
}

type CreateWorkItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkItemResponse) GetBody() *CreateWorkItemResponseBody {
	return s.Body
}

func (s *CreateWorkItemResponse) SetHeaders(v map[string]*string) *CreateWorkItemResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkItemResponse) SetStatusCode(v int32) *CreateWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkItemResponse) SetBody(v *CreateWorkItemResponseBody) *CreateWorkItemResponse {
	s.Body = v
	return s
}

func (s *CreateWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
