// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchReceiveMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumOfMessages(v int32) *BatchReceiveMessageRequest
	GetNumOfMessages() *int32
	SetWaitseconds(v int32) *BatchReceiveMessageRequest
	GetWaitseconds() *int32
}

type BatchReceiveMessageRequest struct {
	NumOfMessages *int32 `json:"numOfMessages,omitempty" xml:"numOfMessages,omitempty"`
	Waitseconds   *int32 `json:"waitseconds,omitempty" xml:"waitseconds,omitempty"`
}

func (s BatchReceiveMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchReceiveMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchReceiveMessageRequest) GetNumOfMessages() *int32 {
	return s.NumOfMessages
}

func (s *BatchReceiveMessageRequest) GetWaitseconds() *int32 {
	return s.Waitseconds
}

func (s *BatchReceiveMessageRequest) SetNumOfMessages(v int32) *BatchReceiveMessageRequest {
	s.NumOfMessages = &v
	return s
}

func (s *BatchReceiveMessageRequest) SetWaitseconds(v int32) *BatchReceiveMessageRequest {
	s.Waitseconds = &v
	return s
}

func (s *BatchReceiveMessageRequest) Validate() error {
	return dara.Validate(s)
}
