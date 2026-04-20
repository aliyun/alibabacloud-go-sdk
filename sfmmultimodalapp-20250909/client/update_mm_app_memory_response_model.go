// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppMemoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppMemoryResponseBody) *UpdateMmAppMemoryResponse
	GetBody() *UpdateMmAppMemoryResponseBody
}

type UpdateMmAppMemoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppMemoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppMemoryResponse) GetBody() *UpdateMmAppMemoryResponseBody {
	return s.Body
}

func (s *UpdateMmAppMemoryResponse) SetHeaders(v map[string]*string) *UpdateMmAppMemoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppMemoryResponse) SetStatusCode(v int32) *UpdateMmAppMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppMemoryResponse) SetBody(v *UpdateMmAppMemoryResponseBody) *UpdateMmAppMemoryResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
