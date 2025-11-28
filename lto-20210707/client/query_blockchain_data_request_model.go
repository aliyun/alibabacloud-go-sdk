// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *QueryBlockchainDataRequest
	GetBizChainId() *string
	SetContractName(v string) *QueryBlockchainDataRequest
	GetContractName() *string
	SetInvokeType(v string) *QueryBlockchainDataRequest
	GetInvokeType() *string
	SetIotDataDID(v string) *QueryBlockchainDataRequest
	GetIotDataDID() *string
	SetRegionId(v string) *QueryBlockchainDataRequest
	GetRegionId() *string
	SetTransactionId(v string) *QueryBlockchainDataRequest
	GetTransactionId() *string
}

type QueryBlockchainDataRequest struct {
	// This parameter is required.
	BizChainId   *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	ContractName *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	InvokeType   *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// This parameter is required.
	IotDataDID    *string `json:"IotDataDID,omitempty" xml:"IotDataDID,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s QueryBlockchainDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainDataRequest) GoString() string {
	return s.String()
}

func (s *QueryBlockchainDataRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *QueryBlockchainDataRequest) GetContractName() *string {
	return s.ContractName
}

func (s *QueryBlockchainDataRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *QueryBlockchainDataRequest) GetIotDataDID() *string {
	return s.IotDataDID
}

func (s *QueryBlockchainDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryBlockchainDataRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *QueryBlockchainDataRequest) SetBizChainId(v string) *QueryBlockchainDataRequest {
	s.BizChainId = &v
	return s
}

func (s *QueryBlockchainDataRequest) SetContractName(v string) *QueryBlockchainDataRequest {
	s.ContractName = &v
	return s
}

func (s *QueryBlockchainDataRequest) SetInvokeType(v string) *QueryBlockchainDataRequest {
	s.InvokeType = &v
	return s
}

func (s *QueryBlockchainDataRequest) SetIotDataDID(v string) *QueryBlockchainDataRequest {
	s.IotDataDID = &v
	return s
}

func (s *QueryBlockchainDataRequest) SetRegionId(v string) *QueryBlockchainDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryBlockchainDataRequest) SetTransactionId(v string) *QueryBlockchainDataRequest {
	s.TransactionId = &v
	return s
}

func (s *QueryBlockchainDataRequest) Validate() error {
	return dara.Validate(s)
}
