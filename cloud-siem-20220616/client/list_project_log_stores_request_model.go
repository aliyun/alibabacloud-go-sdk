// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLogStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListProjectLogStoresRequest
	GetRegionId() *string
	SetSourceLogCode(v string) *ListProjectLogStoresRequest
	GetSourceLogCode() *string
	SetSourceProdCode(v string) *ListProjectLogStoresRequest
	GetSourceProdCode() *string
	SetSubUserId(v int64) *ListProjectLogStoresRequest
	GetSubUserId() *int64
}

type ListProjectLogStoresRequest struct {
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The log code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The code of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListProjectLogStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLogStoresRequest) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectLogStoresRequest) GetSourceLogCode() *string {
	return s.SourceLogCode
}

func (s *ListProjectLogStoresRequest) GetSourceProdCode() *string {
	return s.SourceProdCode
}

func (s *ListProjectLogStoresRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListProjectLogStoresRequest) SetRegionId(v string) *ListProjectLogStoresRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSourceLogCode(v string) *ListProjectLogStoresRequest {
	s.SourceLogCode = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSourceProdCode(v string) *ListProjectLogStoresRequest {
	s.SourceProdCode = &v
	return s
}

func (s *ListProjectLogStoresRequest) SetSubUserId(v int64) *ListProjectLogStoresRequest {
	s.SubUserId = &v
	return s
}

func (s *ListProjectLogStoresRequest) Validate() error {
	return dara.Validate(s)
}
