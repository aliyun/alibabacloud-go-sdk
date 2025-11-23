// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowConstantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagConstants(v []*UpdateTaskFlowConstantsRequestDagConstants) *UpdateTaskFlowConstantsRequest
	GetDagConstants() []*UpdateTaskFlowConstantsRequestDagConstants
	SetDagId(v int64) *UpdateTaskFlowConstantsRequest
	GetDagId() *int64
	SetTid(v int64) *UpdateTaskFlowConstantsRequest
	GetTid() *int64
}

type UpdateTaskFlowConstantsRequest struct {
	// The constants for the task flow.
	DagConstants []*UpdateTaskFlowConstantsRequestDagConstants `json:"DagConstants,omitempty" xml:"DagConstants,omitempty" type:"Repeated"`
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowConstantsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowConstantsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowConstantsRequest) GetDagConstants() []*UpdateTaskFlowConstantsRequestDagConstants {
	return s.DagConstants
}

func (s *UpdateTaskFlowConstantsRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowConstantsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowConstantsRequest) SetDagConstants(v []*UpdateTaskFlowConstantsRequestDagConstants) *UpdateTaskFlowConstantsRequest {
	s.DagConstants = v
	return s
}

func (s *UpdateTaskFlowConstantsRequest) SetDagId(v int64) *UpdateTaskFlowConstantsRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowConstantsRequest) SetTid(v int64) *UpdateTaskFlowConstantsRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowConstantsRequest) Validate() error {
	if s.DagConstants != nil {
		for _, item := range s.DagConstants {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTaskFlowConstantsRequestDagConstants struct {
	// The key name of a constant for the task flow.
	//
	// example:
	//
	// poc_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The key value of a constant for the task flow.
	//
	// example:
	//
	// poc_test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTaskFlowConstantsRequestDagConstants) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowConstantsRequestDagConstants) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowConstantsRequestDagConstants) GetKey() *string {
	return s.Key
}

func (s *UpdateTaskFlowConstantsRequestDagConstants) GetValue() *string {
	return s.Value
}

func (s *UpdateTaskFlowConstantsRequestDagConstants) SetKey(v string) *UpdateTaskFlowConstantsRequestDagConstants {
	s.Key = &v
	return s
}

func (s *UpdateTaskFlowConstantsRequestDagConstants) SetValue(v string) *UpdateTaskFlowConstantsRequestDagConstants {
	s.Value = &v
	return s
}

func (s *UpdateTaskFlowConstantsRequestDagConstants) Validate() error {
	return dara.Validate(s)
}
