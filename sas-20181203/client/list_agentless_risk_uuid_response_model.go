// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRiskUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentlessRiskUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentlessRiskUuidResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentlessRiskUuidResponseBody) *ListAgentlessRiskUuidResponse
	GetBody() *ListAgentlessRiskUuidResponseBody
}

type ListAgentlessRiskUuidResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentlessRiskUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentlessRiskUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRiskUuidResponse) GoString() string {
	return s.String()
}

func (s *ListAgentlessRiskUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentlessRiskUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentlessRiskUuidResponse) GetBody() *ListAgentlessRiskUuidResponseBody {
	return s.Body
}

func (s *ListAgentlessRiskUuidResponse) SetHeaders(v map[string]*string) *ListAgentlessRiskUuidResponse {
	s.Headers = v
	return s
}

func (s *ListAgentlessRiskUuidResponse) SetStatusCode(v int32) *ListAgentlessRiskUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentlessRiskUuidResponse) SetBody(v *ListAgentlessRiskUuidResponseBody) *ListAgentlessRiskUuidResponse {
	s.Body = v
	return s
}

func (s *ListAgentlessRiskUuidResponse) Validate() error {
	return dara.Validate(s)
}
