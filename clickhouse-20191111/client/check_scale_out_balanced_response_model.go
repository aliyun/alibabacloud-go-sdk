// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckScaleOutBalancedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckScaleOutBalancedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckScaleOutBalancedResponse
	GetStatusCode() *int32
	SetBody(v *CheckScaleOutBalancedResponseBody) *CheckScaleOutBalancedResponse
	GetBody() *CheckScaleOutBalancedResponseBody
}

type CheckScaleOutBalancedResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckScaleOutBalancedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckScaleOutBalancedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckScaleOutBalancedResponse) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckScaleOutBalancedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckScaleOutBalancedResponse) GetBody() *CheckScaleOutBalancedResponseBody {
	return s.Body
}

func (s *CheckScaleOutBalancedResponse) SetHeaders(v map[string]*string) *CheckScaleOutBalancedResponse {
	s.Headers = v
	return s
}

func (s *CheckScaleOutBalancedResponse) SetStatusCode(v int32) *CheckScaleOutBalancedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckScaleOutBalancedResponse) SetBody(v *CheckScaleOutBalancedResponseBody) *CheckScaleOutBalancedResponse {
	s.Body = v
	return s
}

func (s *CheckScaleOutBalancedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
