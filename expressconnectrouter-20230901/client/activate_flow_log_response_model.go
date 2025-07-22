// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *ActivateFlowLogResponseBody) *ActivateFlowLogResponse
	GetBody() *ActivateFlowLogResponseBody
}

type ActivateFlowLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateFlowLogResponse) GoString() string {
	return s.String()
}

func (s *ActivateFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateFlowLogResponse) GetBody() *ActivateFlowLogResponseBody {
	return s.Body
}

func (s *ActivateFlowLogResponse) SetHeaders(v map[string]*string) *ActivateFlowLogResponse {
	s.Headers = v
	return s
}

func (s *ActivateFlowLogResponse) SetStatusCode(v int32) *ActivateFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateFlowLogResponse) SetBody(v *ActivateFlowLogResponseBody) *ActivateFlowLogResponse {
	s.Body = v
	return s
}

func (s *ActivateFlowLogResponse) Validate() error {
	return dara.Validate(s)
}
