// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateResourceRequest
	GetProjectId() *int64
	SetResourceFile(v string) *CreateResourceRequest
	GetResourceFile() *string
	SetSpec(v string) *CreateResourceRequest
	GetSpec() *string
}

type CreateResourceRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
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
	//                 "name": "OpenAPITestResource.py",
	//
	//                 "script": {
	//
	//                     "content": "",
	//
	//                     "path": "XX/OpenAPITest/ResourcesTest/OpenAPITestResource.py",
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

func (s CreateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateResourceRequest) GetResourceFile() *string {
	return s.ResourceFile
}

func (s *CreateResourceRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateResourceRequest) SetProjectId(v int64) *CreateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceRequest) SetResourceFile(v string) *CreateResourceRequest {
	s.ResourceFile = &v
	return s
}

func (s *CreateResourceRequest) SetSpec(v string) *CreateResourceRequest {
	s.Spec = &v
	return s
}

func (s *CreateResourceRequest) Validate() error {
	return dara.Validate(s)
}
