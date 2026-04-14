// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlockSendingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int32) *ListBlockSendingRequest
	GetBeginTime() *int32
	SetBlockEmail(v string) *ListBlockSendingRequest
	GetBlockEmail() *string
	SetBlockType(v string) *ListBlockSendingRequest
	GetBlockType() *string
	SetEndTime(v int32) *ListBlockSendingRequest
	GetEndTime() *int32
	SetMaxResults(v int32) *ListBlockSendingRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBlockSendingRequest
	GetNextToken() *string
	SetSenderEmail(v string) *ListBlockSendingRequest
	GetSenderEmail() *string
}

type ListBlockSendingRequest struct {
	// example:
	//
	// 1731463398242
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// xxxx@rcpt.com
	BlockEmail *string `json:"BlockEmail,omitempty" xml:"BlockEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNSUB
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// example:
	//
	// 1732463398242
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// xxxxxyyyyyy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// xxxx@sender.com
	SenderEmail *string `json:"SenderEmail,omitempty" xml:"SenderEmail,omitempty"`
}

func (s ListBlockSendingRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBlockSendingRequest) GoString() string {
	return s.String()
}

func (s *ListBlockSendingRequest) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *ListBlockSendingRequest) GetBlockEmail() *string {
	return s.BlockEmail
}

func (s *ListBlockSendingRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *ListBlockSendingRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *ListBlockSendingRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBlockSendingRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBlockSendingRequest) GetSenderEmail() *string {
	return s.SenderEmail
}

func (s *ListBlockSendingRequest) SetBeginTime(v int32) *ListBlockSendingRequest {
	s.BeginTime = &v
	return s
}

func (s *ListBlockSendingRequest) SetBlockEmail(v string) *ListBlockSendingRequest {
	s.BlockEmail = &v
	return s
}

func (s *ListBlockSendingRequest) SetBlockType(v string) *ListBlockSendingRequest {
	s.BlockType = &v
	return s
}

func (s *ListBlockSendingRequest) SetEndTime(v int32) *ListBlockSendingRequest {
	s.EndTime = &v
	return s
}

func (s *ListBlockSendingRequest) SetMaxResults(v int32) *ListBlockSendingRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBlockSendingRequest) SetNextToken(v string) *ListBlockSendingRequest {
	s.NextToken = &v
	return s
}

func (s *ListBlockSendingRequest) SetSenderEmail(v string) *ListBlockSendingRequest {
	s.SenderEmail = &v
	return s
}

func (s *ListBlockSendingRequest) Validate() error {
	return dara.Validate(s)
}
