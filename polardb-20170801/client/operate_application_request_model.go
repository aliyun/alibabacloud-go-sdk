// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *OperateApplicationRequest
	GetApplicationId() *string
	SetOperation(v string) *OperateApplicationRequest
	GetOperation() *string
}

type OperateApplicationRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The operation type. Valid values:
	//
	// 	- **restart**: restart
	//
	// 	- **stop**: stop
	//
	// 	- **start**: start.
	//
	// example:
	//
	// restart
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s OperateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationRequest) GoString() string {
	return s.String()
}

func (s *OperateApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *OperateApplicationRequest) GetOperation() *string {
	return s.Operation
}

func (s *OperateApplicationRequest) SetApplicationId(v string) *OperateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *OperateApplicationRequest) SetOperation(v string) *OperateApplicationRequest {
	s.Operation = &v
	return s
}

func (s *OperateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
