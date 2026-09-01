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
	SetResourceDirectoryAccountId(v int64) *ModifyStartVulScanRequest
	GetResourceDirectoryAccountId() *int64
	SetTypes(v string) *ModifyStartVulScanRequest
	GetTypes() *string
	SetUuids(v string) *ModifyStartVulScanRequest
	GetUuids() *string
}

type ModifyStartVulScanRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Settings for the types of vulnerabilities to detect by using the one-click scan feature. Valid values:
	//
	// - **cve**: Linux software vulnerability.
	//
	// - **sys**: Windows system vulnerability.
	//
	// - **cms**: Web-CMS vulnerability.
	//
	// - **app**: application vulnerability detected by the web scanner.
	//
	// - **emg**: urgent vulnerability.
	//
	// - **image**: container image vulnerability.
	//
	// - **sca**: application vulnerability detected by software constituency parsing.
	//
	// > If this parameter is left empty, all vulnerability types are detected.
	//
	// example:
	//
	// "cve,sys,cms,app,emg"
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The UUIDs of the servers. Separate multiple UUIDs with commas (,).
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
