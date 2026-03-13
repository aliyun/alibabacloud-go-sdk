// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaUser interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v *QuotaUserResources) *QuotaUser
	GetResources() *QuotaUserResources
	SetUserId(v string) *QuotaUser
	GetUserId() *string
	SetUsername(v string) *QuotaUser
	GetUsername() *string
	SetWorkloadCount(v int32) *QuotaUser
	GetWorkloadCount() *int32
}

type QuotaUser struct {
	Resources     *QuotaUserResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	UserId        *string             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username      *string             `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkloadCount *int32              `json:"WorkloadCount,omitempty" xml:"WorkloadCount,omitempty"`
}

func (s QuotaUser) String() string {
	return dara.Prettify(s)
}

func (s QuotaUser) GoString() string {
	return s.String()
}

func (s *QuotaUser) GetResources() *QuotaUserResources {
	return s.Resources
}

func (s *QuotaUser) GetUserId() *string {
	return s.UserId
}

func (s *QuotaUser) GetUsername() *string {
	return s.Username
}

func (s *QuotaUser) GetWorkloadCount() *int32 {
	return s.WorkloadCount
}

func (s *QuotaUser) SetResources(v *QuotaUserResources) *QuotaUser {
	s.Resources = v
	return s
}

func (s *QuotaUser) SetUserId(v string) *QuotaUser {
	s.UserId = &v
	return s
}

func (s *QuotaUser) SetUsername(v string) *QuotaUser {
	s.Username = &v
	return s
}

func (s *QuotaUser) SetWorkloadCount(v int32) *QuotaUser {
	s.WorkloadCount = &v
	return s
}

func (s *QuotaUser) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuotaUserResources struct {
	Submitted *ResourceAmount `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	Used      *ResourceAmount `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s QuotaUserResources) String() string {
	return dara.Prettify(s)
}

func (s QuotaUserResources) GoString() string {
	return s.String()
}

func (s *QuotaUserResources) GetSubmitted() *ResourceAmount {
	return s.Submitted
}

func (s *QuotaUserResources) GetUsed() *ResourceAmount {
	return s.Used
}

func (s *QuotaUserResources) SetSubmitted(v *ResourceAmount) *QuotaUserResources {
	s.Submitted = v
	return s
}

func (s *QuotaUserResources) SetUsed(v *ResourceAmount) *QuotaUserResources {
	s.Used = v
	return s
}

func (s *QuotaUserResources) Validate() error {
	if s.Submitted != nil {
		if err := s.Submitted.Validate(); err != nil {
			return err
		}
	}
	if s.Used != nil {
		if err := s.Used.Validate(); err != nil {
			return err
		}
	}
	return nil
}
