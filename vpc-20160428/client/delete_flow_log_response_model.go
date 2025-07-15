// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowLogResponseBody) *DeleteFlowLogResponse
	GetBody() *DeleteFlowLogResponseBody
}

type DeleteFlowLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowLogResponse) GetBody() *DeleteFlowLogResponseBody {
	return s.Body
}

func (s *DeleteFlowLogResponse) SetHeaders(v map[string]*string) *DeleteFlowLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowLogResponse) SetStatusCode(v int32) *DeleteFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowLogResponse) SetBody(v *DeleteFlowLogResponseBody) *DeleteFlowLogResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowLogResponse) Validate() error {
	return dara.Validate(s)
}
