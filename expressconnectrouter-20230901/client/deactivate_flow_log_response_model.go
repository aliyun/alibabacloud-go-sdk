// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateFlowLogResponseBody) *DeactivateFlowLogResponse
	GetBody() *DeactivateFlowLogResponseBody
}

type DeactivateFlowLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *DeactivateFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateFlowLogResponse) GetBody() *DeactivateFlowLogResponseBody {
	return s.Body
}

func (s *DeactivateFlowLogResponse) SetHeaders(v map[string]*string) *DeactivateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *DeactivateFlowLogResponse) SetStatusCode(v int32) *DeactivateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateFlowLogResponse) SetBody(v *DeactivateFlowLogResponseBody) *DeactivateFlowLogResponse {
	s.Body = v
	return s
}

func (s *DeactivateFlowLogResponse) Validate() error {
	return dara.Validate(s)
}
