// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveFlowLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveFlowLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveFlowLogResponse
	GetStatusCode() *int32
	SetBody(v *ActiveFlowLogResponseBody) *ActiveFlowLogResponse
	GetBody() *ActiveFlowLogResponseBody
}

type ActiveFlowLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveFlowLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveFlowLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveFlowLogResponse) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveFlowLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveFlowLogResponse) GetBody() *ActiveFlowLogResponseBody {
	return s.Body
}

func (s *ActiveFlowLogResponse) SetHeaders(v map[string]*string) *ActiveFlowLogResponse {
	s.Headers = v
	return s
}

func (s *ActiveFlowLogResponse) SetStatusCode(v int32) *ActiveFlowLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveFlowLogResponse) SetBody(v *ActiveFlowLogResponseBody) *ActiveFlowLogResponse {
	s.Body = v
	return s
}

func (s *ActiveFlowLogResponse) Validate() error {
	return dara.Validate(s)
}
