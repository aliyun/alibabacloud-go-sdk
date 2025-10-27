// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBackupMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachines(v []*DescribeUserBackupMachinesResponseBodyMachines) *DescribeUserBackupMachinesResponseBody
	GetMachines() []*DescribeUserBackupMachinesResponseBodyMachines
	SetRequestId(v string) *DescribeUserBackupMachinesResponseBody
	GetRequestId() *string
}

type DescribeUserBackupMachinesResponseBody struct {
	// An array consisting of the servers to which the anti-ransomware policy is applied.
	Machines []*DescribeUserBackupMachinesResponseBodyMachines `json:"Machines,omitempty" xml:"Machines,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA04B21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBackupMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBackupMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBackupMachinesResponseBody) GetMachines() []*DescribeUserBackupMachinesResponseBodyMachines {
	return s.Machines
}

func (s *DescribeUserBackupMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBackupMachinesResponseBody) SetMachines(v []*DescribeUserBackupMachinesResponseBodyMachines) *DescribeUserBackupMachinesResponseBody {
	s.Machines = v
	return s
}

func (s *DescribeUserBackupMachinesResponseBody) SetRequestId(v string) *DescribeUserBackupMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBackupMachinesResponseBody) Validate() error {
	if s.Machines != nil {
		for _, item := range s.Machines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserBackupMachinesResponseBodyMachines struct {
	// The ID of the anti-ransomware policy that is applied to the server.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anti-ransomware policy that is applied to the server.
	//
	// example:
	//
	// policy_name_A
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The UUID of the server to which the anti-ransomware policy is applied.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeUserBackupMachinesResponseBodyMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBackupMachinesResponseBodyMachines) GoString() string {
	return s.String()
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) GetId() *int64 {
	return s.Id
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) SetId(v int64) *DescribeUserBackupMachinesResponseBodyMachines {
	s.Id = &v
	return s
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) SetPolicyName(v string) *DescribeUserBackupMachinesResponseBodyMachines {
	s.PolicyName = &v
	return s
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) SetUuid(v string) *DescribeUserBackupMachinesResponseBodyMachines {
	s.Uuid = &v
	return s
}

func (s *DescribeUserBackupMachinesResponseBodyMachines) Validate() error {
	return dara.Validate(s)
}
