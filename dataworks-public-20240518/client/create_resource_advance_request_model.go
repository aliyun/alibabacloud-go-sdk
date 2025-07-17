// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateResourceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateResourceAdvanceRequest
	GetProjectId() *int64
	SetResourceFileObject(v io.Reader) *CreateResourceAdvanceRequest
	GetResourceFileObject() io.Reader
	SetSpec(v string) *CreateResourceAdvanceRequest
	GetSpec() *string
}

type CreateResourceAdvanceRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId          *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceFileObject io.Reader `json:"ResourceFile,omitempty" xml:"ResourceFile,omitempty"`
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

func (s CreateResourceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceAdvanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateResourceAdvanceRequest) GetResourceFileObject() io.Reader {
	return s.ResourceFileObject
}

func (s *CreateResourceAdvanceRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateResourceAdvanceRequest) SetProjectId(v int64) *CreateResourceAdvanceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceAdvanceRequest) SetResourceFileObject(v io.Reader) *CreateResourceAdvanceRequest {
	s.ResourceFileObject = v
	return s
}

func (s *CreateResourceAdvanceRequest) SetSpec(v string) *CreateResourceAdvanceRequest {
	s.Spec = &v
	return s
}

func (s *CreateResourceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
