// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeRightAngleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OptimizeRightAngleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OptimizeRightAngleResponse
	GetStatusCode() *int32
	SetBody(v *OptimizeRightAngleResponseBody) *OptimizeRightAngleResponse
	GetBody() *OptimizeRightAngleResponseBody
}

type OptimizeRightAngleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OptimizeRightAngleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OptimizeRightAngleResponse) String() string {
	return dara.Prettify(s)
}

func (s OptimizeRightAngleResponse) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OptimizeRightAngleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OptimizeRightAngleResponse) GetBody() *OptimizeRightAngleResponseBody {
	return s.Body
}

func (s *OptimizeRightAngleResponse) SetHeaders(v map[string]*string) *OptimizeRightAngleResponse {
	s.Headers = v
	return s
}

func (s *OptimizeRightAngleResponse) SetStatusCode(v int32) *OptimizeRightAngleResponse {
	s.StatusCode = &v
	return s
}

func (s *OptimizeRightAngleResponse) SetBody(v *OptimizeRightAngleResponseBody) *OptimizeRightAngleResponse {
	s.Body = v
	return s
}

func (s *OptimizeRightAngleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
