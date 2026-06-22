// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVerifyCodeResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *SubmitVerifyCodeResultRequest
	GetMessageId() *string
	SetResult(v bool) *SubmitVerifyCodeResultRequest
	GetResult() *bool
	SetTo(v string) *SubmitVerifyCodeResultRequest
	GetTo() *string
}

type SubmitVerifyCodeResultRequest struct {
	// The message ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202605012020***
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The verification result.
	//
	// This parameter is required.
	//
	// example:
	//
	// TRUE
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// The recipient number.
	//
	// example:
	//
	// 86138000***
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SubmitVerifyCodeResultRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVerifyCodeResultRequest) GoString() string {
	return s.String()
}

func (s *SubmitVerifyCodeResultRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *SubmitVerifyCodeResultRequest) GetResult() *bool {
	return s.Result
}

func (s *SubmitVerifyCodeResultRequest) GetTo() *string {
	return s.To
}

func (s *SubmitVerifyCodeResultRequest) SetMessageId(v string) *SubmitVerifyCodeResultRequest {
	s.MessageId = &v
	return s
}

func (s *SubmitVerifyCodeResultRequest) SetResult(v bool) *SubmitVerifyCodeResultRequest {
	s.Result = &v
	return s
}

func (s *SubmitVerifyCodeResultRequest) SetTo(v string) *SubmitVerifyCodeResultRequest {
	s.To = &v
	return s
}

func (s *SubmitVerifyCodeResultRequest) Validate() error {
	return dara.Validate(s)
}
