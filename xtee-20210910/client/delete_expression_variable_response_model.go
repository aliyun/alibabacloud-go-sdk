// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressionVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressionVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressionVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressionVariableResponseBody) *DeleteExpressionVariableResponse
	GetBody() *DeleteExpressionVariableResponseBody
}

type DeleteExpressionVariableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressionVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressionVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressionVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressionVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressionVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressionVariableResponse) GetBody() *DeleteExpressionVariableResponseBody {
	return s.Body
}

func (s *DeleteExpressionVariableResponse) SetHeaders(v map[string]*string) *DeleteExpressionVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressionVariableResponse) SetStatusCode(v int32) *DeleteExpressionVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressionVariableResponse) SetBody(v *DeleteExpressionVariableResponseBody) *DeleteExpressionVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressionVariableResponse) Validate() error {
	return dara.Validate(s)
}
