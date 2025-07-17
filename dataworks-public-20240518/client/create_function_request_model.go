// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateFunctionRequest
	GetProjectId() *int64
	SetSpec(v string) *CreateFunctionRequest
	GetSpec() *string
}

type CreateFunctionRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "version": "1.1.0",
	//
	//   "kind": "Function",
	//
	//   "spec": {
	//
	//     "functions": [
	//
	//       {
	//
	//         "name": "function name",
	//
	//         "script": {
	//
	//           "content": "{\\"name\\": \\"function name\\", \\"datasource\\": {\\"type\\": \\"ODPS\\", \\"name\\": \\"ODPS_first\\"}, \\"runtimeResource\\": {\\"resourceGroup\\": \\"s_res_group_xx_xxxx\\"}}",
	//
	//           "path": "XXX/OpenAPI/function/function name",
	//
	//           "runtime": {
	//
	//             "command": "ODPS_FUNCTION"
	//
	//           }
	//
	//         },
	//
	//         "datasource": {
	//
	//           "name": "ODPS_first",
	//
	//           "type": "ODPS"
	//
	//         },
	//
	//         "runtimeResource": {
	//
	//           "resourceGroup": "S_res_group_XXXX_XXXX"
	//
	//         }
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateFunctionRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateFunctionRequest) SetProjectId(v int64) *CreateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFunctionRequest) SetSpec(v string) *CreateFunctionRequest {
	s.Spec = &v
	return s
}

func (s *CreateFunctionRequest) Validate() error {
	return dara.Validate(s)
}
