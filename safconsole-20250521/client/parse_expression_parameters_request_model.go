// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseExpressionParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *ParseExpressionParametersRequest
	GetExpression() *string
}

type ParseExpressionParametersRequest struct {
	// example:
	//
	// score
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
}

func (s ParseExpressionParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ParseExpressionParametersRequest) GoString() string {
	return s.String()
}

func (s *ParseExpressionParametersRequest) GetExpression() *string {
	return s.Expression
}

func (s *ParseExpressionParametersRequest) SetExpression(v string) *ParseExpressionParametersRequest {
	s.Expression = &v
	return s
}

func (s *ParseExpressionParametersRequest) Validate() error {
	return dara.Validate(s)
}
