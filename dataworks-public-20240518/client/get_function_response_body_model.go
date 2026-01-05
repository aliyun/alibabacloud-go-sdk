// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFunction(v *GetFunctionResponseBodyFunction) *GetFunctionResponseBody
	GetFunction() *GetFunctionResponseBodyFunction
	SetRequestId(v string) *GetFunctionResponseBody
	GetRequestId() *string
}

type GetFunctionResponseBody struct {
	// UDF function details
	Function *GetFunctionResponseBodyFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6CF95929-6D12-5A88-8CC3-4B2F4C2EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) GetFunction() *GetFunctionResponseBodyFunction {
	return s.Function
}

func (s *GetFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionResponseBody) SetFunction(v *GetFunctionResponseBodyFunction) *GetFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetFunctionResponseBody) SetRequestId(v string) *GetFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionResponseBody) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionResponseBodyFunction struct {
	// The time when the UDF was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique identifier of the UDF function.
	//
	// >  This field is of type Long in SDK versions prior to 8.0.0, and of type String in SDK version 8.0.0 and later. This change does not affect the normal use of the SDK; parameters are still returned according to the type defined in the SDK. Compilation failures due to the type change may occur only when upgrading the SDK across version 8.0.0, in which case users need to manually correct the data type.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the UDF was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724506661000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the UDF.
	//
	// example:
	//
	// Function name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the UDF.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace to which the UDF belongs.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
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
	//                 "name": "Function_Name",
	//
	//                 "id": "580667964888595XXXX",
	//
	//                 "script": {
	//
	//                     "content": "{  \\"uuid\\": \\"580667964888595XXXX\\",  \\"name\\": \\"Function_Name\\",  \\"datasource\\": {    \\"type\\": \\"odps\\",    \\"name\\": \\"odps_first\\"  },  \\"runtimeResource\\": {    \\"resourceGroup\\": \\"S_res_group_XXXX_XXXX\\",    \\"resourceGroupId\\": 6591XXXX  }}",
	//
	//                     "path": "XXX/OpenAPI/Function/Function_Name",
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
	//                     "resourceGroup": "S_res_group_XXXX_XXXX",
	//
	//                     "id": "723932906364267XXXX",
	//
	//                     "resourceGroupId": "6591XXXX"
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

func (s GetFunctionResponseBodyFunction) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBodyFunction) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFunctionResponseBodyFunction) GetId() *string {
	return s.Id
}

func (s *GetFunctionResponseBodyFunction) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetFunctionResponseBodyFunction) GetName() *string {
	return s.Name
}

func (s *GetFunctionResponseBodyFunction) GetOwner() *string {
	return s.Owner
}

func (s *GetFunctionResponseBodyFunction) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFunctionResponseBodyFunction) GetSpec() *string {
	return s.Spec
}

func (s *GetFunctionResponseBodyFunction) SetCreateTime(v int64) *GetFunctionResponseBodyFunction {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetId(v string) *GetFunctionResponseBodyFunction {
	s.Id = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetModifyTime(v int64) *GetFunctionResponseBodyFunction {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetName(v string) *GetFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetOwner(v string) *GetFunctionResponseBodyFunction {
	s.Owner = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetProjectId(v int64) *GetFunctionResponseBodyFunction {
	s.ProjectId = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetSpec(v string) *GetFunctionResponseBodyFunction {
	s.Spec = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) Validate() error {
	return dara.Validate(s)
}
