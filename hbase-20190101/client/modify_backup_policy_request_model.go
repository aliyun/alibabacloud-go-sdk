// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyBackupPolicyRequest
	GetClusterId() *string
	SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupEndTimeUTC() *string
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupStartTimeUTC() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
}

type ModifyBackupPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The UTC time when the backup ends.
	//
	// example:
	//
	// 18:00Z
	PreferredBackupEndTimeUTC *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	// The backup cycle. Valid values:
	//
	// - Monday: performs backup every Monday.
	//
	// - Tuesday: performs backup every Tuesday.
	//
	// - Wednesday: performs backup every Wednesday.
	//
	// - Thursday: performs backup every Thursday.
	//
	// - Friday: performs backup every Friday.
	//
	// - Saturday: performs backup every Saturday.
	//
	// - Sunday: performs backup every Sunday.
	//
	// This parameter is required.
	//
	// example:
	//
	// Thursday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The UTC time when the backup starts.
	//
	// example:
	//
	// 17:00Z
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	// The backup time range in the current time zone. The interval is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01:00-02:00
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupEndTimeUTC() *string {
	return s.PreferredBackupEndTimeUTC
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupStartTimeUTC() *string {
	return s.PreferredBackupStartTimeUTC
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) SetClusterId(v string) *ModifyBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
