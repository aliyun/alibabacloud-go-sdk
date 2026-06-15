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
	// The start of the time range to query blocked emails.
	//
	// example:
	//
	// 1763973206
	BeginTime *int32 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The recipient email address.
	//
	// example:
	//
	// xxxx@rcpt.com
	BlockEmail *string `json:"BlockEmail,omitempty" xml:"BlockEmail,omitempty"`
	// The type of block.
	//
	// - UNSUB: Unsubscribe
	//
	// - REPORT: Spam report
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSUB
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The end of the time range to query blocked emails.
	//
	// example:
	//
	// 1764146006
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return.<br>
	//
	// Valid values: 1 to 500.<br>
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the query. If you do not specify this parameter, the query starts from the beginning of the results.
	//
	// example:
	//
	// xxxxxyyyyyy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sender email address.
	//
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
