// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryUserProvisioningEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *RetryUserProvisioningEventRequest
	GetDirectoryId() *string
	SetDuplicationStrategy(v string) *RetryUserProvisioningEventRequest
	GetDuplicationStrategy() *string
	SetEventId(v string) *RetryUserProvisioningEventRequest
	GetEventId() *string
}

type RetryUserProvisioningEventRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The conflict handling policy. The policy is used when a RAM user has the same username as the CloudSSO user who is synchronized to RAM. Valid values:
	//
	// 	- KeepBoth: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system creates a RAM user whose username is the username of the CloudSSO user plus the suffix `_sso`.
	//
	// 	- TakeOver: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system replaces the RAM user with the CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	// The ID of the RAM user provisioning event.
	//
	// You can call the [ListUserProvisioningEvents](https://help.aliyun.com/document_detail/2636305.html) operation to query the value of `EventId`.
	//
	// example:
	//
	// upe-wjKyNDmZvyZOiRcJ****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s RetryUserProvisioningEventRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryUserProvisioningEventRequest) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryUserProvisioningEventRequest) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *RetryUserProvisioningEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *RetryUserProvisioningEventRequest) SetDirectoryId(v string) *RetryUserProvisioningEventRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryUserProvisioningEventRequest) SetDuplicationStrategy(v string) *RetryUserProvisioningEventRequest {
	s.DuplicationStrategy = &v
	return s
}

func (s *RetryUserProvisioningEventRequest) SetEventId(v string) *RetryUserProvisioningEventRequest {
	s.EventId = &v
	return s
}

func (s *RetryUserProvisioningEventRequest) Validate() error {
	return dara.Validate(s)
}
