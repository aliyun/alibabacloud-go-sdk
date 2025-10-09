// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponent(v *GetComponentResponseBodyComponent) *GetComponentResponseBody
	GetComponent() *GetComponentResponseBodyComponent
	SetRequestId(v string) *GetComponentResponseBody
	GetRequestId() *string
}

type GetComponentResponseBody struct {
	Component *GetComponentResponseBodyComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComponentResponseBody) GoString() string {
	return s.String()
}

func (s *GetComponentResponseBody) GetComponent() *GetComponentResponseBodyComponent {
	return s.Component
}

func (s *GetComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComponentResponseBody) SetComponent(v *GetComponentResponseBodyComponent) *GetComponentResponseBody {
	s.Component = v
	return s
}

func (s *GetComponentResponseBody) SetRequestId(v string) *GetComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComponentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetComponentResponseBodyComponent struct {
	// example:
	//
	// 43cd873b-235c-44f8-be07-e4d4cf7e73b0
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2017-04-27T05:37:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// None
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2024-01-26T07:44:21Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// dim_whse_epet_warehouse_jz_storage_stock_lot_relation_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 207316543660665792
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 64623
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// {
	//
	//     "kind": "Component",
	//
	//     "name": "com1",
	//
	//     "spec": {
	//
	//         "components": [
	//
	//             {
	//
	//                 "name": "test11",
	//
	//                 "id": "8196828925037*****",
	//
	//                 "owner": "054664",
	//
	//                 "description": "",
	//
	//                 "script": {
	//
	//                     "language": "odps-sql",
	//
	//                     "path": "test11",
	//
	//                     "content": "select \\"@@{bizdate}\\", \\"@@{my_input_table}\\"",
	//
	//                     "runtime": {
	//
	//                         "command": "SQL_COMPONENT"
	//
	//                     }
	//
	//                 },
	//
	//                 "inputs": [
	//
	//                     {
	//
	//                         "name": "bizdate",
	//
	//                         "type": "string"
	//
	//                     },
	//
	//                     {
	//
	//                         "name": "my_input_table",
	//
	//                         "type": "string"
	//
	//                     }
	//
	//                 ],
	//
	//                 "outputs": [
	//
	//                     {
	//
	//                         "name": "my_output_table1",
	//
	//                         "type": "string"
	//
	//                     }
	//
	//                 ]
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

func (s GetComponentResponseBodyComponent) String() string {
	return dara.Prettify(s)
}

func (s GetComponentResponseBodyComponent) GoString() string {
	return s.String()
}

func (s *GetComponentResponseBodyComponent) GetComponentId() *string {
	return s.ComponentId
}

func (s *GetComponentResponseBodyComponent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComponentResponseBodyComponent) GetDescription() *string {
	return s.Description
}

func (s *GetComponentResponseBodyComponent) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetComponentResponseBodyComponent) GetName() *string {
	return s.Name
}

func (s *GetComponentResponseBodyComponent) GetOwner() *string {
	return s.Owner
}

func (s *GetComponentResponseBodyComponent) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetComponentResponseBodyComponent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComponentResponseBodyComponent) GetSpec() *string {
	return s.Spec
}

func (s *GetComponentResponseBodyComponent) SetComponentId(v string) *GetComponentResponseBodyComponent {
	s.ComponentId = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetCreateTime(v string) *GetComponentResponseBodyComponent {
	s.CreateTime = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetDescription(v string) *GetComponentResponseBodyComponent {
	s.Description = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetModifyTime(v string) *GetComponentResponseBodyComponent {
	s.ModifyTime = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetName(v string) *GetComponentResponseBodyComponent {
	s.Name = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetOwner(v string) *GetComponentResponseBodyComponent {
	s.Owner = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetProjectId(v int64) *GetComponentResponseBodyComponent {
	s.ProjectId = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetRegionId(v string) *GetComponentResponseBodyComponent {
	s.RegionId = &v
	return s
}

func (s *GetComponentResponseBodyComponent) SetSpec(v string) *GetComponentResponseBodyComponent {
	s.Spec = &v
	return s
}

func (s *GetComponentResponseBodyComponent) Validate() error {
	return dara.Validate(s)
}
