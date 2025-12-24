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
	// The information about the predefined components.
	Playbooks []*DescribeComponentPlaybookResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	// The request ID.
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
	// The description of the predefined component.
	//
	// example:
	//
	// aegis_kill_process
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the predefined component.
	//
	// example:
	//
	// AegisKillQuara
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
	//
	// >  For more information, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
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
