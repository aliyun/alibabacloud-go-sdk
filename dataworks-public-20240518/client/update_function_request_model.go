// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateFunctionRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateFunctionRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateFunctionRequest
	GetSpec() *string
}

type UpdateFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "version": "1.1.0",
	//
	//     "kind": "Function",
	//
	//     "spec": {
	//
	//         "functions": [
	//
	//             {
	//
	//                 "name": "FunctionName",
	//
	//                 "script": {
	//
	//                     "content": "{\\"name\\": \\"FunctionName\\", \\"datasource\\": {\\"type\\": \\"odps\\", \\"name\\": \\"odps_first\\"}, \\"runtimeResource\\": {\\"resourceGroup\\": \\"S_res_group_XXXX_XXXX\\"}}",
	//
	//                     "path": "XXX/OpenAPI/Function/FunctionName",
	//
	//                     "runtime": {
	//
	//                         "command": "ODPS_FUNCTION"
	//
	//                     }
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
	//                 "runtimeResource": {
	//
	//                     "resourceGroup": "S_res_group_XXXX_XXXX"
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

func (s UpdateFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFunctionRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateFunctionRequest) SetId(v int64) *UpdateFunctionRequest {
	s.Id = &v
	return s
}

func (s *UpdateFunctionRequest) SetProjectId(v int64) *UpdateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpec(v string) *UpdateFunctionRequest {
	s.Spec = &v
	return s
}

func (s *UpdateFunctionRequest) Validate() error {
	return dara.Validate(s)
}
