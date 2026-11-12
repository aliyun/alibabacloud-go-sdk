// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsV1MigrationAdmission interface {
	dara.Model
	String() string
	GoString() string
	SetNewMigrationAllowed(v bool) *CmsV1MigrationAdmission
	GetNewMigrationAllowed() *bool
	SetPolicyRevision(v string) *CmsV1MigrationAdmission
	GetPolicyRevision() *string
	SetReason(v string) *CmsV1MigrationAdmission
	GetReason() *string
}

type CmsV1MigrationAdmission struct {
	// The real-time new migration eligibility evaluated for the current target primary account, which does not replace RAM, account ownership, or workspace validations.
	//
	// example:
	//
	// false
	NewMigrationAllowed *bool `json:"newMigrationAllowed,omitempty" xml:"newMigrationAllowed,omitempty"`
	// The summary of the normalized policy, which excludes the UID list and is not an authorization credential.
	//
	// example:
	//
	// 03ac8b5b92cb9c2675119d64c31d2520c069647aa9b88979df1bb115472c02b4
	PolicyRevision *string `json:"policyRevision,omitempty" xml:"policyRevision,omitempty"`
	// The reason why the new migration is denied. CMSV1_MIGRATION_NOT_ADMITTED indicates that the migration is not admitted, and CMSV1_MIGRATION_POLICY_UNAVAILABLE indicates that the policy is unavailable. Historical tasks remain queryable.
	//
	// example:
	//
	// CMSV1_MIGRATION_NOT_ADMITTED
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CmsV1MigrationAdmission) String() string {
	return dara.Prettify(s)
}

func (s CmsV1MigrationAdmission) GoString() string {
	return s.String()
}

func (s *CmsV1MigrationAdmission) GetNewMigrationAllowed() *bool {
	return s.NewMigrationAllowed
}

func (s *CmsV1MigrationAdmission) GetPolicyRevision() *string {
	return s.PolicyRevision
}

func (s *CmsV1MigrationAdmission) GetReason() *string {
	return s.Reason
}

func (s *CmsV1MigrationAdmission) SetNewMigrationAllowed(v bool) *CmsV1MigrationAdmission {
	s.NewMigrationAllowed = &v
	return s
}

func (s *CmsV1MigrationAdmission) SetPolicyRevision(v string) *CmsV1MigrationAdmission {
	s.PolicyRevision = &v
	return s
}

func (s *CmsV1MigrationAdmission) SetReason(v string) *CmsV1MigrationAdmission {
	s.Reason = &v
	return s
}

func (s *CmsV1MigrationAdmission) Validate() error {
	return dara.Validate(s)
}
