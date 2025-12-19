// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v string) *GetOperationParamsRequest
	GetOperation() *string
	SetServiceType(v string) *GetOperationParamsRequest
	GetServiceType() *string
}

type GetOperationParamsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rename
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetOperationParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationParamsRequest) GoString() string {
	return s.String()
}

func (s *GetOperationParamsRequest) GetOperation() *string {
	return s.Operation
}

func (s *GetOperationParamsRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetOperationParamsRequest) SetOperation(v string) *GetOperationParamsRequest {
	s.Operation = &v
	return s
}

func (s *GetOperationParamsRequest) SetServiceType(v string) *GetOperationParamsRequest {
	s.ServiceType = &v
	return s
}

func (s *GetOperationParamsRequest) Validate() error {
	return dara.Validate(s)
}
