// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchPutKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody
	GetSuccessKeys() []*string
}

type BatchPutKvWithHighCapacityResponseBody struct {
	// The keys that failed to be written.
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The keys that were written.
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchPutKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchPutKvWithHighCapacityResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchPutKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchPutKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchPutKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
