// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybooks(v []*DescribeComponentPlaybookResponseBodyPlaybooks) *DescribeComponentPlaybookResponseBody
	GetPlaybooks() []*DescribeComponentPlaybookResponseBodyPlaybooks
	SetRequestId(v string) *DescribeComponentPlaybookResponseBody
	GetRequestId() *string
}

type DescribeComponentPlaybookResponseBody struct {
	// The list of component playbooks.
	Playbooks []*DescribeComponentPlaybookResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C5F5D6C9-DF1A-5381-92B1-39676F777D20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBody) GetPlaybooks() []*DescribeComponentPlaybookResponseBodyPlaybooks {
	return s.Playbooks
}

func (s *DescribeComponentPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentPlaybookResponseBody) SetPlaybooks(v []*DescribeComponentPlaybookResponseBodyPlaybooks) *DescribeComponentPlaybookResponseBody {
	s.Playbooks = v
	return s
}

func (s *DescribeComponentPlaybookResponseBody) SetRequestId(v string) *DescribeComponentPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBody) Validate() error {
	if s.Playbooks != nil {
		for _, item := range s.Playbooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComponentPlaybookResponseBodyPlaybooks struct {
	// The description of the component playbook.
	//
	// example:
	//
	// aegis_kill_process
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the component playbook.
	//
	// example:
	//
	// AegisKillQuara
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The input parameter configurations of the component playbook. The value is a JSON array.
	//
	// > For more information about the format, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
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
	// The input parameter type of the component playbook.
	//
	// template-ip: IP request template.
	//
	// template-file: file request template.
	//
	// template-process: process request template.
	//
	// custom: custom parameters.
	//
	// example:
	//
	// template-alert
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) GetDescription() *string {
	return s.Description
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) GetInputParams() *string {
	return s.InputParams
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDescription(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.Description = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDisplayName(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.DisplayName = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetInputParams(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.InputParams = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetParamType(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.ParamType = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) Validate() error {
	return dara.Validate(s)
}
