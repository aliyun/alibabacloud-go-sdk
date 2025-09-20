// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStoryResponseBody) *UpdateStoryResponse
	GetBody() *UpdateStoryResponseBody
}

type UpdateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStoryResponse) GetBody() *UpdateStoryResponseBody {
	return s.Body
}

func (s *UpdateStoryResponse) SetHeaders(v map[string]*string) *UpdateStoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoryResponse) SetStatusCode(v int32) *UpdateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoryResponse) SetBody(v *UpdateStoryResponseBody) *UpdateStoryResponse {
	s.Body = v
	return s
}

func (s *UpdateStoryResponse) Validate() error {
	return dara.Validate(s)
}
