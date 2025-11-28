// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBlockchainMetadataResponseBody
	GetCode() *string
	SetData(v *QueryBlockchainMetadataResponseBodyData) *QueryBlockchainMetadataResponseBody
	GetData() *QueryBlockchainMetadataResponseBodyData
	SetMessage(v string) *QueryBlockchainMetadataResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBlockchainMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBlockchainMetadataResponseBody
	GetSuccess() *bool
}

type QueryBlockchainMetadataResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryBlockchainMetadataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBlockchainMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlockchainMetadataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBlockchainMetadataResponseBody) GetData() *QueryBlockchainMetadataResponseBodyData {
	return s.Data
}

func (s *QueryBlockchainMetadataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBlockchainMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBlockchainMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBlockchainMetadataResponseBody) SetCode(v string) *QueryBlockchainMetadataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBody) SetData(v *QueryBlockchainMetadataResponseBodyData) *QueryBlockchainMetadataResponseBody {
	s.Data = v
	return s
}

func (s *QueryBlockchainMetadataResponseBody) SetMessage(v string) *QueryBlockchainMetadataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBody) SetRequestId(v string) *QueryBlockchainMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBody) SetSuccess(v bool) *QueryBlockchainMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBlockchainMetadataResponseBodyData struct {
	BlockHash   *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	BlockNumber *string `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	IotId       *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	MemberName  *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TxHash      *string `json:"TxHash,omitempty" xml:"TxHash,omitempty"`
}

func (s QueryBlockchainMetadataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainMetadataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBlockchainMetadataResponseBodyData) GetBlockHash() *string {
	return s.BlockHash
}

func (s *QueryBlockchainMetadataResponseBodyData) GetBlockNumber() *string {
	return s.BlockNumber
}

func (s *QueryBlockchainMetadataResponseBodyData) GetIotId() *string {
	return s.IotId
}

func (s *QueryBlockchainMetadataResponseBodyData) GetMemberName() *string {
	return s.MemberName
}

func (s *QueryBlockchainMetadataResponseBodyData) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryBlockchainMetadataResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryBlockchainMetadataResponseBodyData) GetTxHash() *string {
	return s.TxHash
}

func (s *QueryBlockchainMetadataResponseBodyData) SetBlockHash(v string) *QueryBlockchainMetadataResponseBodyData {
	s.BlockHash = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetBlockNumber(v string) *QueryBlockchainMetadataResponseBodyData {
	s.BlockNumber = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetIotId(v string) *QueryBlockchainMetadataResponseBodyData {
	s.IotId = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetMemberName(v string) *QueryBlockchainMetadataResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetProductKey(v string) *QueryBlockchainMetadataResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetTimestamp(v int64) *QueryBlockchainMetadataResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) SetTxHash(v string) *QueryBlockchainMetadataResponseBodyData {
	s.TxHash = &v
	return s
}

func (s *QueryBlockchainMetadataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
