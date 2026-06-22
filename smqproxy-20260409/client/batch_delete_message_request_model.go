// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReceiptHandles(v []*string) *BatchDeleteMessageRequest
	GetReceiptHandles() []*string
}

type BatchDeleteMessageRequest struct {
	ReceiptHandles []*string `json:"ReceiptHandles,omitempty" xml:"ReceiptHandles,omitempty" type:"Repeated"`
}

func (s BatchDeleteMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteMessageRequest) GetReceiptHandles() []*string {
	return s.ReceiptHandles
}

func (s *BatchDeleteMessageRequest) SetReceiptHandles(v []*string) *BatchDeleteMessageRequest {
	s.ReceiptHandles = v
	return s
}

func (s *BatchDeleteMessageRequest) Validate() error {
	return dara.Validate(s)
}
