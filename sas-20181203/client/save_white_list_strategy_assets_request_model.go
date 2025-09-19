// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SaveWhiteListStrategyAssetsRequest
	GetLang() *string
	SetOperations(v string) *SaveWhiteListStrategyAssetsRequest
	GetOperations() *string
	SetRelationType(v int32) *SaveWhiteListStrategyAssetsRequest
	GetRelationType() *int32
	SetSourceIp(v string) *SaveWhiteListStrategyAssetsRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *SaveWhiteListStrategyAssetsRequest
	GetStrategyId() *int64
}

type SaveWhiteListStrategyAssetsRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation that you want to perform. This parameter is in the JSON format. The value is case-sensitive. The value contains the following fields:
	//
	// 	- **status**: the operation. Valid values:
	//
	//     	- **0**: the delete operation.
	//
	//     	- **1**: the add operation.
	//
	// 	- **target**: the UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the value of **target*	- from the response parameter Uuid.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"status":0,"target":"c98dcd24-fa57-4759-b5ec-f8a4ffeed****"}]
	Operations *string `json:"Operations,omitempty" xml:"Operations,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **1**: learning policy.
	//
	// 	- **2**: application policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RelationType *int32 `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 115.193.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2730
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s SaveWhiteListStrategyAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyAssetsRequest) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveWhiteListStrategyAssetsRequest) GetOperations() *string {
	return s.Operations
}

func (s *SaveWhiteListStrategyAssetsRequest) GetRelationType() *int32 {
	return s.RelationType
}

func (s *SaveWhiteListStrategyAssetsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *SaveWhiteListStrategyAssetsRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *SaveWhiteListStrategyAssetsRequest) SetLang(v string) *SaveWhiteListStrategyAssetsRequest {
	s.Lang = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsRequest) SetOperations(v string) *SaveWhiteListStrategyAssetsRequest {
	s.Operations = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsRequest) SetRelationType(v int32) *SaveWhiteListStrategyAssetsRequest {
	s.RelationType = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsRequest) SetSourceIp(v string) *SaveWhiteListStrategyAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsRequest) SetStrategyId(v int64) *SaveWhiteListStrategyAssetsRequest {
	s.StrategyId = &v
	return s
}

func (s *SaveWhiteListStrategyAssetsRequest) Validate() error {
	return dara.Validate(s)
}
