// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAssetDetailByUuidRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAssetDetailByUuidRequest
	GetSourceIp() *string
	SetUuid(v string) *DescribeAssetDetailByUuidRequest
	GetUuid() *string
}

type DescribeAssetDetailByUuidRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
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
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server to query.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9e6cad93-a379-46fd-a701-9bbf02f4****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAssetDetailByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAssetDetailByUuidRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAssetDetailByUuidRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAssetDetailByUuidRequest) SetLang(v string) *DescribeAssetDetailByUuidRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetDetailByUuidRequest) SetSourceIp(v string) *DescribeAssetDetailByUuidRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidRequest) SetUuid(v string) *DescribeAssetDetailByUuidRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidRequest) Validate() error {
	return dara.Validate(s)
}
