// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrafficControlTaskExpressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTrafficControlTaskExpressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTrafficControlTaskExpressionResponse
	GetStatusCode() *int32
	SetBody(v *CheckTrafficControlTaskExpressionResponseBody) *CheckTrafficControlTaskExpressionResponse
	GetBody() *CheckTrafficControlTaskExpressionResponseBody
}

type CheckTrafficControlTaskExpressionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTrafficControlTaskExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTrafficControlTaskExpressionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionResponse) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTrafficControlTaskExpressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTrafficControlTaskExpressionResponse) GetBody() *CheckTrafficControlTaskExpressionResponseBody {
	return s.Body
}

func (s *CheckTrafficControlTaskExpressionResponse) SetHeaders(v map[string]*string) *CheckTrafficControlTaskExpressionResponse {
	s.Headers = v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponse) SetStatusCode(v int32) *CheckTrafficControlTaskExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponse) SetBody(v *CheckTrafficControlTaskExpressionResponseBody) *CheckTrafficControlTaskExpressionResponse {
	s.Body = v
	return s
}

func (s *CheckTrafficControlTaskExpressionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
