// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAgentApplicationAgreement interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresAt(v string) *DataAgentApplicationAgreement
	GetExpiresAt() *string
	SetRegionId(v string) *DataAgentApplicationAgreement
	GetRegionId() *string
	SetSignedAt(v string) *DataAgentApplicationAgreement
	GetSignedAt() *string
	SetStatus(v string) *DataAgentApplicationAgreement
	GetStatus() *string
	SetWorkspaceId(v string) *DataAgentApplicationAgreement
	GetWorkspaceId() *string
}

type DataAgentApplicationAgreement struct {
	// The expiration time of the agreement in ISO 8601 UTC string format. This field is omitted if the agreement is unsigned. The agreement expires when this time is reached. The default validity period is 10 calendar years from the signing date.
	//
	// example:
	//
	// 2036-11-23T08:00:00Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The region of the workspace to which the agreement belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The signing time in ISO 8601 UTC string format. This field is omitted if the agreement is unsigned. Repeated signing within the validity period retains the original time, and re-signing after expiration updates the time.
	//
	// example:
	//
	// 2026-11-23T08:00:00Z
	SignedAt *string `json:"SignedAt,omitempty" xml:"SignedAt,omitempty"`
	// The agreement status. Valid values: NOT_SIGNED (unsigned), EXPIRED (expired), and SIGNED (signed and valid). Both NOT_SIGNED and EXPIRED are business statuses returned upon a successful query.
	//
	// example:
	//
	// SIGNED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The stable identifier of the workspace to which the agreement belongs.
	//
	// example:
	//
	// 62q1xxxxxxxxxxxxxxxxxxx0
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DataAgentApplicationAgreement) String() string {
	return dara.Prettify(s)
}

func (s DataAgentApplicationAgreement) GoString() string {
	return s.String()
}

func (s *DataAgentApplicationAgreement) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DataAgentApplicationAgreement) GetRegionId() *string {
	return s.RegionId
}

func (s *DataAgentApplicationAgreement) GetSignedAt() *string {
	return s.SignedAt
}

func (s *DataAgentApplicationAgreement) GetStatus() *string {
	return s.Status
}

func (s *DataAgentApplicationAgreement) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DataAgentApplicationAgreement) SetExpiresAt(v string) *DataAgentApplicationAgreement {
	s.ExpiresAt = &v
	return s
}

func (s *DataAgentApplicationAgreement) SetRegionId(v string) *DataAgentApplicationAgreement {
	s.RegionId = &v
	return s
}

func (s *DataAgentApplicationAgreement) SetSignedAt(v string) *DataAgentApplicationAgreement {
	s.SignedAt = &v
	return s
}

func (s *DataAgentApplicationAgreement) SetStatus(v string) *DataAgentApplicationAgreement {
	s.Status = &v
	return s
}

func (s *DataAgentApplicationAgreement) SetWorkspaceId(v string) *DataAgentApplicationAgreement {
	s.WorkspaceId = &v
	return s
}

func (s *DataAgentApplicationAgreement) Validate() error {
	return dara.Validate(s)
}
