// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestProcessExpressionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestProcessExpressionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestProcessExpressionResponse
	GetStatusCode() *int32
	SetBody(v *TestProcessExpressionResponseBody) *TestProcessExpressionResponse
	GetBody() *TestProcessExpressionResponseBody
}

type TestProcessExpressionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestProcessExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestProcessExpressionResponse) String() string {
	return dara.Prettify(s)
}

func (s TestProcessExpressionResponse) GoString() string {
	return s.String()
}

func (s *TestProcessExpressionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestProcessExpressionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestProcessExpressionResponse) GetBody() *TestProcessExpressionResponseBody {
	return s.Body
}

func (s *TestProcessExpressionResponse) SetHeaders(v map[string]*string) *TestProcessExpressionResponse {
	s.Headers = v
	return s
}

func (s *TestProcessExpressionResponse) SetStatusCode(v int32) *TestProcessExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *TestProcessExpressionResponse) SetBody(v *TestProcessExpressionResponseBody) *TestProcessExpressionResponse {
	s.Body = v
	return s
}

func (s *TestProcessExpressionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
