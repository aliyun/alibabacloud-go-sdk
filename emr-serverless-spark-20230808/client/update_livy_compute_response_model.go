// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivyComputeResponseBody) *UpdateLivyComputeResponse
	GetBody() *UpdateLivyComputeResponseBody
}

type UpdateLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivyComputeResponse) GetBody() *UpdateLivyComputeResponseBody {
	return s.Body
}

func (s *UpdateLivyComputeResponse) SetHeaders(v map[string]*string) *UpdateLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivyComputeResponse) SetStatusCode(v int32) *UpdateLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivyComputeResponse) SetBody(v *UpdateLivyComputeResponseBody) *UpdateLivyComputeResponse {
	s.Body = v
	return s
}

func (s *UpdateLivyComputeResponse) Validate() error {
	return dara.Validate(s)
}
