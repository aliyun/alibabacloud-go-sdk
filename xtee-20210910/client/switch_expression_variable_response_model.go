// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchExpressionVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchExpressionVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchExpressionVariableResponse
	GetStatusCode() *int32
	SetBody(v *SwitchExpressionVariableResponseBody) *SwitchExpressionVariableResponse
	GetBody() *SwitchExpressionVariableResponseBody
}

type SwitchExpressionVariableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchExpressionVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchExpressionVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchExpressionVariableResponse) GoString() string {
	return s.String()
}

func (s *SwitchExpressionVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchExpressionVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchExpressionVariableResponse) GetBody() *SwitchExpressionVariableResponseBody {
	return s.Body
}

func (s *SwitchExpressionVariableResponse) SetHeaders(v map[string]*string) *SwitchExpressionVariableResponse {
	s.Headers = v
	return s
}

func (s *SwitchExpressionVariableResponse) SetStatusCode(v int32) *SwitchExpressionVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchExpressionVariableResponse) SetBody(v *SwitchExpressionVariableResponseBody) *SwitchExpressionVariableResponse {
	s.Body = v
	return s
}

func (s *SwitchExpressionVariableResponse) Validate() error {
	return dara.Validate(s)
}
