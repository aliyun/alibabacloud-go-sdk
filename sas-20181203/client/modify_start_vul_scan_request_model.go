// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStartVulScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyStartVulScanRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyStartVulScanRequest
	GetDryRun() *bool
	SetResourceDirectoryAccountId(v int64) *ModifyStartVulScanRequest
	GetResourceDirectoryAccountId() *int64
	SetTypes(v string) *ModifyStartVulScanRequest
	GetTypes() *string
	SetUuids(v string) *ModifyStartVulScanRequest
	GetUuids() *string
}

type ModifyStartVulScanRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - true: performs only a dry run without performing the actual operation.
	//
	// - false: performs the actual request.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Settings for the vulnerability types to be detected by the one-click scan feature. Valid values:
	//
	// - **cve**: Linux software vulnerability.
	//
	// - **sys**: Windows system vulnerability.
	//
	// - **cms**: Web-CMS vulnerability.
	//
	// - **app**: Application vulnerability detected by the web scanner.
	//
	// - **emg**: Emergency vulnerability.
	//
	// - **image**: Container image vulnerability.
	//
	// - **sca**: Application vulnerability detected by software constituency parsing.
	//
	// > If this parameter is left empty, all vulnerability types are detected.
	//
	// example:
	//
	// "cve,sys,cms,app,emg"
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The list of server UUIDs. Separate multiple UUIDs with commas (,).
	//
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/421726.html) operation to obtain this parameter.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyStartVulScanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStartVulScanRequest) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyStartVulScanRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyStartVulScanRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyStartVulScanRequest) GetTypes() *string {
	return s.Types
}

func (s *ModifyStartVulScanRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyStartVulScanRequest) SetClientToken(v string) *ModifyStartVulScanRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetDryRun(v bool) *ModifyStartVulScanRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetResourceDirectoryAccountId(v int64) *ModifyStartVulScanRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetTypes(v string) *ModifyStartVulScanRequest {
	s.Types = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetUuids(v string) *ModifyStartVulScanRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyStartVulScanRequest) Validate() error {
	return dara.Validate(s)
}
