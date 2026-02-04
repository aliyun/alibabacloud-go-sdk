// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactiveFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactiveFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *DeactiveFlowLogResponseBody) *DeactiveFlowLogResponse
	GetBody() *DeactiveFlowLogResponseBody
}

type DeactiveFlowLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactiveFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactiveFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactiveFlowLogResponse) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactiveFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactiveFlowLogResponse) GetBody() *DeactiveFlowLogResponseBody {
	return s.Body
}

func (s *DeactiveFlowLogResponse) SetHeaders(v map[string]*string) *DeactiveFlowLogResponse {
	s.Headers = v
	return s
}

func (s *DeactiveFlowLogResponse) SetStatusCode(v int32) *DeactiveFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactiveFlowLogResponse) SetBody(v *DeactiveFlowLogResponseBody) *DeactiveFlowLogResponse {
	s.Body = v
	return s
}

func (s *DeactiveFlowLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
