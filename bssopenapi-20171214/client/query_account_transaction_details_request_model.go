// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *QueryAccountTransactionDetailsRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *QueryAccountTransactionDetailsRequest
	GetCreateTimeStart() *string
	SetMaxResults(v int32) *QueryAccountTransactionDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryAccountTransactionDetailsRequest
	GetNextToken() *string
	SetRecordID(v string) *QueryAccountTransactionDetailsRequest
	GetRecordID() *string
	SetTransactionChannel(v string) *QueryAccountTransactionDetailsRequest
	GetTransactionChannel() *string
	SetTransactionChannelSN(v string) *QueryAccountTransactionDetailsRequest
	GetTransactionChannelSN() *string
	SetTransactionNumber(v string) *QueryAccountTransactionDetailsRequest
	GetTransactionNumber() *string
	SetTransactionType(v string) *QueryAccountTransactionDetailsRequest
	GetTransactionType() *string
}

type QueryAccountTransactionDetailsRequest struct {
	// The end of the creation time range to query.
	//
	// example:
	//
	// 2022-12-20
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The beginning of the creation time range to query.
	//
	// example:
	//
	// 2022-01-20
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// This parameter is invalid.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for paging.
	//
	// example:
	//
	// ABEDSDS124DASA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the order or bill.
	//
	// example:
	//
	// 2022120336190912
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The transaction channel.
	//
	// example:
	//
	// ALIPAY
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// The serial number of the transaction channel.
	//
	// example:
	//
	// 2022112122001470591458665933
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	// The number of the transaction.
	//
	// example:
	//
	// 410874027490089
	TransactionNumber *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// The type of the transaction.
	//
	// example:
	//
	// CHARGE
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s QueryAccountTransactionDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *QueryAccountTransactionDetailsRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *QueryAccountTransactionDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryAccountTransactionDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryAccountTransactionDetailsRequest) GetRecordID() *string {
	return s.RecordID
}

func (s *QueryAccountTransactionDetailsRequest) GetTransactionChannel() *string {
	return s.TransactionChannel
}

func (s *QueryAccountTransactionDetailsRequest) GetTransactionChannelSN() *string {
	return s.TransactionChannelSN
}

func (s *QueryAccountTransactionDetailsRequest) GetTransactionNumber() *string {
	return s.TransactionNumber
}

func (s *QueryAccountTransactionDetailsRequest) GetTransactionType() *string {
	return s.TransactionType
}

func (s *QueryAccountTransactionDetailsRequest) SetCreateTimeEnd(v string) *QueryAccountTransactionDetailsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetCreateTimeStart(v string) *QueryAccountTransactionDetailsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetMaxResults(v int32) *QueryAccountTransactionDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetNextToken(v string) *QueryAccountTransactionDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetRecordID(v string) *QueryAccountTransactionDetailsRequest {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionChannel(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionChannelSN(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionNumber(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionType(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) Validate() error {
	return dara.Validate(s)
}
