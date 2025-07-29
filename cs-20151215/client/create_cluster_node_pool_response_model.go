// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateClusterNodePoolResponseBody) *CreateClusterNodePoolResponse
	GetBody() *CreateClusterNodePoolResponseBody
}

type CreateClusterNodePoolResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClusterNodePoolResponse) GetBody() *CreateClusterNodePoolResponseBody {
	return s.Body
}

func (s *CreateClusterNodePoolResponse) SetHeaders(v map[string]*string) *CreateClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterNodePoolResponse) SetStatusCode(v int32) *CreateClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterNodePoolResponse) SetBody(v *CreateClusterNodePoolResponseBody) *CreateClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *CreateClusterNodePoolResponse) Validate() error {
	return dara.Validate(s)
}
