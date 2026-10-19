// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMessageVisibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReceiptHandle(v string) *ChangeMessageVisibilityRequest
	GetReceiptHandle() *string
	SetVisibilityTimeout(v int32) *ChangeMessageVisibilityRequest
	GetVisibilityTimeout() *int32
}

type ChangeMessageVisibilityRequest struct {
	// The temporary handle of the message returned after the last consumption.
	//
	// example:
	//
	// 1-ODU4OTkzNDU5My0xNDMyNzI3ODI3LTItOA==
	ReceiptHandle *string `json:"receiptHandle,omitempty" xml:"receiptHandle,omitempty"`
	// The countdown from the current moment until the message becomes Active again. Valid values: 1 to 43200. Unit: seconds (1 second to 12 hours).
	//
	// example:
	//
	// 50
	VisibilityTimeout *int32 `json:"visibilityTimeout,omitempty" xml:"visibilityTimeout,omitempty"`
}

func (s ChangeMessageVisibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMessageVisibilityRequest) GoString() string {
	return s.String()
}

func (s *ChangeMessageVisibilityRequest) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *ChangeMessageVisibilityRequest) GetVisibilityTimeout() *int32 {
	return s.VisibilityTimeout
}

func (s *ChangeMessageVisibilityRequest) SetReceiptHandle(v string) *ChangeMessageVisibilityRequest {
	s.ReceiptHandle = &v
	return s
}

func (s *ChangeMessageVisibilityRequest) SetVisibilityTimeout(v int32) *ChangeMessageVisibilityRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *ChangeMessageVisibilityRequest) Validate() error {
	return dara.Validate(s)
}
