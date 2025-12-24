// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckTaskInfos(v []*VerifyPlaybookResponseBodyCheckTaskInfos) *VerifyPlaybookResponseBody
	GetCheckTaskInfos() []*VerifyPlaybookResponseBodyCheckTaskInfos
	SetPrerequisites(v []*VerifyPlaybookResponseBodyPrerequisites) *VerifyPlaybookResponseBody
	GetPrerequisites() []*VerifyPlaybookResponseBodyPrerequisites
	SetRequestId(v string) *VerifyPlaybookResponseBody
	GetRequestId() *string
}

type VerifyPlaybookResponseBody struct {
	// The result of the verification.
	CheckTaskInfos []*VerifyPlaybookResponseBodyCheckTaskInfos `json:"CheckTaskInfos,omitempty" xml:"CheckTaskInfos,omitempty" type:"Repeated"`
	Prerequisites  []*VerifyPlaybookResponseBodyPrerequisites  `json:"Prerequisites,omitempty" xml:"Prerequisites,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0DFC9403-54EB-5672-B690-9AA93C9EBB54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBody) GetCheckTaskInfos() []*VerifyPlaybookResponseBodyCheckTaskInfos {
	return s.CheckTaskInfos
}

func (s *VerifyPlaybookResponseBody) GetPrerequisites() []*VerifyPlaybookResponseBodyPrerequisites {
	return s.Prerequisites
}

func (s *VerifyPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyPlaybookResponseBody) SetCheckTaskInfos(v []*VerifyPlaybookResponseBodyCheckTaskInfos) *VerifyPlaybookResponseBody {
	s.CheckTaskInfos = v
	return s
}

func (s *VerifyPlaybookResponseBody) SetPrerequisites(v []*VerifyPlaybookResponseBodyPrerequisites) *VerifyPlaybookResponseBody {
	s.Prerequisites = v
	return s
}

func (s *VerifyPlaybookResponseBody) SetRequestId(v string) *VerifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPlaybookResponseBody) Validate() error {
	if s.CheckTaskInfos != nil {
		for _, item := range s.CheckTaskInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Prerequisites != nil {
		for _, item := range s.Prerequisites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VerifyPlaybookResponseBodyCheckTaskInfos struct {
	// The error message returned when the playbook does not pass the check.
	//
	// example:
	//
	// Node [python3_3] doesn\\"t have the asset information
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the node in the playbook.
	//
	// example:
	//
	// python3_3
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The severity level of the verification information. Valid values:
	//
	// 	- warn: An issue may occur during playbook running.
	//
	// 	- error: The playbook cannot be compiled.
	//
	// 	- remind: The publishing and running of the playbook are not affected. We recommend that you optimize the playbook format.
	//
	// example:
	//
	// error
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) String() string {
	return dara.Prettify(s)
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) GetDetail() *string {
	return s.Detail
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetDetail(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.Detail = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetNodeName(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.NodeName = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetRiskLevel(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.RiskLevel = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) Validate() error {
	return dara.Validate(s)
}

type VerifyPlaybookResponseBodyPrerequisites struct {
	PrerequisiteType  *string `json:"PrerequisiteType,omitempty" xml:"PrerequisiteType,omitempty"`
	PrerequisiteValue *string `json:"PrerequisiteValue,omitempty" xml:"PrerequisiteValue,omitempty"`
}

func (s VerifyPlaybookResponseBodyPrerequisites) String() string {
	return dara.Prettify(s)
}

func (s VerifyPlaybookResponseBodyPrerequisites) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBodyPrerequisites) GetPrerequisiteType() *string {
	return s.PrerequisiteType
}

func (s *VerifyPlaybookResponseBodyPrerequisites) GetPrerequisiteValue() *string {
	return s.PrerequisiteValue
}

func (s *VerifyPlaybookResponseBodyPrerequisites) SetPrerequisiteType(v string) *VerifyPlaybookResponseBodyPrerequisites {
	s.PrerequisiteType = &v
	return s
}

func (s *VerifyPlaybookResponseBodyPrerequisites) SetPrerequisiteValue(v string) *VerifyPlaybookResponseBodyPrerequisites {
	s.PrerequisiteValue = &v
	return s
}

func (s *VerifyPlaybookResponseBodyPrerequisites) Validate() error {
	return dara.Validate(s)
}
