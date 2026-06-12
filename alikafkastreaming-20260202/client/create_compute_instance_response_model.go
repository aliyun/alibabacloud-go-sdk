// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeInstanceResponseBody) *CreateComputeInstanceResponse
	GetBody() *CreateComputeInstanceResponseBody
}

type CreateComputeInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeInstanceResponse) GetBody() *CreateComputeInstanceResponseBody {
	return s.Body
}

func (s *CreateComputeInstanceResponse) SetHeaders(v map[string]*string) *CreateComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeInstanceResponse) SetStatusCode(v int32) *CreateComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeInstanceResponse) SetBody(v *CreateComputeInstanceResponseBody) *CreateComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
