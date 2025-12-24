// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookInputOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExeConfig(v string) *ModifyPlaybookInputOutputRequest
	GetExeConfig() *string
	SetInputParams(v string) *ModifyPlaybookInputOutputRequest
	GetInputParams() *string
	SetLang(v string) *ModifyPlaybookInputOutputRequest
	GetLang() *string
	SetOutputParams(v string) *ModifyPlaybookInputOutputRequest
	GetOutputParams() *string
	SetParamType(v string) *ModifyPlaybookInputOutputRequest
	GetParamType() *string
	SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest
	GetPlaybookUuid() *string
}

type ModifyPlaybookInputOutputRequest struct {
	// The executed mode of a playbook. The value is a JSON array.
	ExeConfig *string `json:"ExeConfig,omitempty" xml:"ExeConfig,omitempty"`
	// The configuration of the input parameters. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The configuration of the output parameters. This parameter is unavailable. Leave it empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The input parameter type.
	//
	// 	- **template-ip**
	//
	// 	- **template-file**
	//
	// 	- **template-process**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInputOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputRequest) GetExeConfig() *string {
	return s.ExeConfig
}

func (s *ModifyPlaybookInputOutputRequest) GetInputParams() *string {
	return s.InputParams
}

func (s *ModifyPlaybookInputOutputRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyPlaybookInputOutputRequest) GetOutputParams() *string {
	return s.OutputParams
}

func (s *ModifyPlaybookInputOutputRequest) GetParamType() *string {
	return s.ParamType
}

func (s *ModifyPlaybookInputOutputRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ModifyPlaybookInputOutputRequest) SetExeConfig(v string) *ModifyPlaybookInputOutputRequest {
	s.ExeConfig = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetInputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.InputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetLang(v string) *ModifyPlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetOutputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.OutputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetParamType(v string) *ModifyPlaybookInputOutputRequest {
	s.ParamType = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) Validate() error {
	return dara.Validate(s)
}
