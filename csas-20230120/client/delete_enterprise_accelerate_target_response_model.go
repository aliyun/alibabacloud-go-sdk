// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAccelerateTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnterpriseAccelerateTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnterpriseAccelerateTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnterpriseAccelerateTargetResponseBody) *DeleteEnterpriseAccelerateTargetResponse
	GetBody() *DeleteEnterpriseAccelerateTargetResponseBody
}

type DeleteEnterpriseAccelerateTargetResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnterpriseAccelerateTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnterpriseAccelerateTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAccelerateTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAccelerateTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnterpriseAccelerateTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnterpriseAccelerateTargetResponse) GetBody() *DeleteEnterpriseAccelerateTargetResponseBody {
	return s.Body
}

func (s *DeleteEnterpriseAccelerateTargetResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseAccelerateTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetResponse) SetStatusCode(v int32) *DeleteEnterpriseAccelerateTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetResponse) SetBody(v *DeleteEnterpriseAccelerateTargetResponseBody) *DeleteEnterpriseAccelerateTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteEnterpriseAccelerateTargetResponse) Validate() error {
	return dara.Validate(s)
}
