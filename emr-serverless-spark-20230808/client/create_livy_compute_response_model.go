// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivyComputeResponseBody) *CreateLivyComputeResponse
	GetBody() *CreateLivyComputeResponseBody
}

type CreateLivyComputeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivyComputeResponse) GetBody() *CreateLivyComputeResponseBody {
	return s.Body
}

func (s *CreateLivyComputeResponse) SetHeaders(v map[string]*string) *CreateLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *CreateLivyComputeResponse) SetStatusCode(v int32) *CreateLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivyComputeResponse) SetBody(v *CreateLivyComputeResponseBody) *CreateLivyComputeResponse {
	s.Body = v
	return s
}

func (s *CreateLivyComputeResponse) Validate() error {
	return dara.Validate(s)
}
