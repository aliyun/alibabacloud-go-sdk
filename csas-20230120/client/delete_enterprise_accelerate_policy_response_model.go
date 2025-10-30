// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseAcceleratePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnterpriseAcceleratePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnterpriseAcceleratePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnterpriseAcceleratePolicyResponseBody) *DeleteEnterpriseAcceleratePolicyResponse
	GetBody() *DeleteEnterpriseAcceleratePolicyResponseBody
}

type DeleteEnterpriseAcceleratePolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnterpriseAcceleratePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnterpriseAcceleratePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseAcceleratePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) GetBody() *DeleteEnterpriseAcceleratePolicyResponseBody {
	return s.Body
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseAcceleratePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) SetStatusCode(v int32) *DeleteEnterpriseAcceleratePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) SetBody(v *DeleteEnterpriseAcceleratePolicyResponseBody) *DeleteEnterpriseAcceleratePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteEnterpriseAcceleratePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
