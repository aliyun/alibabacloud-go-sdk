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
	// The expression to validate.
	//
	// This parameter is required.
	//
	// example:
	//
	// event=exposure
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec_123****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
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
