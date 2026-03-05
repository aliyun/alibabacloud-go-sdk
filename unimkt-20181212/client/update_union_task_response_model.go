// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUnionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUnionTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUnionTaskResponseBody) *UpdateUnionTaskResponse
	GetBody() *UpdateUnionTaskResponseBody
}

type UpdateUnionTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUnionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUnionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnionTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateUnionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUnionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUnionTaskResponse) GetBody() *UpdateUnionTaskResponseBody {
	return s.Body
}

func (s *UpdateUnionTaskResponse) SetHeaders(v map[string]*string) *UpdateUnionTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateUnionTaskResponse) SetStatusCode(v int32) *UpdateUnionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUnionTaskResponse) SetBody(v *UpdateUnionTaskResponseBody) *UpdateUnionTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateUnionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
