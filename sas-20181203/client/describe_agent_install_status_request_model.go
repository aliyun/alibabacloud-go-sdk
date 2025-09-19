// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentInstallStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAgentInstallStatusRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAgentInstallStatusRequest
	GetSourceIp() *string
	SetUuids(v string) *DescribeAgentInstallStatusRequest
	GetUuids() *string
}

type DescribeAgentInstallStatusRequest struct {
	// The language of the content within the request and response. Valid value:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 59.46.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUIDs of servers. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-eae014a7-16c4-4d4e-9f03-5208f4dc****,inet-eae047da-1e5a-41ce-828d-47606e9b****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeAgentInstallStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentInstallStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgentInstallStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAgentInstallStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAgentInstallStatusRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeAgentInstallStatusRequest) SetLang(v string) *DescribeAgentInstallStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAgentInstallStatusRequest) SetSourceIp(v string) *DescribeAgentInstallStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAgentInstallStatusRequest) SetUuids(v string) *DescribeAgentInstallStatusRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeAgentInstallStatusRequest) Validate() error {
	return dara.Validate(s)
}
