// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeClusterResponseBody) *UpdateComputeClusterResponse
	GetBody() *UpdateComputeClusterResponseBody
}

type UpdateComputeClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeClusterResponse) GetBody() *UpdateComputeClusterResponseBody {
	return s.Body
}

func (s *UpdateComputeClusterResponse) SetHeaders(v map[string]*string) *UpdateComputeClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeClusterResponse) SetStatusCode(v int32) *UpdateComputeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeClusterResponse) SetBody(v *UpdateComputeClusterResponseBody) *UpdateComputeClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
