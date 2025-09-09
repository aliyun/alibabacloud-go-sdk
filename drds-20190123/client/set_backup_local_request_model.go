// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupLocalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *SetBackupLocalRequest
	GetDrdsInstanceId() *string
	SetHighSpaceUsageProtection(v string) *SetBackupLocalRequest
	GetHighSpaceUsageProtection() *string
	SetLocalLogRetentionHours(v string) *SetBackupLocalRequest
	GetLocalLogRetentionHours() *string
	SetLocalLogRetentionSpace(v string) *SetBackupLocalRequest
	GetLocalLogRetentionSpace() *string
}

type SetBackupLocalRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbgag23d13fds
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Specifies whether to enable the feature to forcibly delete binary log files if the used storage space reaches 90% of the total storage space or the remaining storage space is less than 5 GB. Valid values: 1 and 0. A value of 1 specifies to enable this feature. A value of 0 specifies not to enable this feature.
	//
	// example:
	//
	// 80
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance. Valid values: 0 to 168. Default value: 18. A value of 0 indicates that log backup files are not retained.
	//
	// example:
	//
	// 12
	LocalLogRetentionHours *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage space usage that is allowed for log files on the instance. Valid values: 0 to 50. Default value: 30.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
}

func (s SetBackupLocalRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBackupLocalRequest) GoString() string {
	return s.String()
}

func (s *SetBackupLocalRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetBackupLocalRequest) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *SetBackupLocalRequest) GetLocalLogRetentionHours() *string {
	return s.LocalLogRetentionHours
}

func (s *SetBackupLocalRequest) GetLocalLogRetentionSpace() *string {
	return s.LocalLogRetentionSpace
}

func (s *SetBackupLocalRequest) SetDrdsInstanceId(v string) *SetBackupLocalRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupLocalRequest) SetHighSpaceUsageProtection(v string) *SetBackupLocalRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionHours(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *SetBackupLocalRequest) SetLocalLogRetentionSpace(v string) *SetBackupLocalRequest {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *SetBackupLocalRequest) Validate() error {
	return dara.Validate(s)
}
