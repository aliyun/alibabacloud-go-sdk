// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchDeleteKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody
	GetSuccessKeys() []*string
}

type BatchDeleteKvWithHighCapacityResponseBody struct {
	FailKeys    []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchDeleteKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
