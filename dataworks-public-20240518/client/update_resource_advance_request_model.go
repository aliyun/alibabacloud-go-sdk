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
	SetId(v string) *UpdateResourceAdvanceRequest
	GetId() *string
	SetProjectId(v int64) *UpdateResourceAdvanceRequest
	GetProjectId() *int64
	SetResourceFileObject(v io.Reader) *UpdateResourceAdvanceRequest
	GetResourceFileObject() io.Reader
	SetSpec(v string) *UpdateResourceAdvanceRequest
	GetSpec() *string
}

type UpdateResourceAdvanceRequest struct {
	// The unique identifier of the Data Studio file resource.
	//
	// >  This field is of type Long in SDK versions prior to 8.0.0, and of type String in SDK version 8.0.0 and later. This change does not affect the normal use of the SDK; parameters are still returned according to the type defined in the SDK. Compilation failures due to the type change may occur only when upgrading the SDK across version 8.0.0, in which case users need to manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The specific file stream or OSS download link contained in the resource.
	//
	// >  This field allows users to provide a file stream or an OSS download link. When providing an OSS download link, ensure that the OSS link is publicly accessible. A presigned URL is recommended.
	//
	// example:
	//
	// http://bucketname1.oss-cn-shanghai.aliyuncs.com/example
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

func (s *UpdateResourceAdvanceRequest) GetId() *string {
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

func (s *UpdateResourceAdvanceRequest) SetId(v string) *UpdateResourceAdvanceRequest {
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
