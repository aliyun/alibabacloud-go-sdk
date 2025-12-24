// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutePlaybooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybookMetrics(v []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics) *DescribeExecutePlaybooksResponseBody
	GetPlaybookMetrics() []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics
	SetRequestId(v string) *DescribeExecutePlaybooksResponseBody
	GetRequestId() *string
}

type DescribeExecutePlaybooksResponseBody struct {
	// The playbook.
	PlaybookMetrics []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics `json:"PlaybookMetrics,omitempty" xml:"PlaybookMetrics,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 88A39217-2802-5B1E-BA2B-CF1BBC43C1F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBody) GetPlaybookMetrics() []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	return s.PlaybookMetrics
}

func (s *DescribeExecutePlaybooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExecutePlaybooksResponseBody) SetPlaybookMetrics(v []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics) *DescribeExecutePlaybooksResponseBody {
	s.PlaybookMetrics = v
	return s
}

func (s *DescribeExecutePlaybooksResponseBody) SetRequestId(v string) *DescribeExecutePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBody) Validate() error {
	if s.PlaybookMetrics != nil {
		for _, item := range s.PlaybookMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExecutePlaybooksResponseBodyPlaybookMetrics struct {
	// The playbook description.
	//
	// example:
	//
	// a demo playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The playbook name.
	//
	// example:
	//
	// demo_playbook
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The configuration of the input parameter. The value is a JSON array.
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
	ParamConfig *string `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty"`
	// The input parameter type of the playbook.
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
	// The playbook UUID.
	//
	// example:
	//
	// c5c88b5e-97ca-435d-8c20-2xxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GetDescription() *string {
	return s.Description
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GetParamConfig() *string {
	return s.ParamConfig
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDescription(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Description = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDisplayName(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamConfig(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamConfig = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamType(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetUuid(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Uuid = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) Validate() error {
	return dara.Validate(s)
}
