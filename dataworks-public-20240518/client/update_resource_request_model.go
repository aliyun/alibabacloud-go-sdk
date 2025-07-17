// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateResourceRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateResourceRequest
	GetProjectId() *int64
	SetResourceFile(v string) *UpdateResourceRequest
	GetResourceFile() *string
	SetSpec(v string) *UpdateResourceRequest
	GetSpec() *string
}

type UpdateResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceFile *string `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
	// The FlowSpec field information about the file resource. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
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
	//                 "script": {
	//
	//                     "content": "",
	//
	//                     "path": "XX/OpenAPI_Test/Resources_Test/OpenAPI_Test_Resource.py",
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

func (s UpdateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateResourceRequest) GetResourceFile() *string {
	return s.ResourceFile
}

func (s *UpdateResourceRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateResourceRequest) SetId(v int64) *UpdateResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceRequest) SetProjectId(v int64) *UpdateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceRequest) SetResourceFile(v string) *UpdateResourceRequest {
	s.ResourceFile = &v
	return s
}

func (s *UpdateResourceRequest) SetSpec(v string) *UpdateResourceRequest {
	s.Spec = &v
	return s
}

func (s *UpdateResourceRequest) Validate() error {
	return dara.Validate(s)
}
