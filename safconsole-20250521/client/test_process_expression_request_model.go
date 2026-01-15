// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestProcessExpressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *TestProcessExpressionRequest
	GetExpression() *string
	SetParams(v string) *TestProcessExpressionRequest
	GetParams() *string
}

type TestProcessExpressionRequest struct {
	// example:
	//
	// score
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// {"score": "1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s TestProcessExpressionRequest) String() string {
	return dara.Prettify(s)
}

func (s TestProcessExpressionRequest) GoString() string {
	return s.String()
}

func (s *TestProcessExpressionRequest) GetExpression() *string {
	return s.Expression
}

func (s *TestProcessExpressionRequest) GetParams() *string {
	return s.Params
}

func (s *TestProcessExpressionRequest) SetExpression(v string) *TestProcessExpressionRequest {
	s.Expression = &v
	return s
}

func (s *TestProcessExpressionRequest) SetParams(v string) *TestProcessExpressionRequest {
	s.Params = &v
	return s
}

func (s *TestProcessExpressionRequest) Validate() error {
	return dara.Validate(s)
}
