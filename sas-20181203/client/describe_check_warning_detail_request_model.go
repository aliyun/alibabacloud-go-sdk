// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v string) *DescribeCheckWarningDetailRequest
	GetCheckId() *string
	SetCheckWarningId(v int64) *DescribeCheckWarningDetailRequest
	GetCheckWarningId() *int64
	SetContainerName(v string) *DescribeCheckWarningDetailRequest
	GetContainerName() *string
	SetLang(v string) *DescribeCheckWarningDetailRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *DescribeCheckWarningDetailRequest
	GetSourceIp() *string
	SetUuid(v string) *DescribeCheckWarningDetailRequest
	GetUuid() *string
}

type DescribeCheckWarningDetailRequest struct {
	// The ID of the check item.
	//
	// > You can call the [ListCheckItemWarningSummary](~~ListCheckItemWarningSummary~~) operation to obtain the check item ID.	Notice: When this parameter is specified, the Uuid parameter is required..
	//
	// example:
	//
	// 1
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The alert ID of the check item.
	//
	// > To query the details of a specified check item, provide the alert ID of the check item. You can call the [DescribeCheckWarnings](~~DescribeCheckWarnings~~) operation to obtain this ID.
	//
	// 	Notice: This parameter is required when both Uuid and CheckId are empty..
	//
	// example:
	//
	// 98675301
	CheckWarningId *int64 `json:"CheckWarningId,omitempty" xml:"CheckWarningId,omitempty"`
	// The container name.
	//
	// example:
	//
	// test_container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the Alibaba Cloud account of the member accounts in the resource folder.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 103.25.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server to query.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.	Notice: When this parameter is specified, the CheckId parameter is required..
	//
	// example:
	//
	// 06125d19-6a02-4451-9f65-2083996e****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCheckWarningDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeCheckWarningDetailRequest) GetCheckWarningId() *int64 {
	return s.CheckWarningId
}

func (s *DescribeCheckWarningDetailRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeCheckWarningDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCheckWarningDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCheckWarningDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCheckWarningDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCheckWarningDetailRequest) SetCheckId(v string) *DescribeCheckWarningDetailRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetCheckWarningId(v int64) *DescribeCheckWarningDetailRequest {
	s.CheckWarningId = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetContainerName(v string) *DescribeCheckWarningDetailRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetLang(v string) *DescribeCheckWarningDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetSourceIp(v string) *DescribeCheckWarningDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetUuid(v string) *DescribeCheckWarningDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) Validate() error {
	return dara.Validate(s)
}
