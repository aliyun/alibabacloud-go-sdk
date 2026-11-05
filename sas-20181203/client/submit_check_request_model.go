// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *SubmitCheckRequest
	GetConfigId() *string
	SetResourceDirectoryAccountId(v int64) *SubmitCheckRequest
	GetResourceDirectoryAccountId() *int64
	SetScanRange(v string) *SubmitCheckRequest
	GetScanRange() *string
	SetTaskSource(v string) *SubmitCheckRequest
	GetTaskSource() *string
}

type SubmitCheckRequest struct {
	// The policy ID. This parameter takes effect only when ScanRange is set to POLICY. You can obtain the value by calling the ListCheckConfigs operation. If this parameter is not specified or is set to default, the scan is triggered based on the system default policy. If the default policy does not exist, NoDefaultStrategy is returned. If another value is specified, the scan is triggered based on the check items of the specified policy. If the policy does not exist or does not belong to the current account, IllegalParam is returned. This parameter is ignored when ScanRange is set to FULL or FREE.
	//
	// example:
	//
	// 22222f91-57af-4907-87f1-f9046da****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Alibaba Cloud account to which the member accounts of the resource folder belong.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The check scope. Valid values:
	//
	// - **FULL**: Scans all check items.
	//
	// - **POLICY**: Scans check items of custom configurations.
	//
	// example:
	//
	// POLICY
	ScanRange *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The source of the task. Valid values:
	//
	// - **YAO_CHI**: ApsaraDB console.
	//
	// example:
	//
	// YAO_CHI
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s SubmitCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitCheckRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *SubmitCheckRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *SubmitCheckRequest) GetScanRange() *string {
	return s.ScanRange
}

func (s *SubmitCheckRequest) GetTaskSource() *string {
	return s.TaskSource
}

func (s *SubmitCheckRequest) SetConfigId(v string) *SubmitCheckRequest {
	s.ConfigId = &v
	return s
}

func (s *SubmitCheckRequest) SetResourceDirectoryAccountId(v int64) *SubmitCheckRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *SubmitCheckRequest) SetScanRange(v string) *SubmitCheckRequest {
	s.ScanRange = &v
	return s
}

func (s *SubmitCheckRequest) SetTaskSource(v string) *SubmitCheckRequest {
	s.TaskSource = &v
	return s
}

func (s *SubmitCheckRequest) Validate() error {
	return dara.Validate(s)
}
