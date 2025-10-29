// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateComponentRequest
	GetClientToken() *string
	SetProjectId(v int64) *CreateComponentRequest
	GetProjectId() *int64
	SetSpec(v string) *CreateComponentRequest
	GetSpec() *string
}

type CreateComponentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/2852607.html) operation to obtain the ID.
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
	//                 "id": "1234",
	//
	//                 "owner": "1234456",
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

func (s CreateComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateComponentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateComponentRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateComponentRequest) SetClientToken(v string) *CreateComponentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateComponentRequest) SetProjectId(v int64) *CreateComponentRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateComponentRequest) SetSpec(v string) *CreateComponentRequest {
	s.Spec = &v
	return s
}

func (s *CreateComponentRequest) Validate() error {
	return dara.Validate(s)
}
