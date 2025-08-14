// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpressionVariableLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckExpressionVariableLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckExpressionVariableLimitResponse
	GetStatusCode() *int32
	SetBody(v *CheckExpressionVariableLimitResponseBody) *CheckExpressionVariableLimitResponse
	GetBody() *CheckExpressionVariableLimitResponseBody
}

type CheckExpressionVariableLimitResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckExpressionVariableLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckExpressionVariableLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckExpressionVariableLimitResponse) GoString() string {
	return s.String()
}

func (s *CheckExpressionVariableLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckExpressionVariableLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckExpressionVariableLimitResponse) GetBody() *CheckExpressionVariableLimitResponseBody {
	return s.Body
}

func (s *CheckExpressionVariableLimitResponse) SetHeaders(v map[string]*string) *CheckExpressionVariableLimitResponse {
	s.Headers = v
	return s
}

func (s *CheckExpressionVariableLimitResponse) SetStatusCode(v int32) *CheckExpressionVariableLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckExpressionVariableLimitResponse) SetBody(v *CheckExpressionVariableLimitResponseBody) *CheckExpressionVariableLimitResponse {
	s.Body = v
	return s
}

func (s *CheckExpressionVariableLimitResponse) Validate() error {
	return dara.Validate(s)
}
