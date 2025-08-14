// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFavoriteVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateFavoriteVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateFavoriteVariableResponse
	GetStatusCode() *int32
	SetBody(v *OperateFavoriteVariableResponseBody) *OperateFavoriteVariableResponse
	GetBody() *OperateFavoriteVariableResponseBody
}

type OperateFavoriteVariableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateFavoriteVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateFavoriteVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateFavoriteVariableResponse) GoString() string {
	return s.String()
}

func (s *OperateFavoriteVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateFavoriteVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateFavoriteVariableResponse) GetBody() *OperateFavoriteVariableResponseBody {
	return s.Body
}

func (s *OperateFavoriteVariableResponse) SetHeaders(v map[string]*string) *OperateFavoriteVariableResponse {
	s.Headers = v
	return s
}

func (s *OperateFavoriteVariableResponse) SetStatusCode(v int32) *OperateFavoriteVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateFavoriteVariableResponse) SetBody(v *OperateFavoriteVariableResponseBody) *OperateFavoriteVariableResponse {
	s.Body = v
	return s
}

func (s *OperateFavoriteVariableResponse) Validate() error {
	return dara.Validate(s)
}
