// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSuspEventUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *SaveSuspEventUserSettingRequest
	GetFrom() *string
	SetLevelsOn(v string) *SaveSuspEventUserSettingRequest
	GetLevelsOn() *string
	SetResourceDirectoryAccountId(v int64) *SaveSuspEventUserSettingRequest
	GetResourceDirectoryAccountId() *int64
}

type SaveSuspEventUserSettingRequest struct {
	// The source of the exception event data. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The alert level for alert notifications. Valid values:
	//
	// - **remind**: Reminder.
	//
	// - **suspicious**: Suspicious.
	//
	// - **serious**: Urgent.
	//
	// example:
	//
	// suspicious,serious,remind
	LevelsOn *string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty"`
	// The ID of the member account in the resource directory.
	//
	// >You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s SaveSuspEventUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSuspEventUserSettingRequest) GoString() string {
	return s.String()
}

func (s *SaveSuspEventUserSettingRequest) GetFrom() *string {
	return s.From
}

func (s *SaveSuspEventUserSettingRequest) GetLevelsOn() *string {
	return s.LevelsOn
}

func (s *SaveSuspEventUserSettingRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *SaveSuspEventUserSettingRequest) SetFrom(v string) *SaveSuspEventUserSettingRequest {
	s.From = &v
	return s
}

func (s *SaveSuspEventUserSettingRequest) SetLevelsOn(v string) *SaveSuspEventUserSettingRequest {
	s.LevelsOn = &v
	return s
}

func (s *SaveSuspEventUserSettingRequest) SetResourceDirectoryAccountId(v int64) *SaveSuspEventUserSettingRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *SaveSuspEventUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
