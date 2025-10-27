// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPolicy(v *CreateBackupPolicyResponseBodyBackupPolicy) *CreateBackupPolicyResponseBody
	GetBackupPolicy() *CreateBackupPolicyResponseBodyBackupPolicy
	SetRequestId(v string) *CreateBackupPolicyResponseBody
	GetRequestId() *string
}

type CreateBackupPolicyResponseBody struct {
	// The information about the anti-ransomware policy.
	BackupPolicy *CreateBackupPolicyResponseBodyBackupPolicy `json:"BackupPolicy,omitempty" xml:"BackupPolicy,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponseBody) GetBackupPolicy() *CreateBackupPolicyResponseBodyBackupPolicy {
	return s.BackupPolicy
}

func (s *CreateBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupPolicyResponseBody) SetBackupPolicy(v *CreateBackupPolicyResponseBodyBackupPolicy) *CreateBackupPolicyResponseBody {
	s.BackupPolicy = v
	return s
}

func (s *CreateBackupPolicyResponseBody) SetRequestId(v string) *CreateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPolicyResponseBody) Validate() error {
	if s.BackupPolicy != nil {
		if err := s.BackupPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBackupPolicyResponseBodyBackupPolicy struct {
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 1301575
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// >  After you create an anti-ransomware policy, the policy is enabled by default.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBackupPolicyResponseBodyBackupPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyResponseBodyBackupPolicy) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponseBodyBackupPolicy) GetId() *string {
	return s.Id
}

func (s *CreateBackupPolicyResponseBodyBackupPolicy) GetStatus() *string {
	return s.Status
}

func (s *CreateBackupPolicyResponseBodyBackupPolicy) SetId(v string) *CreateBackupPolicyResponseBodyBackupPolicy {
	s.Id = &v
	return s
}

func (s *CreateBackupPolicyResponseBodyBackupPolicy) SetStatus(v string) *CreateBackupPolicyResponseBodyBackupPolicy {
	s.Status = &v
	return s
}

func (s *CreateBackupPolicyResponseBodyBackupPolicy) Validate() error {
	return dara.Validate(s)
}
