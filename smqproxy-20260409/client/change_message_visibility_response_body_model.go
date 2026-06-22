// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMessageVisibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextVisibleTime(v int64) *ChangeMessageVisibilityResponseBody
	GetNextVisibleTime() *int64
	SetReceiptHandle(v string) *ChangeMessageVisibilityResponseBody
	GetReceiptHandle() *string
}

type ChangeMessageVisibilityResponseBody struct {
	NextVisibleTime *int64  `json:"NextVisibleTime,omitempty" xml:"NextVisibleTime,omitempty"`
	ReceiptHandle   *string `json:"ReceiptHandle,omitempty" xml:"ReceiptHandle,omitempty"`
}

func (s ChangeMessageVisibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMessageVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMessageVisibilityResponseBody) GetNextVisibleTime() *int64 {
	return s.NextVisibleTime
}

func (s *ChangeMessageVisibilityResponseBody) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *ChangeMessageVisibilityResponseBody) SetNextVisibleTime(v int64) *ChangeMessageVisibilityResponseBody {
	s.NextVisibleTime = &v
	return s
}

func (s *ChangeMessageVisibilityResponseBody) SetReceiptHandle(v string) *ChangeMessageVisibilityResponseBody {
	s.ReceiptHandle = &v
	return s
}

func (s *ChangeMessageVisibilityResponseBody) Validate() error {
	return dara.Validate(s)
}
