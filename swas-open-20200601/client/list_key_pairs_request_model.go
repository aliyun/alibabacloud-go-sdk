// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListKeyPairsRequest
	GetClientToken() *string
	SetKeyPairName(v string) *ListKeyPairsRequest
	GetKeyPairName() *string
	SetPageNumber(v int32) *ListKeyPairsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKeyPairsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKeyPairsRequest
	GetRegionId() *string
}

type ListKeyPairsRequest struct {
	// The client token that is used to ensure the idempotence of the request. Generate a unique value for this parameter from your client to ensure that different requests use unique values. ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters and Chinese characters. It must start with an uppercase or lowercase letter or a Chinese character, and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The page number of the key pair list. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyPairsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListKeyPairsRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ListKeyPairsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKeyPairsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeyPairsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKeyPairsRequest) SetClientToken(v string) *ListKeyPairsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListKeyPairsRequest) SetKeyPairName(v string) *ListKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *ListKeyPairsRequest) SetPageNumber(v int32) *ListKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeyPairsRequest) SetPageSize(v int32) *ListKeyPairsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeyPairsRequest) SetRegionId(v string) *ListKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
