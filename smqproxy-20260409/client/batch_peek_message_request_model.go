// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPeekMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumOfMessages(v int32) *BatchPeekMessageRequest
	GetNumOfMessages() *int32
	SetPeekonly(v bool) *BatchPeekMessageRequest
	GetPeekonly() *bool
}

type BatchPeekMessageRequest struct {
	NumOfMessages *int32 `json:"numOfMessages,omitempty" xml:"numOfMessages,omitempty"`
	Peekonly      *bool  `json:"peekonly,omitempty" xml:"peekonly,omitempty"`
}

func (s BatchPeekMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPeekMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchPeekMessageRequest) GetNumOfMessages() *int32 {
	return s.NumOfMessages
}

func (s *BatchPeekMessageRequest) GetPeekonly() *bool {
	return s.Peekonly
}

func (s *BatchPeekMessageRequest) SetNumOfMessages(v int32) *BatchPeekMessageRequest {
	s.NumOfMessages = &v
	return s
}

func (s *BatchPeekMessageRequest) SetPeekonly(v bool) *BatchPeekMessageRequest {
	s.Peekonly = &v
	return s
}

func (s *BatchPeekMessageRequest) Validate() error {
	return dara.Validate(s)
}
