// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationActivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOperationActivityRequest
	GetInstanceId() *string
	SetOperationId(v string) *ListOperationActivityRequest
	GetOperationId() *string
}

type ListOperationActivityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// op-f49743caa809****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s ListOperationActivityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationActivityRequest) GoString() string {
	return s.String()
}

func (s *ListOperationActivityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationActivityRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *ListOperationActivityRequest) SetInstanceId(v string) *ListOperationActivityRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationActivityRequest) SetOperationId(v string) *ListOperationActivityRequest {
	s.OperationId = &v
	return s
}

func (s *ListOperationActivityRequest) Validate() error {
	return dara.Validate(s)
}
