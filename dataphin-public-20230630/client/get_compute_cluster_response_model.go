// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeClusterResponseBody) *GetComputeClusterResponse
	GetBody() *GetComputeClusterResponseBody
}

type GetComputeClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeClusterResponse) GoString() string {
	return s.String()
}

func (s *GetComputeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeClusterResponse) GetBody() *GetComputeClusterResponseBody {
	return s.Body
}

func (s *GetComputeClusterResponse) SetHeaders(v map[string]*string) *GetComputeClusterResponse {
	s.Headers = v
	return s
}

func (s *GetComputeClusterResponse) SetStatusCode(v int32) *GetComputeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeClusterResponse) SetBody(v *GetComputeClusterResponseBody) *GetComputeClusterResponse {
	s.Body = v
	return s
}

func (s *GetComputeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
