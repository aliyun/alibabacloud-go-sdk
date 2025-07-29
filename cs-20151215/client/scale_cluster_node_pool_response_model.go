// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *ScaleClusterNodePoolResponseBody) *ScaleClusterNodePoolResponse
	GetBody() *ScaleClusterNodePoolResponseBody
}

type ScaleClusterNodePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleClusterNodePoolResponse) GetBody() *ScaleClusterNodePoolResponseBody {
	return s.Body
}

func (s *ScaleClusterNodePoolResponse) SetHeaders(v map[string]*string) *ScaleClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterNodePoolResponse) SetStatusCode(v int32) *ScaleClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleClusterNodePoolResponse) SetBody(v *ScaleClusterNodePoolResponseBody) *ScaleClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *ScaleClusterNodePoolResponse) Validate() error {
	return dara.Validate(s)
}
