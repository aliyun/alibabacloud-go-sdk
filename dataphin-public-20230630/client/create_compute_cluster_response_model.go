// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeClusterResponseBody) *CreateComputeClusterResponse
	GetBody() *CreateComputeClusterResponseBody
}

type CreateComputeClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeClusterResponse) GetBody() *CreateComputeClusterResponseBody {
	return s.Body
}

func (s *CreateComputeClusterResponse) SetHeaders(v map[string]*string) *CreateComputeClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeClusterResponse) SetStatusCode(v int32) *CreateComputeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeClusterResponse) SetBody(v *CreateComputeClusterResponseBody) *CreateComputeClusterResponse {
	s.Body = v
	return s
}

func (s *CreateComputeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
