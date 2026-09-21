// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSuspEventQuaraFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *RollbackSuspEventQuaraFileRequest
	GetFrom() *string
	SetQuaraFileId(v int32) *RollbackSuspEventQuaraFileRequest
	GetQuaraFileId() *int32
	SetResourceDirectoryAccountId(v int64) *RollbackSuspEventQuaraFileRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *RollbackSuspEventQuaraFileRequest
	GetSourceIp() *string
}

type RollbackSuspEventQuaraFileRequest struct {
	// The source of the request. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the quarantined file. You can call [DescribeSuspEventQuaraFiles](~~DescribeSuspEventQuaraFiles~~) to obtain this value from the Id field in the response. This parameter is required. If this parameter is not specified, the API returns HTTP 400 with error code -101.
	//
	// Before you call this operation, make sure that the Security Center agent is installed on the ECS instance, and that file-related security events and corresponding quarantined files exist. After a file is quarantined, call DescribeSuspEventQuaraFiles to query the quarantined file ID, and then call this operation to restore the file.
	//
	// example:
	//
	// 3921797
	QuaraFileId *int32 `json:"QuaraFileId,omitempty" xml:"QuaraFileId,omitempty"`
	// The Alibaba Cloud account ID of the member account in the resource directory.
	//
	// >You can call [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) GetFrom() *string {
	return s.From
}

func (s *RollbackSuspEventQuaraFileRequest) GetQuaraFileId() *int32 {
	return s.QuaraFileId
}

func (s *RollbackSuspEventQuaraFileRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *RollbackSuspEventQuaraFileRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *RollbackSuspEventQuaraFileRequest) SetFrom(v string) *RollbackSuspEventQuaraFileRequest {
	s.From = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFileId(v int32) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFileId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetResourceDirectoryAccountId(v int64) *RollbackSuspEventQuaraFileRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetSourceIp(v string) *RollbackSuspEventQuaraFileRequest {
	s.SourceIp = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) Validate() error {
	return dara.Validate(s)
}
