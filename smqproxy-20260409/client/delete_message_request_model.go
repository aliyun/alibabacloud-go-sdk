// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReceiptHandle(v string) *DeleteMessageRequest
	GetReceiptHandle() *string
}

type DeleteMessageRequest struct {
	ReceiptHandle *string `json:"ReceiptHandle,omitempty" xml:"ReceiptHandle,omitempty"`
}

func (s DeleteMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageRequest) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *DeleteMessageRequest) SetReceiptHandle(v string) *DeleteMessageRequest {
	s.ReceiptHandle = &v
	return s
}

func (s *DeleteMessageRequest) Validate() error {
	return dara.Validate(s)
}
