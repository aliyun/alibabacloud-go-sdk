// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAssetDetailByUuidsRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeAssetDetailByUuidsRequest
	GetResourceDirectoryAccountId() *int64
	SetUuids(v string) *DescribeAssetDetailByUuidsRequest
	GetUuids() *string
}

type DescribeAssetDetailByUuidsRequest struct {
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
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the IDs of Alibaba Cloud accounts.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The UUIDs of the instances. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0687b17f-2a36-4e5****,0687b17f-2a36-4e****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeAssetDetailByUuidsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAssetDetailByUuidsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeAssetDetailByUuidsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeAssetDetailByUuidsRequest) SetLang(v string) *DescribeAssetDetailByUuidsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetDetailByUuidsRequest) SetResourceDirectoryAccountId(v int64) *DescribeAssetDetailByUuidsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsRequest) SetUuids(v string) *DescribeAssetDetailByUuidsRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeAssetDetailByUuidsRequest) Validate() error {
	return dara.Validate(s)
}
