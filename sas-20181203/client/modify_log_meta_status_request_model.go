// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogMetaStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *ModifyLogMetaStatusRequest
	GetFrom() *string
	SetLogStore(v string) *ModifyLogMetaStatusRequest
	GetLogStore() *string
	SetProject(v string) *ModifyLogMetaStatusRequest
	GetProject() *string
	SetResourceDirectoryAccountId(v int64) *ModifyLogMetaStatusRequest
	GetResourceDirectoryAccountId() *int64
	SetStatus(v string) *ModifyLogMetaStatusRequest
	GetStatus() *string
}

type ModifyLogMetaStatusRequest struct {
	// The request source identifier. Set this parameter to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The name of the dedicated Logstore where logs are stored.
	//
	// >You can call the [DescribeLogMeta](~~DescribeLogMeta~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// aegis-log-login
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The project name.
	//
	// > You can call the [DescribeLogMeta](~~DescribeLogMeta~~) operation to obtain this parameter.
	//
	// example:
	//
	// aegis-log
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the member account in the resource directory (Alibaba Cloud account).
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The status to which you want to change the log. Valid values:
	//
	// - **enabled**: enabled
	//
	// - **disabled**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyLogMetaStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogMetaStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogMetaStatusRequest) GetFrom() *string {
	return s.From
}

func (s *ModifyLogMetaStatusRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *ModifyLogMetaStatusRequest) GetProject() *string {
	return s.Project
}

func (s *ModifyLogMetaStatusRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyLogMetaStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyLogMetaStatusRequest) SetFrom(v string) *ModifyLogMetaStatusRequest {
	s.From = &v
	return s
}

func (s *ModifyLogMetaStatusRequest) SetLogStore(v string) *ModifyLogMetaStatusRequest {
	s.LogStore = &v
	return s
}

func (s *ModifyLogMetaStatusRequest) SetProject(v string) *ModifyLogMetaStatusRequest {
	s.Project = &v
	return s
}

func (s *ModifyLogMetaStatusRequest) SetResourceDirectoryAccountId(v int64) *ModifyLogMetaStatusRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyLogMetaStatusRequest) SetStatus(v string) *ModifyLogMetaStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyLogMetaStatusRequest) Validate() error {
	return dara.Validate(s)
}
