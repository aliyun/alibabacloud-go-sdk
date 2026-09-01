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
	// The verification results.
	CheckTaskInfos []*VerifyPlaybookResponseBodyCheckTaskInfos `json:"CheckTaskInfos,omitempty" xml:"CheckTaskInfos,omitempty" type:"Repeated"`
	// The prerequisite check information for the playbook.
	Prerequisites []*VerifyPlaybookResponseBodyPrerequisites `json:"Prerequisites,omitempty" xml:"Prerequisites,omitempty" type:"Repeated"`
	// The ID of the request. Alibaba Cloud generates this unique identifier for the request. Use this ID to troubleshoot and locate issues.
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
	// The specific error message that is returned if the verification fails.
	//
	// example:
	//
	// Node [python3_3] doesn\\"t have the asset information
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the playbook node.
	//
	// example:
	//
	// python3_3
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The severity level of the verification message. Valid values:
	//
	// - **warn**: A warning message. An issue may occur when the playbook runs.
	//
	// - **error**: An error message. The playbook fails to be compiled.
	//
	// - **remind**: A suggestion. This does not affect publishing or running the playbook. Optimize the playbook format.
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
	// The check type. Valid values:
	//
	// - **role**: The name of the custom RAM role.
	//
	// - **policies**: The list of RAM system policies.
	//
	// example:
	//
	// role
	PrerequisiteType *string `json:"PrerequisiteType,omitempty" xml:"PrerequisiteType,omitempty"`
	// The check content. The value is determined as follows:
	//
	// - If PrerequisiteType is **role**, the value is the static field AliyunSiemSoarExecutionDefaultRole.
	//
	// - If PrerequisiteType is **policies**, the value is a collection of policy names.
	//
	// example:
	//
	// AliyunSiemSoarExecutionDefaultRole
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
