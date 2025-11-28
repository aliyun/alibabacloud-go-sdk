// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *QueryBlockchainMetadataRequest
	GetBizChainId() *string
	SetContractName(v string) *QueryBlockchainMetadataRequest
	GetContractName() *string
	SetInvokeType(v string) *QueryBlockchainMetadataRequest
	GetInvokeType() *string
	SetIotDataDID(v string) *QueryBlockchainMetadataRequest
	GetIotDataDID() *string
	SetRegionId(v string) *QueryBlockchainMetadataRequest
	GetRegionId() *string
	SetTransactionId(v string) *QueryBlockchainMetadataRequest
	GetTransactionId() *string
}

type QueryBlockchainMetadataRequest struct {
	// This parameter is required.
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	ContractName *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	InvokeType   *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// This parameter is required.
	IotDataDID    *string `json:"IotDataDID,omitempty" xml:"IotDataDID,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s QueryBlockchainMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainMetadataRequest) GoString() string {
	return s.String()
}

func (s *QueryBlockchainMetadataRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *QueryBlockchainMetadataRequest) GetContractName() *string {
	return s.ContractName
}

func (s *QueryBlockchainMetadataRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *QueryBlockchainMetadataRequest) GetIotDataDID() *string {
	return s.IotDataDID
}

func (s *QueryBlockchainMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryBlockchainMetadataRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *QueryBlockchainMetadataRequest) SetBizChainId(v string) *QueryBlockchainMetadataRequest {
	s.BizChainId = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) SetContractName(v string) *QueryBlockchainMetadataRequest {
	s.ContractName = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) SetInvokeType(v string) *QueryBlockchainMetadataRequest {
	s.InvokeType = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) SetIotDataDID(v string) *QueryBlockchainMetadataRequest {
	s.IotDataDID = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) SetRegionId(v string) *QueryBlockchainMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) SetTransactionId(v string) *QueryBlockchainMetadataRequest {
	s.TransactionId = &v
	return s
}

func (s *QueryBlockchainMetadataRequest) Validate() error {
	return dara.Validate(s)
}
