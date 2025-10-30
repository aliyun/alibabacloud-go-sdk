// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeSourceResponseBody) *UpdateComputeSourceResponse
	GetBody() *UpdateComputeSourceResponseBody
}

type UpdateComputeSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeSourceResponse) GetBody() *UpdateComputeSourceResponseBody {
	return s.Body
}

func (s *UpdateComputeSourceResponse) SetHeaders(v map[string]*string) *UpdateComputeSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeSourceResponse) SetStatusCode(v int32) *UpdateComputeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeSourceResponse) SetBody(v *UpdateComputeSourceResponseBody) *UpdateComputeSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
