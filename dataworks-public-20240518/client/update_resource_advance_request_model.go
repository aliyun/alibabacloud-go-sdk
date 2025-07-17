// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUpdateResourceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateResourceAdvanceRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateResourceAdvanceRequest
	GetProjectId() *int64
	SetResourceFileObject(v io.Reader) *UpdateResourceAdvanceRequest
	GetResourceFileObject() io.Reader
	SetSpec(v string) *UpdateResourceAdvanceRequest
	GetSpec() *string
}

type UpdateResourceAdvanceRequest struct {
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

func (s UpdateResourceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceAdvanceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateResourceAdvanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateResourceAdvanceRequest) GetResourceFileObject() io.Reader {
	return s.ResourceFileObject
}

func (s *UpdateResourceAdvanceRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateResourceAdvanceRequest) SetId(v int64) *UpdateResourceAdvanceRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceAdvanceRequest) SetProjectId(v int64) *UpdateResourceAdvanceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceAdvanceRequest) SetResourceFileObject(v io.Reader) *UpdateResourceAdvanceRequest {
	s.ResourceFileObject = v
	return s
}

func (s *UpdateResourceAdvanceRequest) SetSpec(v string) *UpdateResourceAdvanceRequest {
	s.Spec = &v
	return s
}

func (s *UpdateResourceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
