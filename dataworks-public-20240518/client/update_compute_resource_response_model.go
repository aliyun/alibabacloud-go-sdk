// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeResourceResponseBody) *UpdateComputeResourceResponse
	GetBody() *UpdateComputeResourceResponseBody
}

type UpdateComputeResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeResourceResponse) GetBody() *UpdateComputeResourceResponseBody {
	return s.Body
}

func (s *UpdateComputeResourceResponse) SetHeaders(v map[string]*string) *UpdateComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeResourceResponse) SetStatusCode(v int32) *UpdateComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeResourceResponse) SetBody(v *UpdateComputeResourceResponseBody) *UpdateComputeResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
