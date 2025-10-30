// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleOutClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleOutClusterResponse
	GetStatusCode() *int32
	SetBody(v *ScaleOutClusterResponseBody) *ScaleOutClusterResponse
	GetBody() *ScaleOutClusterResponseBody
}

type ScaleOutClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleOutClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleOutClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleOutClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleOutClusterResponse) GetBody() *ScaleOutClusterResponseBody {
	return s.Body
}

func (s *ScaleOutClusterResponse) SetHeaders(v map[string]*string) *ScaleOutClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutClusterResponse) SetStatusCode(v int32) *ScaleOutClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleOutClusterResponse) SetBody(v *ScaleOutClusterResponseBody) *ScaleOutClusterResponse {
	s.Body = v
	return s
}

func (s *ScaleOutClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
