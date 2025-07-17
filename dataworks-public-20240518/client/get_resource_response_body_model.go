// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceResponseBody
	GetRequestId() *string
	SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody
	GetResource() *GetResourceResponseBodyResource
}

type GetResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E871F6C0-2EFF-5790-A00D-C57543EEXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the file resource.
	Resource *GetResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s GetResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceResponseBody) GetResource() *GetResourceResponseBodyResource {
	return s.Resource
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody {
	s.Resource = v
	return s
}

func (s *GetResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceResponseBodyResource struct {
	// The time when the file resource was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file resource.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the file resource was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the file resource.
	//
	// example:
	//
	// OpenAPI_Test_Resource. py
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the file resource.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the file resource belongs.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the file resource. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
	//
	// example:
	//
	// {
	//
	//     "version": "1.1.0",
	//
	//     "kind": "Resource",
	//
	//     "spec": {
	//
	//         "fileResources": [
	//
	//             {
	//
	//                 "name": "OpenAPI_Test_Resource.py",
	//
	//                 "id": "631478864897630XXXX",
	//
	//                 "script": {
	//
	//                     "content": "",
	//
	//                     "path": "XX/OpenAPI_Test/Resource_Test/OpenAPI_Test_Resource.py",
	//
	//                     "runtime": {
	//
	//                         "command": "ODPS_PYTHON"
	//
	//                     }
	//
	//                 },
	//
	//                 "type": "python",
	//
	//                 "file": {
	//
	//                     "storage": {}
	//
	//                 },
	//
	//                 "datasource": {
	//
	//                     "name": "odps_first",
	//
	//                     "type": "odps"
	//
	//                 },
	//
	//                 "metadata": {
	//
	//                     "owner": "110755000425XXXX"
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

func (s GetResourceResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s GetResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetResourceResponseBodyResource) GetId() *int64 {
	return s.Id
}

func (s *GetResourceResponseBodyResource) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetResourceResponseBodyResource) GetName() *string {
	return s.Name
}

func (s *GetResourceResponseBodyResource) GetOwner() *string {
	return s.Owner
}

func (s *GetResourceResponseBodyResource) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceResponseBodyResource) GetSpec() *string {
	return s.Spec
}

func (s *GetResourceResponseBodyResource) SetCreateTime(v int64) *GetResourceResponseBodyResource {
	s.CreateTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetId(v int64) *GetResourceResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetModifyTime(v int64) *GetResourceResponseBodyResource {
	s.ModifyTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetName(v string) *GetResourceResponseBodyResource {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetOwner(v string) *GetResourceResponseBodyResource {
	s.Owner = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetProjectId(v int64) *GetResourceResponseBodyResource {
	s.ProjectId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetSpec(v string) *GetResourceResponseBodyResource {
	s.Spec = &v
	return s
}

func (s *GetResourceResponseBodyResource) Validate() error {
	return dara.Validate(s)
}
