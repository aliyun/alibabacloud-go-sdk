// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *GetCheckConfigRequest
	GetConfigId() *string
	SetResourceDirectoryAccountId(v int64) *GetCheckConfigRequest
	GetResourceDirectoryAccountId() *int64
}

type GetCheckConfigRequest struct {
	// The policy ID.
	//
	// example:
	//
	// 22222f91-57af-4907-87f1-f9046da*****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource directory member account belongs.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1234567890123456
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s GetCheckConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCheckConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetCheckConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetCheckConfigRequest) SetConfigId(v string) *GetCheckConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetCheckConfigRequest) SetResourceDirectoryAccountId(v int64) *GetCheckConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetCheckConfigRequest) Validate() error {
	return dara.Validate(s)
}
