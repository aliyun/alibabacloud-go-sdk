// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetImportantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportantCode(v int32) *ModifyAssetImportantRequest
	GetImportantCode() *int32
	SetUuidList(v string) *ModifyAssetImportantRequest
	GetUuidList() *string
}

type ModifyAssetImportantRequest struct {
	// The importance of the asset. Valid values:
	//
	// 	- **0**: test
	//
	// 	- **1**: normal
	//
	// 	- **2**: important
	//
	// example:
	//
	// 0
	ImportantCode *int32 `json:"ImportantCode,omitempty" xml:"ImportantCode,omitempty"`
	// The UUIDs of servers. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 076a446d-df7d-424c-bdc5-bb5dc7f1****
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s ModifyAssetImportantRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetImportantRequest) GoString() string {
	return s.String()
}

func (s *ModifyAssetImportantRequest) GetImportantCode() *int32 {
	return s.ImportantCode
}

func (s *ModifyAssetImportantRequest) GetUuidList() *string {
	return s.UuidList
}

func (s *ModifyAssetImportantRequest) SetImportantCode(v int32) *ModifyAssetImportantRequest {
	s.ImportantCode = &v
	return s
}

func (s *ModifyAssetImportantRequest) SetUuidList(v string) *ModifyAssetImportantRequest {
	s.UuidList = &v
	return s
}

func (s *ModifyAssetImportantRequest) Validate() error {
	return dara.Validate(s)
}
