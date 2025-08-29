// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrafficControlTaskExpressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *CheckTrafficControlTaskExpressionRequest
	GetExpression() *string
	SetInstanceId(v string) *CheckTrafficControlTaskExpressionRequest
	GetInstanceId() *string
	SetTableMetaId(v string) *CheckTrafficControlTaskExpressionRequest
	GetTableMetaId() *string
}

type CheckTrafficControlTaskExpressionRequest struct {
	// This parameter is required.
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
}

func (s CheckTrafficControlTaskExpressionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTrafficControlTaskExpressionRequest) GoString() string {
	return s.String()
}

func (s *CheckTrafficControlTaskExpressionRequest) GetExpression() *string {
	return s.Expression
}

func (s *CheckTrafficControlTaskExpressionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckTrafficControlTaskExpressionRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *CheckTrafficControlTaskExpressionRequest) SetExpression(v string) *CheckTrafficControlTaskExpressionRequest {
	s.Expression = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionRequest) SetInstanceId(v string) *CheckTrafficControlTaskExpressionRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionRequest) SetTableMetaId(v string) *CheckTrafficControlTaskExpressionRequest {
	s.TableMetaId = &v
	return s
}

func (s *CheckTrafficControlTaskExpressionRequest) Validate() error {
	return dara.Validate(s)
}
