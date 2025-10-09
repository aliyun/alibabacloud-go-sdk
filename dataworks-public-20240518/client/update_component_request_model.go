// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *UpdateComponentRequest
	GetComponentId() *string
	SetProjectId(v int64) *UpdateComponentRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateComponentRequest
	GetSpec() *string
}

type UpdateComponentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 568780096083*******
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec information for this UDF function. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "kind": "Component",
	//
	//     "version": "1.1.2",
	//
	//     "spec": {
	//
	//         "components": [
	//
	//             {
	//
	//                 "id": "568780096083*******",
	//
	//                 "script": {
	//
	//                     "content": "select \\"@@{para1}\\", \\"@@{para2}\\""
	//
	//                 }
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *UpdateComponentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateComponentRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateComponentRequest) SetComponentId(v string) *UpdateComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *UpdateComponentRequest) SetProjectId(v int64) *UpdateComponentRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateComponentRequest) SetSpec(v string) *UpdateComponentRequest {
	s.Spec = &v
	return s
}

func (s *UpdateComponentRequest) Validate() error {
	return dara.Validate(s)
}
