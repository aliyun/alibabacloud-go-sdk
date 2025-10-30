// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressionVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressionVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressionVariableResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressionVariableResponseBody) *CreateExpressionVariableResponse
	GetBody() *CreateExpressionVariableResponseBody
}

type CreateExpressionVariableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressionVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressionVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressionVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressionVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressionVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressionVariableResponse) GetBody() *CreateExpressionVariableResponseBody {
	return s.Body
}

func (s *CreateExpressionVariableResponse) SetHeaders(v map[string]*string) *CreateExpressionVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressionVariableResponse) SetStatusCode(v int32) *CreateExpressionVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressionVariableResponse) SetBody(v *CreateExpressionVariableResponseBody) *CreateExpressionVariableResponse {
	s.Body = v
	return s
}

func (s *CreateExpressionVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
