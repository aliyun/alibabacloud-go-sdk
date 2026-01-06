// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateFunctionRequest
	GetId() *string
	SetProjectId(v int64) *UpdateFunctionRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateFunctionRequest
	GetSpec() *string
}

type UpdateFunctionRequest struct {
	// The unique identifier of the UDF.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK.. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique identifier of the UDF.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
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

func (s *UpdateFunctionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFunctionRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateFunctionRequest) SetId(v string) *UpdateFunctionRequest {
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
