// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressionVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressionVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressionVariableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressionVariableResponseBody) *ModifyExpressionVariableResponse
	GetBody() *ModifyExpressionVariableResponseBody
}

type ModifyExpressionVariableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressionVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressionVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressionVariableResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressionVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressionVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressionVariableResponse) GetBody() *ModifyExpressionVariableResponseBody {
	return s.Body
}

func (s *ModifyExpressionVariableResponse) SetHeaders(v map[string]*string) *ModifyExpressionVariableResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressionVariableResponse) SetStatusCode(v int32) *ModifyExpressionVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressionVariableResponse) SetBody(v *ModifyExpressionVariableResponseBody) *ModifyExpressionVariableResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressionVariableResponse) Validate() error {
	return dara.Validate(s)
}
