// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleClusterResponse
	GetStatusCode() *int32
	SetBody(v *ScaleClusterResponseBody) *ScaleClusterResponse
	GetBody() *ScaleClusterResponseBody
}

type ScaleClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleClusterResponse) GetBody() *ScaleClusterResponseBody {
	return s.Body
}

func (s *ScaleClusterResponse) SetHeaders(v map[string]*string) *ScaleClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterResponse) SetStatusCode(v int32) *ScaleClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleClusterResponse) SetBody(v *ScaleClusterResponseBody) *ScaleClusterResponse {
	s.Body = v
	return s
}

func (s *ScaleClusterResponse) Validate() error {
	return dara.Validate(s)
}
