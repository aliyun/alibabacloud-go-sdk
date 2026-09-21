// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmMachineCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *GetAlarmMachineCountRequest
	GetFrom() *string
	SetResourceDirectoryAccountId(v int64) *GetAlarmMachineCountRequest
	GetResourceDirectoryAccountId() *int64
}

type GetAlarmMachineCountRequest struct {
	// The request source identifier. Set this parameter to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The Alibaba Cloud account ID of the member account in the resource directory.
	//
	// >Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s GetAlarmMachineCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmMachineCountRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmMachineCountRequest) GetFrom() *string {
	return s.From
}

func (s *GetAlarmMachineCountRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetAlarmMachineCountRequest) SetFrom(v string) *GetAlarmMachineCountRequest {
	s.From = &v
	return s
}

func (s *GetAlarmMachineCountRequest) SetResourceDirectoryAccountId(v int64) *GetAlarmMachineCountRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetAlarmMachineCountRequest) Validate() error {
	return dara.Validate(s)
}
