// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnblockSendingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockEmail(v string) *UnblockSendingRequest
	GetBlockEmail() *string
	SetBlockType(v string) *UnblockSendingRequest
	GetBlockType() *string
	SetSenderEmail(v string) *UnblockSendingRequest
	GetSenderEmail() *string
}

type UnblockSendingRequest struct {
	// The blacklisted recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// recipient@yyy.com
	BlockEmail *string `json:"BlockEmail,omitempty" xml:"BlockEmail,omitempty"`
	// The blacklist type.
	//
	// - UNSUB: Unsubscribe
	//
	// - REPORT: Complaint
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSUB
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The sender address.
	//
	// This parameter is required.
	//
	// example:
	//
	// sender@xxx.com
	SenderEmail *string `json:"SenderEmail,omitempty" xml:"SenderEmail,omitempty"`
}

func (s UnblockSendingRequest) String() string {
	return dara.Prettify(s)
}

func (s UnblockSendingRequest) GoString() string {
	return s.String()
}

func (s *UnblockSendingRequest) GetBlockEmail() *string {
	return s.BlockEmail
}

func (s *UnblockSendingRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *UnblockSendingRequest) GetSenderEmail() *string {
	return s.SenderEmail
}

func (s *UnblockSendingRequest) SetBlockEmail(v string) *UnblockSendingRequest {
	s.BlockEmail = &v
	return s
}

func (s *UnblockSendingRequest) SetBlockType(v string) *UnblockSendingRequest {
	s.BlockType = &v
	return s
}

func (s *UnblockSendingRequest) SetSenderEmail(v string) *UnblockSendingRequest {
	s.SenderEmail = &v
	return s
}

func (s *UnblockSendingRequest) Validate() error {
	return dara.Validate(s)
}
