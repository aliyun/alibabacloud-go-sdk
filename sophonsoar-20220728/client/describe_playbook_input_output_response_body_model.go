// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookInputOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *DescribePlaybookInputOutputResponseBodyConfig) *DescribePlaybookInputOutputResponseBody
	GetConfig() *DescribePlaybookInputOutputResponseBodyConfig
	SetRequestId(v string) *DescribePlaybookInputOutputResponseBody
	GetRequestId() *string
}

type DescribePlaybookInputOutputResponseBody struct {
	// The configurations.
	Config *DescribePlaybookInputOutputResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 688B4CCD-5272-5DCF-9D76-FE5EFEF545F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBody) GetConfig() *DescribePlaybookInputOutputResponseBodyConfig {
	return s.Config
}

func (s *DescribePlaybookInputOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookInputOutputResponseBody) SetConfig(v *DescribePlaybookInputOutputResponseBodyConfig) *DescribePlaybookInputOutputResponseBody {
	s.Config = v
	return s
}

func (s *DescribePlaybookInputOutputResponseBody) SetRequestId(v string) *DescribePlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlaybookInputOutputResponseBodyConfig struct {
	// The execution method of the playbook is in JSONObject format.
	ExeConfig *string `json:"ExeConfig,omitempty" xml:"ExeConfig,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
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
	// The output parameter configuration. This parameter is unavailable and is always left empty.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The input parameter type of the playbook. Valid values:
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
	// example:
	//
	// 9030076b-6733-4842-b05a-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) GetExeConfig() *string {
	return s.ExeConfig
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) GetInputParams() *string {
	return s.InputParams
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) GetOutputParams() *string {
	return s.OutputParams
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) GetParamType() *string {
	return s.ParamType
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetExeConfig(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.ExeConfig = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetInputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetOutputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.OutputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetParamType(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.ParamType = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetPlaybookUuid(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
