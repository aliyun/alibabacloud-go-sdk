// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResultList(v []*DescribeCheckResultResponseBodyCheckResultList) *DescribeCheckResultResponseBody
	GetCheckResultList() []*DescribeCheckResultResponseBodyCheckResultList
	SetRequestId(v string) *DescribeCheckResultResponseBody
	GetRequestId() *string
}

type DescribeCheckResultResponseBody struct {
	// The check results.
	CheckResultList []*DescribeCheckResultResponseBodyCheckResultList `json:"CheckResultList,omitempty" xml:"CheckResultList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 571B2642-BF51-5BDD-906B-D2340DB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckResultResponseBody) GetCheckResultList() []*DescribeCheckResultResponseBodyCheckResultList {
	return s.CheckResultList
}

func (s *DescribeCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckResultResponseBody) SetCheckResultList(v []*DescribeCheckResultResponseBodyCheckResultList) *DescribeCheckResultResponseBody {
	s.CheckResultList = v
	return s
}

func (s *DescribeCheckResultResponseBody) SetRequestId(v string) *DescribeCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckResultResponseBodyCheckResultList struct {
	// The compliance status. Valid values:
	//
	// 	- **1**: compliant
	//
	// 	- **0**: non-compliant
	//
	// example:
	//
	// 1
	ComplianceStatus *int32 `json:"ComplianceStatus,omitempty" xml:"ComplianceStatus,omitempty"`
	// The name of the corresponding section. Valid values:
	//
	// 	- **information_classification**: information classification
	//
	// 	- **information_mark**: information labeling
	//
	// 	- **network_security_policy**: access to networks and network services
	//
	// 	- **login_control**: secure logon procedures
	//
	// 	- **week_password**: password management system
	//
	// 	- **key_manage**: key management
	//
	// 	- **malicious_software**: protection against malware
	//
	// 	- **information_backup**: information backup
	//
	// 	- **audit_policy**: information system audit control mechanisms
	//
	// example:
	//
	// information_mark
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCheckResultResponseBodyCheckResultList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckResultResponseBodyCheckResultList) GoString() string {
	return s.String()
}

func (s *DescribeCheckResultResponseBodyCheckResultList) GetComplianceStatus() *int32 {
	return s.ComplianceStatus
}

func (s *DescribeCheckResultResponseBodyCheckResultList) GetName() *string {
	return s.Name
}

func (s *DescribeCheckResultResponseBodyCheckResultList) SetComplianceStatus(v int32) *DescribeCheckResultResponseBodyCheckResultList {
	s.ComplianceStatus = &v
	return s
}

func (s *DescribeCheckResultResponseBodyCheckResultList) SetName(v string) *DescribeCheckResultResponseBodyCheckResultList {
	s.Name = &v
	return s
}

func (s *DescribeCheckResultResponseBodyCheckResultList) Validate() error {
	return dara.Validate(s)
}
